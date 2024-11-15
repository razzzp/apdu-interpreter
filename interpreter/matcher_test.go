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

func TestByteMatcher_Matches_DifferentCase_ReturnTrue(t *testing.T) {
	obj, _ := interpreter.NewByteMatcher("aB", "")

	assert.True(t, obj.Matches("Ab"))
}

func TestByteMatcher_Matches_NoMatch_ReturnFalse(t *testing.T) {
	obj, _ := interpreter.NewByteMatcher("aX", "")

	assert.False(t, obj.Matches("bb"))
}
