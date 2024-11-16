package interpreter_test

import (
	"apdu-interpreter/interpreter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPadStringLeft_5Chars_PadsTo8(t *testing.T) {
	s := interpreter.PadStringLeft("1", 'x', 8)

	assert.Len(t, s, 8)
	assert.Equal(t, s, "xxxxxxx1")
}

func TestPadStringLeft_8Chars_NotPad(t *testing.T) {
	input := "12345678"
	s := interpreter.PadStringLeft(input, 'x', 8)

	assert.Len(t, s, 8)
	assert.Equal(t, input, s)
}

func TestBitPattern_1BitLsb_SameValue(t *testing.T) {
	bitIntp, err := interpreter.BitPattern("1", "desc")
	assert.Nil(t, err)
	assert.Equal(t, byte(1), bitIntp.ExpectedValue)
	assert.Equal(t, byte(1), bitIntp.Mask)
}

func TestBitPattern_1BitMsb_MaskCorrect(t *testing.T) {
	bitIntp, err := interpreter.BitPattern("1xxxxxxx", "desc")
	assert.Nil(t, err)
	assert.Equal(t, byte(128), bitIntp.ExpectedValue)
	assert.Equal(t, byte(128), bitIntp.Mask)
}

func TestBitPattern_8Bit_MaskCorrect(t *testing.T) {
	bitIntp, err := interpreter.BitPattern("10000001", "desc")
	assert.Nil(t, err)
	assert.Equal(t, byte(129), bitIntp.ExpectedValue)
	assert.Equal(t, byte(255), bitIntp.Mask)
}
