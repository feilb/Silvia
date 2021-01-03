package mcp9600

type Register byte

const (
	RegisterTHot         Register = 0
	RegisterTDelta       Register = 1
	RegisterTCold        Register = 2
	RegisterRawADC       Register = 3
	RegisterStatus       Register = 4
	RegisterSensorConfig Register = 5
	RegisterDeviceConfig Register = 6
)

type ADCResolution byte

const (
	ADCResolution18Bit ADCResolution = 0
	ADCResolution16Bit ADCResolution = 1
	ADCResolution14Bit ADCResolution = 2
	ADCResolution12Bit ADCResolution = 3
)

func (adr ADCResolution) String() string {
	switch adr {
	case ADCResolution18Bit:
		return "18 bit"
	case ADCResolution16Bit:
		return "16 bit"
	case ADCResolution14Bit:
		return "14 bit"
	case ADCResolution12Bit:
		return "12 bit"
	default:
		return "Invalid"
	}
}

type ThermocoupleType byte

const (
	TypeK ThermocoupleType = 0
	TypeJ ThermocoupleType = 1
	TypeT ThermocoupleType = 2
	TypeN ThermocoupleType = 3
	TypeS ThermocoupleType = 4
	TypeE ThermocoupleType = 5
	TypeB ThermocoupleType = 6
	TypeR ThermocoupleType = 7
)

func (tt ThermocoupleType) String() string {
	switch tt {
	case TypeK:
		return "K"
	case TypeJ:
		return "J"
	case TypeT:
		return "T"
	case TypeN:
		return "N"
	case TypeS:
		return "S"
	case TypeE:
		return "E"
	case TypeB:
		return "B"
	case TypeR:
		return "R"
	default:
		return "Invalid"
	}
}

type ColdJuncitonResolution byte

const (
	ColdJuncitonResolution0p0625 ColdJuncitonResolution = 0
	ColdJuncitonResolution0p25   ColdJuncitonResolution = 1
)

func (cjr ColdJuncitonResolution) String() string {
	switch cjr {
	case ColdJuncitonResolution0p0625:
		return "0.0625 C"
	case ColdJuncitonResolution0p25:
		return "0.25 C"
	default:
		return "Invalid"
	}
}

type FilterCoefficient byte

const (
	FilterCoefficient0 FilterCoefficient = 0
	FilterCoefficient1 FilterCoefficient = 1
	FilterCoefficient2 FilterCoefficient = 2
	FilterCoefficient3 FilterCoefficient = 3
	FilterCoefficient4 FilterCoefficient = 4
	FilterCoefficient5 FilterCoefficient = 5
	FilterCoefficient6 FilterCoefficient = 6
	FilterCoefficient7 FilterCoefficient = 7
)

type BurstModeSamples byte

const (
	BurstMode1Sample    BurstModeSamples = 0
	BurstMode2Samples   BurstModeSamples = 1
	BurstMode4Samples   BurstModeSamples = 2
	BurstMode8Samples   BurstModeSamples = 3
	BurstMode16Samples  BurstModeSamples = 4
	BurstMode32Samples  BurstModeSamples = 5
	BurstMode64Samples  BurstModeSamples = 6
	BurstMode128Samples BurstModeSamples = 7
)

func (bm BurstModeSamples) String() string {
	switch bm {
	case BurstMode1Sample:
		return "1 Sample"
	case BurstMode2Samples:
		return "2 Samples"
	case BurstMode4Samples:
		return "4 Samples"
	case BurstMode8Samples:
		return "8 Samples"
	case BurstMode16Samples:
		return "16 Samples"
	case BurstMode32Samples:
		return "32 Samples"
	case BurstMode64Samples:
		return "64 Samples"
	case BurstMode128Samples:
		return "128 Samples"
	default:
		return "Invalid"
	}
}

type ShutdownMode byte

const (
	ModeNormal   ShutdownMode = 0
	ModeShutdown ShutdownMode = 1
	ModeBurst    ShutdownMode = 2
)

func (sm ShutdownMode) String() string {
	switch sm {
	case ModeNormal:
		return "Normal"
	case ModeShutdown:
		return "Shutdown"
	case ModeBurst:
		return "Burst"
	default:
		return "Invalid"
	}
}
