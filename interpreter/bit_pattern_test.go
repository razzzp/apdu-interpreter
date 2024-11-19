package interpreter_test

import (
	"testing"

	"github.com/razzzp/apdu-interpreter/interpreter"
	"github.com/razzzp/apdu-interpreter/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestInterpret_1BitMatch_AddInterpretation(t *testing.T) {
	bitIntp, err := interpreter.BitPattern("1", "first bit on")
	byteI := mocks.NewMockInterpretation(t)
	assert.Nil(t, err)

	byteI.On("Add", mock.Anything)

	bitIntp.Interpret(byteI, ByteFromHex("F1"))

	byteI.AssertCalled(t, "Add", "0bxxxxxxx1: first bit on")
}

func TestInterpret_1BitNoMatch_NoInterpretation(t *testing.T) {
	bitIntp, err := interpreter.BitPattern("0", "first bit on")
	byteI := mocks.NewMockInterpretation(t)
	assert.Nil(t, err)

	bitIntp.Interpret(byteI, ByteFromHex("F1"))

	byteI.AssertNotCalled(t, "Add", "0bxxxxxxx0: first bit on")
}

func TestInterpret_8BitMatch_AddInterpretation(t *testing.T) {
	bitIntp, err := interpreter.BitPattern("11001101", "exactly")
	byteI := mocks.NewMockInterpretation(t)
	assert.Nil(t, err)

	byteI.On("Add", mock.Anything)

	bitIntp.Interpret(byteI, ByteFromHex("CD"))

	byteI.AssertCalled(t, "Add", "0b11001101: exactly")
}

func TestBitDef_Bit0_BuildsBitPatternIntp(t *testing.T) {
	bitIntp, err := interpreter.BitDef(1, false, "bit 1")

	assert.Nil(t, err)
	assert.Equal(t, byte(1), bitIntp.Mask)
	assert.Equal(t, byte(1), bitIntp.ExpectedValue)
	assert.Equal(t, "bit 1", bitIntp.Description)
}

func TestBitDef_Bit8ZeroOn_BuildsBitPatternIntp(t *testing.T) {
	bitIntp, err := interpreter.BitDef(8, true, "bit 8")

	assert.Nil(t, err)
	assert.Equal(t, byte(128), bitIntp.Mask)
	assert.Equal(t, byte(0), bitIntp.ExpectedValue)
	assert.Equal(t, "bit 8", bitIntp.Description)
}
