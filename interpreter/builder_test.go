package interpreter_test

import (
	"testing"

	"github.com/razzzp/apdu-interpreter/interpreter"
	"github.com/razzzp/apdu-interpreter/schema"
	"github.com/stretchr/testify/assert"
)

func TestInterpreterBuilder_BuildCommandIntp_Cla_NotNil(t *testing.T) {
	builder := interpreter.InterpreterBuilder{}
	inputDef := schema.CommandDefinition{
		Name:        "Cmd",
		Decsription: "Test",
		Cla: []schema.ByteDefinition{
			{
				BitPattern: &schema.BitPatternDefinition{
					Description: "Equals 1",
					Pattern:     "00000001",
				},
			},
		},
	}
	expectedIntp := interpreter.BitPatternIntp{
		ExpectedValue: 1,
		Mask:          ByteFromHex("FF"),
		Pattern:       "00000001",
		Description:   "Equals 1",
	}

	intp, err := builder.BuildCommandIntp(&inputDef)

	assert.Nil(t, err)
	assert.NotNil(t, intp.ClaMatcher)
	assert.Len(t, intp.ClaMatcher, 1)
	assert.Equal(t, &expectedIntp, intp.ClaMatcher[0])
}
