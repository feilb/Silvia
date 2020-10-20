package mcp9600

type Register byte

const (
	RegisterTHot         Register = 0
	RegisterTDelta                = 1
	RegisterTCold                 = 2
	RegisterRawADC                = 3
	RegisterStatus                = 4
	RegisterSensorConfig          = 5
	RegisterDeviceConfig          = 6
)

type ADCResolution byte

const (
	ADCResolution18Bit ADCResolution = 0
	ADCResolution16Bit               = 1
	ADCResolution14Bit               = 2
	ADCResolution12Bit               = 3
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
	TypeJ                  = 1
	TypeT                  = 2
	TypeN                  = 3
	TypeS                  = 4
	TypeE                  = 5
	TypeB                  = 6
	TypeR                  = 7
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
	ColdJuncitonResolution0p0625 = 0
	ColdJuncitonResolution0p25   = 1
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
	FilterCoefficient1                   = 1
	FilterCoefficient2                   = 2
	FilterCoefficient3                   = 3
	FilterCoefficient4                   = 4
	FilterCoefficient5                   = 5
	FilterCoefficient6                   = 6
	FilterCoefficient7                   = 7
)

type BurstModeSamples byte

const (
	BurstMode1Sample    BurstModeSamples = 0
	BurstMode2Samples                    = 1
	BurstMode4Samples                    = 2
	BurstMode8Samples                    = 3
	BurstMode16Samples                   = 4
	BurstMode32Samples                   = 5
	BurstMode64Samples                   = 6
	BurstMode128Samples                  = 7
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
	ModeShutdown              = 1
	ModeBurst                 = 2
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
