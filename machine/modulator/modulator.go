package modulator

type Modulator struct {
	err   float64
	state int

	Setpoint float64
}

func (m *Modulator) Modulate() int {
	m.err += float64(m.state) - m.Setpoint

	if m.err < 0 {
		m.state = 1
	} else {
		m.state = 0
	}

	return m.state
}
