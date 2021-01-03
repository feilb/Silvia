package pid

import "fmt"

type PID struct {
	Kp, Ki, Kd     float64
	Setpoint       float64
	OutMin, OutMax float64
	lastInput      float64
	iTerm          float64
}

const debug = true

func limit(a, min, max float64) float64 {
	switch {
	case a < min:
		return min
	case a > max:
		return max
	default:
		return a
	}
}

func (p *PID) Compute(input float64) float64 {
	e := p.Setpoint - input
	p.iTerm += p.Ki * e

	p.iTerm = limit(p.iTerm, p.OutMin, p.OutMax)

	d := input - p.lastInput
	p.lastInput = input

	output := p.Kp*e + p.iTerm - p.Kd*d
	output = limit(output, p.OutMin, p.OutMax)

	if debug {
		fmt.Printf("i: %.4f, s: %.4f, e: %.4f, i: %.4f, d: %.4f, PT: %.4f, IT: %.4f, DT: %.4f, O: %.4f\n", input, p.Setpoint, e, p.iTerm, d, p.Kp*e, p.iTerm, p.Kd*d, output)
	}
	return output
}
