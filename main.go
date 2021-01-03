package main

import (
	"fmt"
	"time"

	"github.com/feilb/Silvia/chips/ads1115"
	"github.com/feilb/Silvia/chips/mcp9600"
	"github.com/feilb/Silvia/machine/modulator"
	"github.com/feilb/Silvia/utils/i2cUtils"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

func main() {
	// **************************************************
	// periph.io Init
	// **************************************************
	host.Init()

	// I2C
	i2cBus, err := i2creg.Open("")

	if err != nil {
		fmt.Println(err)
		return
	}

	//GPIO
	boilerPin := gpioreg.ByName("GPIO19")
	boilerPin.Out(gpio.Low)

	valvePin := gpioreg.ByName("GPIO13")
	valvePin.Out(gpio.Low)

	pumpPin := gpioreg.ByName("GPIO26")
	pumpPin.Out(gpio.Low)

	zeroCrossingPin := gpioreg.ByName("GPIO24")
	zeroCrossingPin.In(gpio.PullUp, gpio.RisingEdge)

	// **************************************************
	// MCP9600 Thermocouple Sensor Init
	// **************************************************
	mcp := mcp9600.Dev{Device: &i2c.Dev{Addr: 0x60, Bus: i2cBus}}

	mcp.SetThermocoupleType(mcp9600.TypeK)
	mcp.SetFilterCoefficient(mcp9600.FilterCoefficient4)
	mcp.SetADCResolution(mcp9600.ADCResolution16Bit)
	mcp.SetColdJunctionResolution(mcp9600.ColdJuncitonResolution0p25)

	// **************************************************
	// ADS1115 ADC Init
	// **************************************************
	adc := ads1115.Dev{Device: &i2c.Dev{Addr: 0x48, Bus: i2cBus}}
	adcRange := ads1115.Range4p096

	adc.SetRange(adcRange)
	adc.SetMux(ads1115.Mux0v3)
	adc.SetMode(ads1115.ModeSingle)

	r, _ := adc.GetRange()
	d, _ := adc.GetDataRate()
	m, _ := adc.GetMux()
	o, _ := adc.GetMode()
	fmt.Printf("range: %v\n", r)
	fmt.Printf("rate: %v\n", d)
	fmt.Printf("mux: %v\n", m)
	fmt.Printf("mode: %v\n", o)

	cfg, _ := i2cUtils.ReadI2C(adc.Device, byte(ads1115.RegisterConfig), 2)
	fmt.Printf("raw: %08b\n", cfg)

	// **************************************************
	// PID Init
	// **************************************************
	/*pid := pid.PID{
		Kp:       0.0375,
		Ki:       0.0002727,
		Kd:       1.289,
		OutMax:   1,
		OutMin:   0,
		Setpoint: 40,
	}*/

	// **************************************************
	// START
	// **************************************************

	start := time.Now()

	fmt.Println("Modulator output = 0")
	mod := modulator.Modulator{Setpoint: 0.00}

	go func() {
		fmt.Println("zero crossing function")
		for {
			zeroCrossingPin.WaitForEdge(time.Second)
			if mod.Modulate() == 1 {
				boilerPin.Out(gpio.High)
			} else {
				boilerPin.Out(gpio.Low)
			}
		}
	}()

	defer func() {
		fmt.Println("boiler low func")
		boilerPin.Out(gpio.Low)
	}()

	var temp float64

	go func() {
		fmt.Println("temp reading fn")
		for {
			temp, _ = mcp.GetTemp()

			//mod.Setpoint = pid.Compute(temp)
			fmt.Printf("%v,%v,%v\n", time.Since(start).Milliseconds(), temp, mod.Setpoint)
			time.Sleep(time.Second * 5)
		}

	}()

	/*
		fmt.Println("Waiting 10s...")
		time.Sleep(time.Second * 10)
		mod.Setpoint = 0.05
		time.Sleep(time.Hour * 6)
	*/

	pumpPin.Out(gpio.High)
	time.Sleep(time.Millisecond * 500)
	pumpPin.Out(gpio.Low)
}
