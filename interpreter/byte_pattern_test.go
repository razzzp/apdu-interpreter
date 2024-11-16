package interpreter_test

import (
	"apdu-interpreter/interpreter"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ByteFromHex(hexstr string) byte {
	res, err := hex.DecodeString(hexstr)
	if err != nil {
		panic(err)
	}

	return res[0]
}

func TestByteMatcher_Matches_ExactMatch_ReturnTrue(t *testing.T) {
	obj, _ := interpreter.BytePattern("00", "")

	assert.True(t, obj.Matches(ByteFromHex("00")))
}

func TestByteMatcher_Matches_HighNibbleAny_ReturnTrue(t *testing.T) {
	obj, _ := interpreter.BytePattern("X0", "")

	assert.True(t, obj.Matches(ByteFromHex("50")))
}

func TestByteMatcher_Matches_LowNibbleAny_ReturnTrue(t *testing.T) {
	obj, _ := interpreter.BytePattern("0X", "")

	assert.True(t, obj.Matches(ByteFromHex("05")))
}

func TestByteMatcher_Matches_DifferentCase_ReturnTrue(t *testing.T) {
	obj, _ := interpreter.BytePattern("aB", "")

	assert.True(t, obj.Matches(ByteFromHex("Ab")))
}

func TestByteMatcher_Matches_NoMatch_ReturnFalse(t *testing.T) {
	obj, _ := interpreter.BytePattern("aX", "")

	assert.False(t, obj.Matches(ByteFromHex("bb")))
}
