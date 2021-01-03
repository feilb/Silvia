package modulator

import (
	"fmt"
	"testing"
)

func TestModulator_Modulate(t *testing.T) {

	m := Modulator{Setpoint: 0.0}
	for i := 0; i < 20; i++ {
		fmt.Printf("%v\n", m.Modulate())
	}

	t.Error("err")
}
