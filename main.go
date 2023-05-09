package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	env "github.com/caarlos0/env/v6"
	"github.com/jake-dog/opensimdash/codemasters"
	"github.com/tarm/serial"
)

type Config struct {
	Index      int    `env:"INDEX" envDefault:"0"`
	Listen     string `env:"LISTEN_UDP" envDefault:"127.0.0.1:20777"`
	ListenHttp string `env:"LISTEN_HTTP" envDefault:"127.0.0.1:8123"`
}

type Params struct {
	Steer    float32
	Clutch   float32
	Brake    float32
	Throttle float32
	Gear     int
	Active   bool
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
			p := status.Get()
			p.Active = false
			ch <- p
		})
		var old codemasters.DirtPacket
		b := make([]byte, 4096)
		last := time.Now()
		for {
			n, _, err := conn.ReadFrom(b)
			if err != nil {
				done <- err
			}
			now := time.Now()
			if now.Sub(last) < 15*time.Millisecond {
				continue
			}
			last = now
			var pkt codemasters.DirtPacket
			pkt.Decode(b[:n])
			changed := false
			changed = changed || old.Brake != pkt.Brake
			changed = changed || old.Throttle != pkt.Throttle
			changed = changed || old.Clutch != pkt.Clutch
			changed = changed || old.Steer != pkt.Steer
			changed = changed || old.LapTime != pkt.LapTime
			changed = changed || old.LapDistance != pkt.LapDistance
			old = pkt
			if changed {
				timer.Reset(5 * time.Second)
				status.Activate()
			}
			status.Update(&pkt)
			ch <- status.Get()
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
	timeout := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-r.Context().Done():
			return
		case v := <-ch:
			timeout.Reset(30 * time.Second)
			b, _ := json.Marshal(v)
			fmt.Fprintf(w, "data: %s\n\n", string(b))
		case <-timeout.C:
			fmt.Fprintf(w, "data: \n\n")
		}
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
}

func forwardProc(ctx context.Context, p1, p2 string) error {
	src, err := serial.OpenPort(&serial.Config{Name: os.Args[1], Baud: 115200})
	if err != nil {
		return err
	}
	srcOnce := sync.Once{}
	defer srcOnce.Do(func() { src.Close() })

	dst, err := serial.OpenPort(&serial.Config{Name: os.Args[2], Baud: 115200})
	if err != nil {
		return err
	}
	defer dst.Close()

	go func() {
		defer srcOnce.Do(func() { src.Close() })
		io.Copy(os.Stdout, dst)
	}()
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func forward(ctx context.Context) {
	if len(os.Args) != 3 {
		return
	}
	p1, p2 := os.Args[1], os.Args[2]
	for {
		if err := forwardProc(ctx, p1, p2); err != nil {
			log.Print(err)
		}
		select {
		case <-ctx.Done():
			return
		default:
		}
		time.Sleep(10 * time.Second)
	}
}

//go:embed static/*
var contents embed.FS

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go forward(ctx)
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
