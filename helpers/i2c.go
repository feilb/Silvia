package helpers

import (
	"fmt"

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
