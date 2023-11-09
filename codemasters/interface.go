package codemasters

type Telemetry interface {
	Steering() float32
	Throttle() float32
	Brake() float32
	Clutch() float32
	Handbrake() float32
	Gear() int
	RPM() float32
	MaxRPM() float32
	Speed() float32
	StageDistance() float32
}

func Decode(b []byte) (Telemetry, error) {
	switch len(b) {
	case PacketEASportsWRCLength:
		pkt := &PacketEASportsWRC{}
		if err := pkt.UnmarshalBinary(b); err != nil {
			return nil, err
		}
		return pkt, nil
	default:
		pkt := &PacketDirtSeries{}
		if err := pkt.UnmarshalBinary(b); err != nil {
			return nil, err
		}
		return pkt, nil
	}
}
