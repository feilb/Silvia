package main

import (
	"fmt"

	"github.com/feilb/Silvia/chips/MCP9600"
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
	mcp := MCP9600.Dev{Device: &i2c.Dev{Addr: 0x60, Bus: i2cBus}}

	mcp.SetThermocoupleType(MCP9600.TypeK)
	tcType, _ := mcp.GetThermocoupleType()
	fmt.Printf("Type: %v\n", tcType)

	mcp.SetFilterCoefficient(MCP9600.FilterCoefficient0)
	coef, _ := mcp.GetFilterCoefficient()
	fmt.Printf("Coef: %v\n", coef)

	mcp.SetADCResolution(MCP9600.ADCResolution18Bit)
	adc, _ := mcp.GetADCResolution()
	fmt.Printf("ADCR: %v\n", adc)

	mcp.SetColdJunctionResolution(MCP9600.ColdJuncitonResolution0p25)
	cjr, _ := mcp.GetColdJunctionResolution()
	fmt.Printf("CJR: %v\n", cjr)

	mcp.SetBurstModeSamples(MCP9600.BurstMode2Samples)
	sam, _ := mcp.GetBurstModeSamples()
	fmt.Printf("BMS: %v\n", sam)

	mcp.SetShutdownMode(MCP9600.ModeNormal)
	sdm, _ := mcp.GetShutdownMode()
	fmt.Printf("SDM: %v\n", sdm)

	temp, _ := mcp.GetTemp()
	fmt.Printf("temp: %v\n", temp)
	ctemp, _ := mcp.GetAmbientTemp()
	fmt.Printf("tamb: %v\n", ctemp)
}
