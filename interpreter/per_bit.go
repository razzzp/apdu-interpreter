package interpreter

import (
	"encoding/hex"
)

func NewBytePerBitIntp() *BytePerBitIntp {
	return &BytePerBitIntp{}
}

type BytePerBitIntp struct {
	BitDefs []BitIntp
}

func (bb *BytePerBitIntp) Matches(bytes string) bool {
	// matches if any bit matches
	for _, intp := range bb.BitDefs {
		if intp.Matches(bytes) {
			return true
		}
	}
	return false
}

func (bb *BytePerBitIntp) Interpret(bytes string) (string, error) {
	panic("not implemented") // TODO: Implement
}

type BitIntp struct {
	BitPos      int
	OneIsOff    bool
	Description string
}

func (bi *BitIntp) Matches(bytes string) bool {
	// parse to byte
	hex.DecodeString(bytes)
	return false
}

func (bi *BitIntp) Interpret(bytes string) (map[string]any, error) {
	panic("not implemented") // TODO: Implement
}
