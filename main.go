package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	env "github.com/caarlos0/env/v6"
	"github.com/jake-dog/opensimdash/codemasters"
	"github.com/simulatedsimian/joystick"
)

type Config struct {
	Index      int    `env:"INDEX" envDefault:"0"`
	Listen     string `env:"LISTEN_UDP" envDefault:"127.0.0.1:20777"`
	ListenHttp string `env:"LISTEN_HTTP" envDefault:"127.0.0.1:8123"`
}

type Params struct {
	Steer     float32
	Clutch    float32
	Brake     float32
	Throttle  float32
	HandBrake float32
	Gear      int
	Active    bool
}

type Status struct {
	mu sync.RWMutex
	Params
}

func (status *Status) Activate() {
	status.mu.Lock()
	defer status.mu.Unlock()
	status.Active = true
}

func (status *Status) Deactivate() {
	status.mu.Lock()
	defer status.mu.Unlock()
	status.Active = false
}

func (status *Status) Update(pkt *codemasters.DirtPacket) {
	status.mu.Lock()
	defer status.mu.Unlock()
	pol := float32(1)
	if pkt.Track_size == 0 { // for WRC Generations
		pol = -1
	}
	status.Steer = pol * pkt.Steer
	status.Clutch = pkt.Clutch
	status.Brake = pkt.Brake
	status.Throttle = pkt.Throttle
	status.Gear = int(pkt.Gear)
}

func (status *Status) SetHandBrake(v float32) {
	status.mu.Lock()
	defer status.mu.Unlock()
	v = v/0.7 - 0.1
	if v < 0 {
		v = 0
	}
	if v > 1 {
		v = 1
	}
	status.HandBrake = v
}

func (status *Status) Get() Params {
	status.mu.RLock()
	defer status.mu.RUnlock()
	return status.Params
}

var (
	config Config
	status Status
)

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	config = Config{}
	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}
}

func jsReciver(ctx context.Context) error {
	var js joystick.Joystick
	for i := 0; i < 8; i++ {
		id := (config.Index + i) % 8
		d, err := joystick.Open(id)
		if err != nil {
			continue
		}
		log.Printf("found joystick device id=%d", id)
		js = d
		break
	}
	if js == nil {
		return fmt.Errorf("not found controller")
	}
	ticker := time.NewTicker(10 * time.Millisecond)
	defer js.Close()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			state, err := js.Read()
			if err != nil {
				return err
			}
			status.SetHandBrake((float32(state.AxisData[2]) + 32767) / 65535)
		}
	}
}

func udpReceiver(ctx context.Context, ch chan<- Params) error {
	host, port, err := net.SplitHostPort(config.Listen)
	if err != nil {
		return err
	}
	addr := net.JoinHostPort(host, port)
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Println("listen udp:", addr)
	defer log.Println("udp closed:", addr)
	done := make(chan error, 1)
	go func() {
		timer := time.AfterFunc(5*time.Second, func() {
			status.Deactivate()
			ch <- status.Get()
		})
		lapT := float32(0.0)
		lapD := float32(0.0)
		b := make([]byte, 4096)
		for {
			n, _, err := conn.ReadFrom(b)
			if err != nil {
				done <- err
			}
			var pkt codemasters.DirtPacket
			pkt.Decode(b[:n])
			changed := lapT != pkt.LapTime || lapD != pkt.LapDistance
			changed = changed || pkt.Throttle > 0.0 || pkt.Speed > 0.0
			if changed {
				lapT = pkt.LapTime
				lapD = pkt.LapDistance
				timer.Reset(5 * time.Second)
				status.Activate()
			}
			status.Update(&pkt)
			p := status.Get()
			ch <- p
		}
	}()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
	}
	return nil
}

var (
	subscribe   = make(chan chan<- Params, 1)
	unsubscribe = make(chan chan<- Params, 1)
)

func proc(ctx context.Context, publish <-chan Params) {
	m := map[chan<- Params]struct{}{}
	for {
		select {
		case <-ctx.Done():
			return
		case v := <-subscribe:
			m[v] = struct{}{}
		case v := <-unsubscribe:
			delete(m, v)
			close(v)
		case v := <-publish:
			for c := range m {
				c <- v
			}
		}
	}
}

func sse(w http.ResponseWriter, r *http.Request) {
	log.Printf("connect from: %v", r.RemoteAddr)
	defer log.Printf("disconnect from: %v", r.RemoteAddr)
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	ch := make(chan Params, 64)
	subscribe <- ch
	defer func() {
		unsubscribe <- ch
	}()
	timeout := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-r.Context().Done():
			return
		case v := <-ch:
			timeout.Reset(5 * time.Second)
			b, _ := json.Marshal(v)
			fmt.Fprintf(w, "data: %s\n\n", string(b))
		case <-timeout.C:
			v := Params{Active: false}
			b, _ := json.Marshal(v)
			fmt.Fprintf(w, "data: %s\n\n", string(b))
		}
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
}

//go:embed static/*
var contents embed.FS

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for {
			if err := jsReciver(ctx); err != nil {
				log.Print(err)
				time.Sleep(5 * time.Second)
				continue
			}
			break
		}
	}()
	ch := make(chan Params, 64)
	go func() {
		for {
			if err := udpReceiver(ctx, ch); err != nil {
				log.Print(err)
				time.Sleep(5 * time.Second)
				continue
			}
			break
		}
	}()
	go proc(ctx, ch)
	static, err := fs.Sub(contents, "static")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(static)))
	http.Handle("/sse", http.HandlerFunc(sse))
	log.Print("listen start http:", config.ListenHttp)
	defer log.Print("program terminated")
	if err := http.ListenAndServe(config.ListenHttp, nil); err != nil {
		log.Fatal(err)
	}
}
