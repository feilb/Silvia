package main

import (
	"fmt"
	"time"

	mcp9600 "github.com/feilb/Silvia/chips/mcp9600"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

func main() {
	host.Init()

	i2cBus, err := i2creg.Open("")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(i2cBus)
	mcp := mcp9600.Dev{Device: &i2c.Dev{Addr: 0x60, Bus: i2cBus}}

	mcp.SetThermocoupleType(mcp9600.TypeK)
	tcType, _ := mcp.GetThermocoupleType()
	fmt.Printf("Type: %v\n", tcType)

	mcp.SetFilterCoefficient(mcp9600.FilterCoefficient2)
	coef, _ := mcp.GetFilterCoefficient()
	fmt.Printf("Coef: %v\n", coef)

	mcp.SetADCResolution(mcp9600.ADCResolution18Bit)
	adc, _ := mcp.GetADCResolution()
	fmt.Printf("ADCR: %v\n", adc)

	mcp.SetColdJunctionResolution(mcp9600.ColdJuncitonResolution0p0625)
	cjr, _ := mcp.GetColdJunctionResolution()
	fmt.Printf("CJR: %v\n", cjr)

	mcp.SetBurstModeSamples(mcp9600.BurstMode2Samples)
	sam, _ := mcp.GetBurstModeSamples()
	fmt.Printf("BMS: %v\n", sam)

	mcp.SetShutdownMode(mcp9600.ModeNormal)
	sdm, _ := mcp.GetShutdownMode()
	fmt.Printf("SDM: %v\n", sdm)

	ctemp, _ := mcp.GetAmbientTemp()
	fmt.Printf("tamb: %v\n", ctemp)

	for {
		temp, _ := mcp.GetTemp()
		fmt.Printf("temp: %v\n", temp)
		time.Sleep(time.Second)
	}

}
