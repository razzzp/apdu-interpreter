package interpreter

import (
	"fmt"
)

type BitPatternInterpreter struct {
	ExpectedValue byte
	Mask          byte
	Pattern       string
	Description   string
}

func PadStringLeft(s string, c rune, length int) string {
	res := s

	for i := 0; i < length-len(s); i++ {
		res = string(c) + res
	}
	return res
}

// Helper builder to define just one bit
//
//	bitNum is a 1-based bit index
//	oneIsOn determines whether it will match a 1 or 0 at bitNum
func SingleBitDefinition(bitNum int, zeroIsOn bool, desc string) (*BitPatternInterpreter, error) {
	if bitNum < 1 || bitNum > 8 {
		return nil, fmt.Errorf("bit number must be between 0 and 7, got: %d", bitNum)
	}
	runes := []rune("xxxxxxxx")
	if zeroIsOn {
		runes[8-bitNum] = '0'
	} else {
		runes[8-bitNum] = '1'
	}
	return BitPattern(string(runes), desc)
}

// Pattern in the from of '0011xx'
// 'x's are don't cares
// if < 8 bits pads on left with 'x's
func BitPattern(pattern string, desc string) (*BitPatternInterpreter, error) {
	// pad left to 8 bytes with don't cares

	pattern = PadStringLeft(pattern, 'x', 8)
	expVal := 0
	mask := 0
	for pos, c := range pattern {
		if c == '1' {
			expVal += 1 << (7 - pos)
			mask += 1 << (7 - pos)
		} else if c == '0' {
			mask += 1 << (7 - pos)
		} else if c == 'x' {
			// do nothing
		} else {
			return nil, fmt.Errorf("failed to build bit pattern, invalid char %c", c)
		}
	}
	return &BitPatternInterpreter{
		ExpectedValue: byte(expVal),
		Mask:          byte(mask),
		Description:   desc,
		Pattern:       pattern,
	}, nil
}

func (bp *BitPatternInterpreter) Matches(b byte) bool {
	return b&bp.Mask == bp.ExpectedValue
}

func (bp *BitPatternInterpreter) Interpret(i Interpretation, b byte) error {
	if bp.Matches(b) {
		i.Add(fmt.Sprintf("0b%s: %s", bp.Pattern, bp.Description))
	}

	return nil
}
