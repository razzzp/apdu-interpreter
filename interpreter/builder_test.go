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
		Description: "Test",
		Cla: []schema.ByteDefinition{
			{
				BitPattern: &schema.BitPatternDefinition{
					Description: "Equals 1",
					Pattern:     "00000001",
				},
			},
		},
	}
	expectedIntp := interpreter.BitPatternInterpreter{
		ExpectedValue: 1,
		Mask:          ByteFromHex("FF"),
		Pattern:       "00000001",
		Description:   "Equals 1",
	}

	intp, err := builder.BuildCommandInterpreter(&inputDef)

	assert.Nil(t, err)
	assert.NotNil(t, intp.ClaMatcher)
	assert.Len(t, intp.ClaMatcher, 1)
	assert.Equal(t, &expectedIntp, intp.ClaMatcher[0])
}

func TestInterpreterBuilder_BytePatterns(t *testing.T) {
	builder := interpreter.InterpreterBuilder{}
	inputDef := schema.CommandDefinition{
		Name:        "Cmd",
		Description: "Test",
		Cla: []schema.ByteDefinition{
			{
				BytePatterns: &schema.BytePatternsDefinition{
					Patterns: []string{
						"AX",
						"8X",
						"X1",
					},
				},
			},
		},
	}
	hm := byte(160)
	expectedFirstIntp := interpreter.BytePatternMatcher{
		HighMask: &hm,
		Pattern:  "AX",
	}

	lm := byte(1)
	expectedLastIntp := interpreter.BytePatternMatcher{
		LowMask: &lm,
		Pattern: "X1",
	}

	intp, err := builder.BuildCommandInterpreter(&inputDef)
	if err != nil {
		t.Error(err)
	}
	assert.Len(t, intp.ClaMatcher, 3)
	assert.Equal(t, &expectedFirstIntp, intp.ClaMatcher[0])
	assert.Equal(t, &expectedLastIntp, intp.ClaMatcher[2])
}
