package ads1115

type Register byte

const (
	RegisterConversion Register = 0
	RegisterConfig     Register = 1
	RegisterLoThresh   Register = 2
	RegisterHiThresh   Register = 3
)

type InputMuxSetting byte

const (
	Mux0v1 InputMuxSetting = 0
	Mux0v3 InputMuxSetting = 1
	Mux1v3 InputMuxSetting = 2
	Mux2v3 InputMuxSetting = 3
	Mux0vG InputMuxSetting = 4
	Mux1vG InputMuxSetting = 5
	Mux2vG InputMuxSetting = 6
	Mux3vG InputMuxSetting = 7
)

func (i InputMuxSetting) String() string {
	switch i {
	case Mux0v1:
		return "AIN0 - AIN1"
	case Mux0v3:
		return "AIN0 - AIN3"
	case Mux1v3:
		return "AIN1 - AIN3"
	case Mux2v3:
		return "AIN2 - AIN3"
	case Mux0vG:
		return "AIN0 - GND"
	case Mux1vG:
		return "AIN1 - GND"
	case Mux2vG:
		return "AIN2 - GND"
	case Mux3vG:
		return "AIN3 - GND"
	default:
		return "Invalid"
	}
}

type FullScaleRange byte

const (
	Range6p144 FullScaleRange = 0
	Range4p096 FullScaleRange = 1
	Range2p048 FullScaleRange = 2
	Range1p024 FullScaleRange = 3
	Range0p512 FullScaleRange = 4
	Range0p256 FullScaleRange = 5
)

func (f FullScaleRange) String() string {
	switch f {
	case Range6p144:
		return "+/- 6.144V"
	case Range4p096:
		return "+/- 4.096V"
	case Range2p048:
		return "+/- 2.048V"
	case Range1p024:
		return "+/- 1.024V"
	case Range0p512:
		return "+/- 0.512V"
	case Range0p256:
		return "+/- 0.256V"
	default:
		return "Invalid"
	}
}

func (f FullScaleRange) ToFloat() float64 {
	switch f {
	case Range6p144:
		return 6.144
	case Range4p096:
		return 4.096
	case Range2p048:
		return 2.048
	case Range1p024:
		return 1.024
	case Range0p512:
		return 0.512
	case Range0p256:
		return 0.256
	default:
		return 0
	}
}

type OperatingMode byte

const (
	ModeContinuous OperatingMode = 0
	ModeSingle     OperatingMode = 1
)

func (o OperatingMode) String() string {
	switch o {
	case ModeContinuous:
		return "Continuous"
	case ModeSingle:
		return "Single"
	default:
		return "Invalid"
	}
}

type DataRate byte

const (
	Rate8SPS   DataRate = 0
	Rate16SPS  DataRate = 1
	Rate32SPS  DataRate = 2
	Rate64SPS  DataRate = 3
	Rate128SPS DataRate = 4
	Rate250SPS DataRate = 5
	Rate475SPS DataRate = 6
	Rate860SPS DataRate = 7
)

func (d DataRate) String() string {
	switch d {
	case Rate8SPS:
		return "8 SPS"
	case Rate16SPS:
		return "16 SPS"
	case Rate32SPS:
		return "32 SPS"
	case Rate64SPS:
		return "64 SPS"
	case Rate128SPS:
		return "128 SPS"
	case Rate250SPS:
		return "250 SPS"
	case Rate475SPS:
		return "475 SPS"
	case Rate860SPS:
		return "860 SPS"
	default:
		return "invalid"
	}
}

type ComparatorMode byte

const (
	ComparatorModeTraditional ComparatorMode = 0
	ComparatorModeWindow      ComparatorMode = 1
)

func (c ComparatorMode) String() string {
	switch c {
	case ComparatorModeTraditional:
		return "Traditional"
	case ComparatorModeWindow:
		return "Window"
	default:
		return "Invalid"
	}
}

type ComparatorPolarity byte

const (
	ComparatorPolarityActiveLow  ComparatorPolarity = 0
	ComparatorPolarityActiveHigh ComparatorPolarity = 1
)

func (c ComparatorPolarity) String() string {
	switch c {
	case ComparatorPolarityActiveLow:
		return "Active Low"
	case ComparatorPolarityActiveHigh:
		return "Active High"
	default:
		return "Invalid"
	}
}

type ComparatorLatching byte

const (
	ComparatorLatchingFalse ComparatorLatching = 0
	ComparatorLatchingTrue  ComparatorLatching = 1
)

func (c ComparatorLatching) String() string {
	switch c {
	case ComparatorLatchingFalse:
		return "Non-Latching"
	case ComparatorLatchingTrue:
		return "Latching"
	default:
		return "Invalid"
	}
}

type ComparatorQueue byte

const (
	ComparatorQueue1Sample ComparatorQueue = 0
	ComparatorQueue2Sample ComparatorQueue = 1
	ComparatorQueue3Sample ComparatorQueue = 2
	ComparatorDisable      ComparatorQueue = 3
)
