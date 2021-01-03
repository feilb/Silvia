package manager

import (
	"time"

	"periph.io/x/periph/conn/gpio"
)

func (m *Manager) manageMode() {
	switch m.Mode {
	case ModeOff:
		m.PumpPin.Out(gpio.Low)
		m.ValvePin.Out(gpio.Low)
		m.BoilerPin.Out(gpio.Low)
		m.modulator.Setpoint = 0
	case ModeHeat:
		m.PumpPin.Out(gpio.Low)
		m.ValvePin.Out(gpio.Low)
		m.setpoint = m.BrewSetpoint
	case ModeBrew:
		m.PumpPin.Out(gpio.High)
		m.ValvePin.Out(gpio.High)
		m.setpoint = m.BrewSetpoint
	case ModeWater:
		m.PumpPin.Out(gpio.High)
		m.ValvePin.Out(gpio.Low)
		m.setpoint = m.BrewSetpoint
	case ModeSteam:
		m.PumpPin.Out(gpio.Low)
		m.ValvePin.Out(gpio.Low)
		m.setpoint = m.SteamSetpoint
	}
}

func (m *Manager) CurrentSetpoint() float64 {
	return m.setpoint
}

func (m *Manager) Run() {
	if m.closer == nil {
		m.closer = make(chan time.Time)
	}
	if m.Tick == nil {
		m.Tick = make(chan time.Time)
	}
	if m.TempUpdate == nil {
		m.TempUpdate = make(chan float64)
	}

	// Tick
	go func() {

		for range m.Tick {
			m.manageMode()

			if m.modulator.Modulate() == 1 {
				m.BoilerPin.Out(gpio.High)
			} else {
				m.BoilerPin.Out(gpio.Low)
			}
		}
	}()

	// Temp
	go func() {
		for temp := range m.TempUpdate {
			if m.Mode != ModeOff && m.Mode != ModeInvalid {
				m.modulator.Setpoint = m.Controller.Compute(m.setpoint, temp)
			}
		}
	}()

	<-m.closer
}

func (m *Manager) Close() {
	m.closer <- time.Now()
	m.BoilerPin.Out(gpio.Low)
	m.ValvePin.Out(gpio.Low)
	m.PumpPin.Out(gpio.Low)
}
