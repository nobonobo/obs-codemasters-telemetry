package codemasters

import (
	"encoding/binary"
	"fmt"
	"math"
)

const PacketDirtSeriesLength = 263

type PacketDirtSeries struct {
	Time                     float32
	LapTime                  float32
	LapDistance              float32
	TotalDistance            float32
	VehiclePosX              float32 // World space position
	VehiclePosY              float32 // World space position
	VehiclePosZ              float32 // World space position
	VehicleSpeed             float32
	VehicleVelX              float32 // Velocity in world space
	VehicleVelY              float32 // Velocity in world space
	VehicleVelZ              float32 // Velocity in world space
	VehicleRightDirectionX   float32 // World space right direction
	VehicleRightDirectionY   float32 // World space right direction
	VehicleRightDirectionZ   float32 // World space right direction
	VehicleForwardDirectionX float32 // World space forward direction
	VehicleForwardDirectionY float32 // World space forward direction
	VehicleForwardDirectionZ float32 // World space forward direction
	SuspPosBl                float32
	SuspPosBr                float32
	SuspPosFl                float32
	SuspPosFr                float32
	SuspVelBl                float32
	SuspVelBr                float32
	SuspVelFl                float32
	SuspVelFr                float32
	WheelSpeedBl             float32
	WheelSpeedBr             float32
	WheelSpeedFl             float32
	WheelSpeedFr             float32
	VehicleThrottle          float32
	VehicleSteering          float32
	VehicleBrake             float32
	VehicleClutch            float32
	VehicleGear              float32
	GforceLat                float32
	GforceLon                float32
	Lap                      float32
	EngineRate               float32
	//############################################################# unknown start
	SliProNativeSupport float32 // SLI Pro support
	CarPosition         float32 // car race position
	KersLevel           float32 // kers energy left
	KersMaxLevel        float32 // kers maximum energy
	Drs                 float32 // 0 = off, 1 = on
	TractionControl     float32 // 0 (off) - 2 (high)
	AntiLockBrakes      float32 // 0 (off) - 1 (on)
	FuelInTank          float32 // current fuel mass
	FuelCapacity        float32 // fuel capacity
	InPits              float32 // 0 = none, 1 = pitting, 2 = in pit area
	Sector              float32 // 0 = sector1, 1 = sector2 2 = sector3
	Sector1Time         float32 // time of sector1 (or 0)
	Sector2Time         float32 // time of sector2 (or 0)
	//############################################################# unknown end
	BrakesTemp [4]float32 // brakes temperature (centigrade)
	//############################################################# unknown start
	WheelsPressure [4]float32 // wheels pressure PSI
	TeamInfo       float32    // team ID
	//############################################################# unknown end
	TotalLaps   float32 // total number of laps in this race
	TrackSize   float32 // track size meters
	LastLapTime float32 // last lap time
	MaxRpm      float32 // cars max RPM, at which point the rev limiter will kick in
	//Idle_rpm               float32    // cars idle RPM
	//Max_gears              float32    // maximum number of gears
	//SessionType            float32    // 0 = unknown, 1 = practice, 2 = qualifying, 3 = race
	//DrsAllowed             float32    // 0 = not allowed, 1 = allowed, -1 = invalid / unknown
	//Track_number           float32    // -1 for unknown, 0-21 for tracks
	//VehicleFIAFlags        float32    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
}

func (p *PacketDirtSeries) UnmarshalBinary(b []byte) error {
	if len(b) < PacketDirtSeriesLength {
		return fmt.Errorf("invalid packet size: %d", len(b))
	}
	p.Time = math.Float32frombits(binary.LittleEndian.Uint32(b[:4]))
	p.LapTime = math.Float32frombits(binary.LittleEndian.Uint32(b[4:8]))
	p.LapDistance = math.Float32frombits(binary.LittleEndian.Uint32(b[8:12]))
	p.TotalDistance = math.Float32frombits(binary.LittleEndian.Uint32(b[12:16]))
	p.VehiclePosX = math.Float32frombits(binary.LittleEndian.Uint32(b[16:20]))
	p.VehiclePosY = math.Float32frombits(binary.LittleEndian.Uint32(b[20:24]))
	p.VehiclePosZ = math.Float32frombits(binary.LittleEndian.Uint32(b[24:28]))
	p.VehicleSpeed = math.Float32frombits(binary.LittleEndian.Uint32(b[28:32]))
	p.VehicleVelX = math.Float32frombits(binary.LittleEndian.Uint32(b[32:36]))
	p.VehicleVelY = math.Float32frombits(binary.LittleEndian.Uint32(b[36:40]))
	p.VehicleVelZ = math.Float32frombits(binary.LittleEndian.Uint32(b[40:44]))
	p.VehicleRightDirectionX = math.Float32frombits(binary.LittleEndian.Uint32(b[44:48]))
	p.VehicleRightDirectionY = math.Float32frombits(binary.LittleEndian.Uint32(b[48:52]))
	p.VehicleRightDirectionZ = math.Float32frombits(binary.LittleEndian.Uint32(b[52:56]))
	p.VehicleForwardDirectionX = math.Float32frombits(binary.LittleEndian.Uint32(b[56:60]))
	p.VehicleForwardDirectionY = math.Float32frombits(binary.LittleEndian.Uint32(b[60:64]))
	p.VehicleForwardDirectionZ = math.Float32frombits(binary.LittleEndian.Uint32(b[64:68]))
	p.SuspPosBl = math.Float32frombits(binary.LittleEndian.Uint32(b[68:72]))
	p.SuspPosBr = math.Float32frombits(binary.LittleEndian.Uint32(b[72:76]))
	p.SuspPosFl = math.Float32frombits(binary.LittleEndian.Uint32(b[76:80]))
	p.SuspPosFr = math.Float32frombits(binary.LittleEndian.Uint32(b[80:84]))
	p.SuspVelBl = math.Float32frombits(binary.LittleEndian.Uint32(b[84:88]))
	p.SuspVelBr = math.Float32frombits(binary.LittleEndian.Uint32(b[88:92]))
	p.SuspVelFl = math.Float32frombits(binary.LittleEndian.Uint32(b[92:96]))
	p.SuspVelFr = math.Float32frombits(binary.LittleEndian.Uint32(b[96:100]))
	p.WheelSpeedBl = math.Float32frombits(binary.LittleEndian.Uint32(b[100:104]))
	p.WheelSpeedBr = math.Float32frombits(binary.LittleEndian.Uint32(b[104:108]))
	p.WheelSpeedFl = math.Float32frombits(binary.LittleEndian.Uint32(b[108:112]))
	p.WheelSpeedFr = math.Float32frombits(binary.LittleEndian.Uint32(b[112:116]))
	p.VehicleThrottle = math.Float32frombits(binary.LittleEndian.Uint32(b[116:120]))
	p.VehicleSteering = math.Float32frombits(binary.LittleEndian.Uint32(b[120:124]))
	p.VehicleBrake = math.Float32frombits(binary.LittleEndian.Uint32(b[124:128]))
	p.VehicleClutch = math.Float32frombits(binary.LittleEndian.Uint32(b[128:132]))
	p.VehicleGear = math.Float32frombits(binary.LittleEndian.Uint32(b[132:136]))
	p.GforceLat = math.Float32frombits(binary.LittleEndian.Uint32(b[136:140]))
	p.GforceLon = math.Float32frombits(binary.LittleEndian.Uint32(b[140:144]))
	p.Lap = math.Float32frombits(binary.LittleEndian.Uint32(b[144:148]))
	p.EngineRate = math.Float32frombits(binary.LittleEndian.Uint32(b[148:152]))
	p.SliProNativeSupport = math.Float32frombits(binary.LittleEndian.Uint32(b[152:156]))
	p.CarPosition = math.Float32frombits(binary.LittleEndian.Uint32(b[156:160]))
	p.KersLevel = math.Float32frombits(binary.LittleEndian.Uint32(b[160:164]))
	p.KersMaxLevel = math.Float32frombits(binary.LittleEndian.Uint32(b[164:168]))
	p.Drs = math.Float32frombits(binary.LittleEndian.Uint32(b[168:172]))
	p.TractionControl = math.Float32frombits(binary.LittleEndian.Uint32(b[172:176]))
	p.AntiLockBrakes = math.Float32frombits(binary.LittleEndian.Uint32(b[176:180]))
	p.FuelInTank = math.Float32frombits(binary.LittleEndian.Uint32(b[180:184]))
	p.FuelCapacity = math.Float32frombits(binary.LittleEndian.Uint32(b[184:188]))
	p.InPits = math.Float32frombits(binary.LittleEndian.Uint32(b[188:192]))
	p.Sector = math.Float32frombits(binary.LittleEndian.Uint32(b[192:196]))
	p.Sector1Time = math.Float32frombits(binary.LittleEndian.Uint32(b[196:200]))
	p.Sector2Time = math.Float32frombits(binary.LittleEndian.Uint32(b[200:204]))
	p.BrakesTemp[0] = math.Float32frombits(binary.LittleEndian.Uint32(b[204:208]))
	p.BrakesTemp[1] = math.Float32frombits(binary.LittleEndian.Uint32(b[208:212]))
	p.BrakesTemp[2] = math.Float32frombits(binary.LittleEndian.Uint32(b[212:216]))
	p.BrakesTemp[3] = math.Float32frombits(binary.LittleEndian.Uint32(b[216:220]))
	p.WheelsPressure[0] = math.Float32frombits(binary.LittleEndian.Uint32(b[220:224]))
	p.WheelsPressure[1] = math.Float32frombits(binary.LittleEndian.Uint32(b[224:228]))
	p.WheelsPressure[2] = math.Float32frombits(binary.LittleEndian.Uint32(b[228:232]))
	p.WheelsPressure[3] = math.Float32frombits(binary.LittleEndian.Uint32(b[232:236]))
	p.TeamInfo = math.Float32frombits(binary.LittleEndian.Uint32(b[236:240]))
	p.TotalLaps = math.Float32frombits(binary.LittleEndian.Uint32(b[240:244]))
	p.TrackSize = math.Float32frombits(binary.LittleEndian.Uint32(b[244:248]))
	p.LastLapTime = math.Float32frombits(binary.LittleEndian.Uint32(b[248:252]))
	p.MaxRpm = math.Float32frombits(binary.LittleEndian.Uint32(b[252:256]))
	return nil
}

func (p *PacketDirtSeries) Steering() float32 {
	return p.VehicleSteering
}

func (p *PacketDirtSeries) Throttle() float32 {
	return p.VehicleThrottle
}

func (p *PacketDirtSeries) Brake() float32 {
	return p.VehicleBrake
}

func (p *PacketDirtSeries) Clutch() float32 {
	return p.VehicleClutch
}

func (p *PacketDirtSeries) Handbrake() float32 {
	return 0
}

func (p *PacketDirtSeries) Gear() int {
	return int(p.VehicleGear)
}

func (p *PacketDirtSeries) RPM() float32 {
	return p.EngineRate
}

func (p *PacketDirtSeries) MaxRPM() float32 {
	return p.MaxRpm
}

func (p *PacketDirtSeries) Speed() float32 {
	return p.VehicleSpeed
}

func (p *PacketDirtSeries) StageDistance() float32 {
	return p.TrackSize
}
