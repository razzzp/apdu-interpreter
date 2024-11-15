package interpreter_test

import (
	"apdu-interpreter/interpreter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteMatcher_Matches_ExactMatch_ReturnTrue(t *testing.T) {
	obj, _ := interpreter.NewByteMatcher("00", "")

	assert.True(t, obj.Matches("00"))
}

func TestByteMatcher_Matches_HighNibbleAny_ReturnTrue(t *testing.T) {
	obj, _ := interpreter.NewByteMatcher("X0", "")

	assert.True(t, obj.Matches("50"))
}

func TestByteMatcher_Matches_LowNibbleAny_ReturnTrue(t *testing.T) {
	obj, _ := interpreter.NewByteMatcher("0X", "")

	assert.True(t, obj.Matches("05"))
}
