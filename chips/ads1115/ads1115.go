package ads1115

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
func convertFS(t []byte) float64 {
	return float64(int16(binary.BigEndian.Uint16(t))) / 32768.0
}

// **************************************************
// Multiplexer
// **************************************************

// SetMux - Set conversion multiplexer configuration
func (d *Dev) SetMux(v InputMuxSetting) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 14, 12, byte(v))
}

// GetMux - Get conversion multiplexer configuration
func (d *Dev) GetMux() (InputMuxSetting, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 14, 12)

	return InputMuxSetting(raw), err
}

// **************************************************
// Range
// **************************************************

// SetRange - Set programmable gain amplifier config
func (d *Dev) SetRange(v FullScaleRange) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 11, 9, byte(v))
}

// GetRange - Set programmable gain amplifier config
func (d *Dev) GetRange() (FullScaleRange, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 11, 9)

	return FullScaleRange(raw), err
}

// **************************************************
// Operating Mode
// **************************************************

// SetMode - Set operating mode
func (d *Dev) SetMode(v OperatingMode) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 8, 8, byte(v))
}

// GetMode - Get operating mode
func (d *Dev) GetMode() (OperatingMode, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 8, 8)

	return OperatingMode(raw), err
}

// **************************************************
// Data Rate
// **************************************************

// SetDataRate -
func (d *Dev) SetDataRate(v DataRate) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 7, 5, byte(v))
}

// GetDataRate -
func (d *Dev) GetDataRate() (DataRate, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 7, 5)

	return DataRate(raw), err
}

// **************************************************
// Comparator Mode
// **************************************************

// SetComparatorMode -
func (d *Dev) SetComparatorMode(v ComparatorMode) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 4, 4, byte(v))
}

// GetComparatorMode -
func (d *Dev) GetComparatorMode() (ComparatorMode, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 4, 4)

	return ComparatorMode(raw), err
}

// **************************************************
// Comparator Polarity
// **************************************************

// SetComparatorPolarity -
func (d *Dev) SetComparatorPolarity(v ComparatorPolarity) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 3, 3, byte(v))
}

// GetComparatorPolarity -
func (d *Dev) GetComparatorPolarity() (ComparatorPolarity, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 3, 3)

	return ComparatorPolarity(raw), err
}

// **************************************************
// Comparator Latch
// **************************************************

// SetComparatorLatch -
func (d *Dev) SetComparatorLatch(v ComparatorLatching) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 2, 2, byte(v))
}

// GetComparatorLatch -
func (d *Dev) GetComparatorLatch() (ComparatorLatching, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 2, 2)

	return ComparatorLatching(raw), err
}

// **************************************************
// Comparator Queue
// **************************************************

// SetComparatorQueue -
func (d *Dev) SetComparatorQueue(v ComparatorQueue) error {
	return i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 1, 0, byte(v))
}

// GetComparatorQueue -
func (d *Dev) GetComparatorQueue() (ComparatorQueue, error) {
	raw, err := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 1, 0)

	return ComparatorQueue(raw), err
}

// **************************************************
// Get Conversion
// **************************************************

// GetNormalizedConversion - return ranges between 1 and -1 of full-scale value
func (d *Dev) GetNormalizedConversion() (float64, error) {
	raw, err := i2cUtils.ReadI2C(d.Device, byte(RegisterConversion), 2)

	return convertFS(raw), err
}

// GetScaledConversion - return current scaled value given fullscalerange
func (d *Dev) GetScaledConversion(r FullScaleRange) (float64, error) {
	raw, err := d.GetNormalizedConversion()

	return raw * r.ToFloat(), err
}

// GetNormalizedSingle - In single mode, trigger, wait for, then return a conversion
func (d *Dev) GetNormalizedSingle() (float64, error) {
	i2cUtils.SetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 15, 15, 1)

	for {
		val, _ := i2cUtils.GetRegisterBitsI2C(d.Device, byte(RegisterConfig), 2, 15, 15)

		if val == 1 {
			break
		}
	}

	return d.GetNormalizedConversion()
}

// GetScaledSingle - Same as GetNormalizedSingle, but scaled given fullscalerange
func (d *Dev) GetScaledSingle(r FullScaleRange) (float64, error) {
	raw, err := d.GetNormalizedSingle()

	return raw * r.ToFloat(), err
}

// **************************************************
// Low Threshold
// **************************************************

// GetLowThreshold -
func (d *Dev) GetLowThreshold() (int16, error) {
	raw, err := i2cUtils.ReadI2C(d.Device, byte(RegisterLoThresh), 2)

	return int16(binary.BigEndian.Uint16(raw)), err
}

// SetLowThreshold -
func (d *Dev) SetLowThreshold(v int16) error {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(v))

	return i2cUtils.WriteI2C(d.Device, byte(RegisterLoThresh), b)
}

// **************************************************
// High Threshold
// **************************************************

// GetHighThreshold -
func (d *Dev) GetHighThreshold() (int16, error) {
	raw, err := i2cUtils.ReadI2C(d.Device, byte(RegisterHiThresh), 2)

	return int16(binary.BigEndian.Uint16(raw)), err
}

// SetHighThreshold -
func (d *Dev) SetHighThreshold(v int16) error {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(v))

	return i2cUtils.WriteI2C(d.Device, byte(RegisterHiThresh), b)
}
