package apdu_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/razzzp/apdu-interpreter/apdu"
	"github.com/stretchr/testify/assert"
)

func TestApduLogParser_FullCommandResponse(t *testing.T) {
	// arrange
	input := bufio.NewReader(strings.NewReader("00 01 02 03 04 FF FF FF FF 05\n FF FF 06 07"))
	sut := apdu.NewApduLogParser(input)

	// execute
	cmd, err := sut.GetNextCommandResponse()
	if err != nil {
		t.Error(err)
	}

	// assert
	assert.Equal(t, byte(0), cmd.Command.Cla)
	assert.Equal(t, byte(1), cmd.Command.Ins)
	assert.Equal(t, byte(2), cmd.Command.P1)
	assert.Equal(t, byte(3), cmd.Command.P2)
	assert.Equal(t, byte(4), cmd.Command.P3)
	assert.Equal(t, []byte{255, 255, 255, 255}, cmd.Command.Data)
	assert.Equal(t, byte(5), cmd.Command.Le)

	assert.Equal(t, []byte{255, 255}, cmd.Response.Data)
	assert.Equal(t, byte(06), cmd.Response.SW1)
	assert.Equal(t, byte(07), cmd.Response.SW2)
}

func TestApduLogParser_NoP3NoRData(t *testing.T) {
	// arrange
	input := bufio.NewReader(strings.NewReader("00 01 02 03\n 06 07"))
	sut := apdu.NewApduLogParser(input)

	// execute
	cmd, err := sut.GetNextCommandResponse()
	if err != nil {
		t.Error(err)
	}

	// assert
	assert.Equal(t, byte(0), cmd.Command.Cla)
	assert.Equal(t, byte(1), cmd.Command.Ins)
	assert.Equal(t, byte(2), cmd.Command.P1)
	assert.Equal(t, byte(3), cmd.Command.P2)
	assert.Equal(t, byte(0), cmd.Command.P3)
	assert.Nil(t, cmd.Command.Data)
	assert.Equal(t, byte(0), cmd.Command.Le)

	assert.Nil(t, cmd.Response.Data)
	assert.Equal(t, byte(06), cmd.Response.SW1)
	assert.Equal(t, byte(07), cmd.Response.SW2)
}
