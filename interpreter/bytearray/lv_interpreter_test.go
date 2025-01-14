package bytearray_test

import (
	"testing"

	"github.com/razzzp/apdu-interpreter/interpreter"
	"github.com/razzzp/apdu-interpreter/interpreter/bytearray"
	"github.com/stretchr/testify/assert"
)

func TestLVInterpreter_Start0LastByteAtEnd(t *testing.T) {
	// arrange
	sut := bytearray.LvInterpreter{
		Label: "label 1",
	}
	i := interpreter.DataInterpretations{}
	input := []byte{05, 00, 01, 02, 03, 04}

	// execute
	nextIdx, err := sut.Interpret(&i, input, 0)

	// assert
	assert.Equal(t, 6, nextIdx)
	assert.Nil(t, err)
	assert.Equal(t, "label 1: 00 01 02 03 04", i.Intps[0])
}

func TestLVInterpreter_Start0LastByteInMiddle(t *testing.T) {
	// arrange
	sut := bytearray.LvInterpreter{
		Label: "label 1",
	}
	i := interpreter.DataInterpretations{}
	input := []byte{04, 00, 01, 02, 03, 04}

	// execute
	nextIdx, err := sut.Interpret(&i, input, 0)

	// assert
	assert.Equal(t, 5, nextIdx)
	assert.Nil(t, err)
	assert.Equal(t, "label 1: 00 01 02 03", i.Intps[0])
}

func TestLVInterpreter_Start0NotEnoughBytes(t *testing.T) {
	// arrange
	sut := bytearray.LvInterpreter{
		Label: "label 1",
	}
	i := interpreter.DataInterpretations{}
	input := []byte{06, 00, 01, 02, 03, 04}

	// execute
	nextIdx, err := sut.Interpret(&i, input, 0)

	// assert
	assert.Equal(t, 6, nextIdx)
	assert.Nil(t, err)
	assert.Equal(t, "label 1: 00 01 02 03 04", i.Intps[0])
}
