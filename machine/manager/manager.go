package manager

import (
	"fmt"
	"time"

	"periph.io/x/periph/conn/gpio"
)

func (m *Manager) manageMode() {
	if m.lastMode == m.Mode {
		return
	}

	switch m.Mode {
	case ModeOff:
		m.PumpPin.Out(gpio.Low)
		m.ValvePin.Out(gpio.Low)
		m.BoilerPin.Out(gpio.Low)
		m.modulator.Setpoint = 0
		m.autoOffTimer.Stop()
	case ModeHeat:
		m.PumpPin.Out(gpio.Low)
		m.ValvePin.Out(gpio.Low)
		m.setpoint = m.BrewSetpoint
		m.resetTimer()
	case ModeBrew:
		m.PumpPin.Out(gpio.High)
		m.ValvePin.Out(gpio.High)
		m.setpoint = m.BrewSetpoint
		m.resetTimer()
	case ModeWater:
		m.PumpPin.Out(gpio.High)
		m.ValvePin.Out(gpio.Low)
		m.setpoint = m.BrewSetpoint
		m.resetTimer()
	case ModeSteam:
		m.PumpPin.Out(gpio.Low)
		m.ValvePin.Out(gpio.Low)
		m.setpoint = m.SteamSetpoint
		m.resetTimer()
	}

	m.lastMode = m.Mode
}

func (m *Manager) resetTimer() {
	if m.Mode != ModeOff {
		fmt.Printf("Resetting timer Duration %v", m.autoOffDurration)
		m.autoOffTimer.Stop()
		m.autoOffTimer.Reset(m.autoOffDurration)
	}
}

func (m *Manager) CurrentSetpoint() float64 {
	return m.setpoint
}

func (m *Manager) EnableAutoOff() {
	m.autoOffEnabled = true
	m.resetTimer()
}

func (m *Manager) DisableAutoOff() {
	m.autoOffEnabled = false
	m.autoOffTimer.Stop()
}

func (m *Manager) AutoOffEnabled() bool {
	return m.autoOffEnabled
}

func (m *Manager) SetAutoOffDuration(d time.Duration) {
	m.autoOffDurration = d
	m.resetTimer()
}

func (m *Manager) AutoOffDuration() time.Duration {
	return m.autoOffDurration
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

	m.lastMode = m.Mode

	fmt.Println("timer init")
	m.autoOffTimer = time.NewTimer(m.autoOffDurration)
	m.autoOffTimer.Stop()

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

	// AutoOff
	go func() {
		for range m.autoOffTimer.C {
			if m.autoOffEnabled {
				fmt.Println("autooff")
				m.Mode = ModeOff
			}
		}
	}()

	<-m.closer
}

func (m *Manager) Close() {
	m.closer <- time.Now()
	m.autoOffTimer.Stop()
	m.BoilerPin.Out(gpio.Low)
	m.ValvePin.Out(gpio.Low)
	m.PumpPin.Out(gpio.Low)
}
