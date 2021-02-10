package manager

import (
	"strings"
	"time"

	"github.com/feilb/Silvia/machine/modulator"
	"periph.io/x/periph/conn/gpio"
)

type ProcessController interface {
	Compute(s, i float64) float64
}

type Mode byte

const (
	ModeOff     Mode = 0
	ModeHeat    Mode = 1
	ModeBrew    Mode = 2
	ModeWater   Mode = 3
	ModeSteam   Mode = 4
	ModeInvalid Mode = 99
)

func (m Mode) String() string {
	switch m {
	case ModeOff:
		return "OFF"
	case ModeHeat:
		return "HEAT"
	case ModeBrew:
		return "BREW"
	case ModeWater:
		return "WATER"
	case ModeSteam:
		return "STEAM"
	default:
		return "Invalid"
	}
}

func ModeFromString(s string) Mode {
	switch strings.ToUpper(s) {
	case "OFF":
		return ModeOff
	case "HEAT":
		return ModeHeat
	case "BREW":
		return ModeBrew
	case "WATER":
		return ModeWater
	case "STEAM":
		return ModeSteam
	default:
		return ModeInvalid
	}
}

type Manager struct {
	TempUpdate    chan float64
	Tick          chan time.Time
	BrewSetpoint  float64
	SteamSetpoint float64
	Mode          Mode
	Controller    ProcessController

	BoilerPin, ValvePin, PumpPin gpio.PinIO

	setpoint         float64
	modulator        modulator.Modulator
	closer           chan time.Time
	autoOffEnabled   bool
	autoOffTimer     *time.Timer
	autoOffDurration time.Duration
	lastMode         Mode
}
