package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBits(t *testing.T) {
	test := byte(0xE0)
	SetBits(&test, 3, 1, 0x7)
	assert.Equal(t, byte(0xEE), test)
}
