package i2cUtils

import (
	"errors"
	"fmt"

	"github.com/feilb/Silvia/utils/bitUtils"
	"periph.io/x/periph/conn/i2c"
)

func ReadI2C(d *i2c.Dev, reg byte, count int) ([]byte, error) {
	pkt := make([]byte, count)

	err := d.Tx([]byte{reg}, pkt)

	if err != nil {
		return nil, err
	}

	return pkt, nil
}

func WriteI2C(d *i2c.Dev, reg byte, msg []byte) error {
	pkt := append([]byte{reg}, msg...)

	count, err := d.Write(pkt)

	if err != nil {
		return err
	}
	if count < len(msg)+1 {
		return fmt.Errorf("Expected to send %v bytes but sent %v", len(msg)+1, count)
	}

	return nil
}

func SetRegisterBitsI2C(d *i2c.Dev, reg byte, bytes int, highBit, lowBit, val byte) error {
	raw, err := ReadI2C(d, reg, bytes)

	if err != nil {
		return err
	}

	if highBit/8 != lowBit/8 {
		return errors.New("setRegister does not support breaking values across byte boundaries")
	}

	i := byte(bytes) - highBit/8 - 1

	bitUtils.SetBits(&raw[i], highBit%8, lowBit%8, byte(val))

	return WriteI2C(d, reg, raw)
}

func GetRegisterBitsI2C(d *i2c.Dev, reg byte, bytes int, highBit, lowBit byte) (byte, error) {
	raw, err := ReadI2C(d, reg, bytes)

	if err != nil {
		return 0, err
	}

	if highBit/8 != lowBit/8 {
		return 0, errors.New("getRegister does not support breaking values across byte boundaries")
	}

	i := byte(bytes) - highBit/8 - 1

	return bitUtils.SelectBits(raw[i], highBit%8, lowBit%8), nil
}
