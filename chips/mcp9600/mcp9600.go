package mcp9600

import (
	"encoding/binary"

	"github.com/feilb/Silvia/utils/i2cUtils"
	"periph.io/x/periph/conn/i2c"
)

// Dev - i2c.Dev associated with mcp9600
type Dev struct {
	Device *i2c.Dev
}

// **************************************************
// Boilerplate reducing functions
// **************************************************
func convertTemp(t []byte) float64 {
	return float64(int16(binary.BigEndian.Uint16(t))) / 16.0
}

// **************************************************
// Thermocouple Type
// **************************************************

// SetThermocoupleType - Set thermocouple type
func (d *Dev) SetThermocoupleType(t ThermocoupleType) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterSensorConfig), 1, 6, 4, byte(t))
}

// GetThermocoupleType - Get thermocouple type
func (d *Dev) GetThermocoupleType() (ThermocoupleType, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterSensorConfig), 1, 6, 4)

	return ThermocoupleType(raw), err
}

// **************************************************
// Filter Coefficients
// **************************************************

// SetFilterCoefficient - Set coefficient used in temperature filter
func (d *Dev) SetFilterCoefficient(f FilterCoefficient) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterSensorConfig), 1, 2, 0, byte(f))
}

// GetFilterCoefficient - Get coefficient used in temperature filter
func (d *Dev) GetFilterCoefficient() (FilterCoefficient, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterSensorConfig), 1, 2, 0)

	return FilterCoefficient(raw), err
}

// **************************************************
// Cold Junction Resolution
// **************************************************

// SetColdJunctionResolution - Set ADC resolution used when converting
// cold-junction temperature
func (d *Dev) SetColdJunctionResolution(r ColdJuncitonResolution) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterDeviceConfig), 1, 7, 7, byte(r))
}

// GetColdJunctionResolution - Get ADC resolution used when converting
// cold-junction temperature
func (d *Dev) GetColdJunctionResolution() (ColdJuncitonResolution, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterDeviceConfig), 1, 7, 7)

	return ColdJuncitonResolution(raw), err
}

// **************************************************
// ADC Resolution
// **************************************************

// SetADCResolution - Set ADC resolution used in converting
// hot junction temperature
func (d *Dev) SetADCResolution(r ADCResolution) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterDeviceConfig), 1, 6, 5, byte(r))
}

// GetADCResolution - Return ADC resolution used in converting
// hot junction temperature
func (d *Dev) GetADCResolution() (ADCResolution, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterDeviceConfig), 1, 6, 5)

	return ADCResolution(raw), err
}

// **************************************************
// Burst Mode Samples
// **************************************************

// SetBurstModeSamples - Set number of samples to be captured in
// burst mode
func (d *Dev) SetBurstModeSamples(s BurstModeSamples) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterDeviceConfig), 1, 4, 2, byte(s))
}

// GetBurstModeSamples - Return number of samples that will be captured
// in burst mode
func (d *Dev) GetBurstModeSamples() (BurstModeSamples, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterDeviceConfig), 1, 4, 2)

	return BurstModeSamples(raw), err
}

// **************************************************
// Shutdown Mode
// **************************************************

// SetShutdownMode - Sets mcp9600 operating mode
func (d *Dev) SetShutdownMode(s ShutdownMode) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterDeviceConfig), 1, 1, 0, byte(s))
}

// GetShutdownMode - Returns current mcp9600 operating mode
func (d *Dev) GetShutdownMode() (ShutdownMode, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterDeviceConfig), 1, 1, 0)

	return ShutdownMode(raw), err
}

// **************************************************
// Temperatures
// **************************************************

// GetTemp returns cold-junction compensated thermocouple temperature
func (d *Dev) GetTemp() (float64, error) {
	raw, err := i2cUtils.ReadI2C(d.Device, byte(RegisterTHot), 2)

	if err != nil {
		return 0, err
	}

	return convertTemp(raw), nil
}

// GetAmbientTemp returns cold junction temp
func (d *Dev) GetAmbientTemp() (float64, error) {
	raw, err := i2cUtils.ReadI2C(d.Device, byte(RegisterTCold), 2)

	if err != nil {
		return 0, err
	}

	return convertTemp(raw), nil
}
