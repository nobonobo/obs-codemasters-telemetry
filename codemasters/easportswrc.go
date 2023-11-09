package codemasters

import (
	"encoding/binary"
	"fmt"
	"math"
)

const PacketEASportsWRCLength = 237

type PacketEASportsWRC struct {
	PacketUid                 uint64  // 0
	GameTotalTime             float32 // 1
	GameDeltaTime             float32 // 2
	GameFrameCount            uint64  // 3
	ShiftlightsFraction       float32 // 4
	ShiftlightsRpmStart       float32 // 5
	ShiftlightsRpmEnd         float32 // 6
	ShiftlightsRpmValid       bool    // 7
	VehicleGearIndex          uint8   // 8
	VehicleGearIndexNeutral   uint8   // 9
	VehicleGearIndexReverse   uint8   // 10
	VehicleGearMaximum        uint8   // 11
	VehicleSpeed              float32 // 12
	VehicleTransmissionSpeed  float32 // 13
	VehiclePositionX          float32 // 14
	VehiclePositionY          float32 // 15
	VehiclePositionZ          float32 // 16
	VehicleVelocityX          float32 // 17
	VehicleVelocityY          float32 // 18
	VehicleVelocityZ          float32 // 19
	VehicleAccelerationX      float32 // 20
	VehicleAccelerationY      float32 // 21
	VehicleAccelerationZ      float32 // 22
	VehicleLeftDirectionX     float32 // 23
	VehicleLeftDirectionY     float32 // 24
	VehicleLeftDirectionZ     float32 // 25
	VehicleForwardDirectionX  float32 // 26
	VehicleForwardDirectionY  float32 // 27
	VehicleForwardDirectionZ  float32 // 28
	VehicleUpDirectionX       float32 // 29
	VehicleUpDirectionY       float32 // 30
	VehicleUpDirectionZ       float32 // 31
	VehicleHubPositionBl      float32 // 32
	VehicleHubPositionBr      float32 // 33
	VehicleHubPositionFl      float32 // 34
	VehicleHubPositionFr      float32 // 35
	VehicleHubVelocityBl      float32 // 36
	VehicleHubVelocityBr      float32 // 37
	VehicleHubVelocityFl      float32 // 38
	VehicleHubVelocityFr      float32 // 39
	VehicleCpForwardSpeedBl   float32 // 40
	VehicleCpForwardSpeedBr   float32 // 41
	VehicleCpForwardSpeedFl   float32 // 42
	VehicleCpForwardSpeedFr   float32 // 43
	VehicleBrakeTemperatureBl float32 // 44
	VehicleBrakeTemperatureBr float32 // 45
	VehicleBrakeTemperatureFl float32 // 46
	VehicleBrakeTemperatureFr float32 // 47
	VehicleEngineRpmMax       float32 // 48
	VehicleEngineRpmIdle      float32 // 49
	VehicleEngineRpmCurrent   float32 // 50
	VehicleThrottle           float32 // 51
	VehicleBrake              float32 // 52
	VehicleClutch             float32 // 53
	VehicleSteering           float32 // 54
	VehicleHandbrake          float32 // 55
	StageCurrentTime          float32 // 56
	StageCurrentDistance      float64 // 57
	StageLength               float64 // 58
}

func (p *PacketEASportsWRC) UnmarshalBinary(b []byte) error {
	if len(b) < PacketEASportsWRCLength {
		return fmt.Errorf("invalid packet size: %d", len(b))
	}
	p.PacketUid = binary.LittleEndian.Uint64(b[0:8])
	p.GameTotalTime = math.Float32frombits(binary.LittleEndian.Uint32(b[8:12]))
	p.GameDeltaTime = math.Float32frombits(binary.LittleEndian.Uint32(b[12:16]))
	p.GameFrameCount = binary.LittleEndian.Uint64(b[16:24])
	p.ShiftlightsFraction = math.Float32frombits(binary.LittleEndian.Uint32(b[24:28]))
	p.ShiftlightsRpmStart = math.Float32frombits(binary.LittleEndian.Uint32(b[28:32]))
	p.ShiftlightsRpmEnd = math.Float32frombits(binary.LittleEndian.Uint32(b[32:36]))
	p.ShiftlightsRpmValid = b[36] != 0
	p.VehicleGearIndex = b[37]
	p.VehicleGearIndexNeutral = b[38]
	p.VehicleGearIndexReverse = b[39]
	p.VehicleGearMaximum = b[40]
	p.VehicleSpeed = math.Float32frombits(binary.LittleEndian.Uint32(b[41:45]))
	p.VehicleTransmissionSpeed = math.Float32frombits(binary.LittleEndian.Uint32(b[45:49]))
	p.VehiclePositionX = math.Float32frombits(binary.LittleEndian.Uint32(b[49:53]))
	p.VehiclePositionY = math.Float32frombits(binary.LittleEndian.Uint32(b[53:57]))
	p.VehiclePositionZ = math.Float32frombits(binary.LittleEndian.Uint32(b[57:61]))
	p.VehicleVelocityX = math.Float32frombits(binary.LittleEndian.Uint32(b[61:65]))
	p.VehicleVelocityY = math.Float32frombits(binary.LittleEndian.Uint32(b[65:69]))
	p.VehicleVelocityZ = math.Float32frombits(binary.LittleEndian.Uint32(b[69:73]))
	p.VehicleAccelerationX = math.Float32frombits(binary.LittleEndian.Uint32(b[73:77]))
	p.VehicleAccelerationY = math.Float32frombits(binary.LittleEndian.Uint32(b[77:81]))
	p.VehicleAccelerationZ = math.Float32frombits(binary.LittleEndian.Uint32(b[81:85]))
	p.VehicleLeftDirectionX = math.Float32frombits(binary.LittleEndian.Uint32(b[85:89]))
	p.VehicleLeftDirectionY = math.Float32frombits(binary.LittleEndian.Uint32(b[89:93]))
	p.VehicleLeftDirectionZ = math.Float32frombits(binary.LittleEndian.Uint32(b[93:97]))
	p.VehicleForwardDirectionX = math.Float32frombits(binary.LittleEndian.Uint32(b[97:101]))
	p.VehicleForwardDirectionY = math.Float32frombits(binary.LittleEndian.Uint32(b[101:105]))
	p.VehicleForwardDirectionZ = math.Float32frombits(binary.LittleEndian.Uint32(b[105:109]))
	p.VehicleUpDirectionX = math.Float32frombits(binary.LittleEndian.Uint32(b[109:113]))
	p.VehicleUpDirectionY = math.Float32frombits(binary.LittleEndian.Uint32(b[113:117]))
	p.VehicleUpDirectionZ = math.Float32frombits(binary.LittleEndian.Uint32(b[117:121]))
	p.VehicleHubPositionBl = math.Float32frombits(binary.LittleEndian.Uint32(b[121:125]))
	p.VehicleHubPositionBr = math.Float32frombits(binary.LittleEndian.Uint32(b[125:129]))
	p.VehicleHubPositionFl = math.Float32frombits(binary.LittleEndian.Uint32(b[129:133]))
	p.VehicleHubPositionFr = math.Float32frombits(binary.LittleEndian.Uint32(b[133:137]))
	p.VehicleHubVelocityBl = math.Float32frombits(binary.LittleEndian.Uint32(b[137:141]))
	p.VehicleHubVelocityBr = math.Float32frombits(binary.LittleEndian.Uint32(b[141:145]))
	p.VehicleHubVelocityFl = math.Float32frombits(binary.LittleEndian.Uint32(b[145:149]))
	p.VehicleHubVelocityFr = math.Float32frombits(binary.LittleEndian.Uint32(b[149:153]))
	p.VehicleCpForwardSpeedBl = math.Float32frombits(binary.LittleEndian.Uint32(b[153:157]))
	p.VehicleCpForwardSpeedBr = math.Float32frombits(binary.LittleEndian.Uint32(b[157:161]))
	p.VehicleCpForwardSpeedFl = math.Float32frombits(binary.LittleEndian.Uint32(b[161:165]))
	p.VehicleCpForwardSpeedFr = math.Float32frombits(binary.LittleEndian.Uint32(b[165:169]))
	p.VehicleBrakeTemperatureBl = math.Float32frombits(binary.LittleEndian.Uint32(b[169:173]))
	p.VehicleBrakeTemperatureBr = math.Float32frombits(binary.LittleEndian.Uint32(b[173:177]))
	p.VehicleBrakeTemperatureFl = math.Float32frombits(binary.LittleEndian.Uint32(b[177:181]))
	p.VehicleBrakeTemperatureFr = math.Float32frombits(binary.LittleEndian.Uint32(b[181:185]))
	p.VehicleEngineRpmMax = math.Float32frombits(binary.LittleEndian.Uint32(b[185:189]))
	p.VehicleEngineRpmIdle = math.Float32frombits(binary.LittleEndian.Uint32(b[189:193]))
	p.VehicleEngineRpmCurrent = math.Float32frombits(binary.LittleEndian.Uint32(b[193:197]))
	p.VehicleThrottle = math.Float32frombits(binary.LittleEndian.Uint32(b[197:201]))
	p.VehicleBrake = math.Float32frombits(binary.LittleEndian.Uint32(b[201:205]))
	p.VehicleClutch = math.Float32frombits(binary.LittleEndian.Uint32(b[205:209]))
	p.VehicleSteering = math.Float32frombits(binary.LittleEndian.Uint32(b[209:213]))
	p.VehicleHandbrake = math.Float32frombits(binary.LittleEndian.Uint32(b[213:217]))
	p.StageCurrentTime = math.Float32frombits(binary.LittleEndian.Uint32(b[217:221]))
	p.StageCurrentDistance = math.Float64frombits(binary.LittleEndian.Uint64(b[221:229]))
	p.StageLength = math.Float64frombits(binary.LittleEndian.Uint64(b[229:237]))
	return nil
}

func (p *PacketEASportsWRC) Steering() float32 {
	return p.VehicleSteering
}

func (p *PacketEASportsWRC) Throttle() float32 {
	return p.VehicleThrottle
}

func (p *PacketEASportsWRC) Brake() float32 {
	return p.VehicleBrake
}

func (p *PacketEASportsWRC) Clutch() float32 {
	return p.VehicleClutch
}

func (p *PacketEASportsWRC) Handbrake() float32 {
	return p.VehicleHandbrake
}

func (p *PacketEASportsWRC) Gear() int {
	return int(p.VehicleGearIndex)
}

func (p *PacketEASportsWRC) RPM() float32 {
	return p.VehicleEngineRpmCurrent
}

func (p *PacketEASportsWRC) MaxRPM() float32 {
	return p.VehicleEngineRpmMax
}

func (p *PacketEASportsWRC) Speed() float32 {
	return p.VehicleSpeed
}

func (p *PacketEASportsWRC) StageDistance() float32 {
	return float32(p.StageLength)
}
