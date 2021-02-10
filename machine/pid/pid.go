package pid

import "fmt"

type PID struct {
	Kp, Ki, Kd, Kt float64
	OutMin, OutMax float64
	LastInput      []float64
	iTerm          float64
	outputError    float64
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

func (p *PID) Compute(setpoint, input float64) float64 {
	negativeD := true

	e := setpoint - input

	// only integrate if the output is within its control bounds
	if p.outputError == 0.0 {
		//p.iTerm += p.Ki * p.LastError[0] // + p.Kt*p.outputError
		p.iTerm += p.Ki * e
	}

	d := input - p.LastInput[0]
	p.LastInput = append(p.LastInput[1:], input)

	dTerm := p.Kd * d

	if negativeD && d > 0 {
		dTerm = 0
	}

	outputTarget := p.Kp*e + p.iTerm - dTerm

	output := limit(outputTarget, p.OutMin, p.OutMax)

	p.outputError = output - outputTarget

	if debug {
		fmt.Printf("i: %.4f, s: %.4f, e: %.4f, PT: %.4f, IT: %.4f, DT: %.4f, O: %.4f, OE: %.4f\n", input, setpoint, e, p.Kp*e, p.iTerm, -dTerm, output, p.outputError)
	}
	return output
}
