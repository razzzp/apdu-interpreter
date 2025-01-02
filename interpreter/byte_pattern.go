package interpreter

import (
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type BytePatternMatcher struct {
	HighMask    *byte
	LowMask     *byte
	Pattern     string
	Description string
}

// Interpret implements Matcher.
func (bp *BytePatternMatcher) Interpret(i Interpretations, b byte) error {
	if bp.Matches(b) && bp.Description != "" {
		i.Add(bp.Description)
	}
	return nil
}

// Matches implements Matcher.
func (bp *BytePatternMatcher) Matches(b byte) bool {
	// check high nibble
	if bp.HighMask != nil {
		// 240 = 0b1111_0000
		if b&240 != *bp.HighMask {
			return false
		}
	}
	// check high nibble
	if bp.LowMask != nil {
		// 15 = 0b0000_1111
		if b&15 != *bp.LowMask {
			return false
		}
	}
	return true
}

func BytePattern(pattern string, description string) (ByteInterpreter, error) {
	pattern = strings.TrimSpace(pattern)
	if pattern == "" {
		return nil, nil
	}
	if len(pattern) != 2 {
		return nil, errors.New("pattern must be 2 chars long")
	}
	// check invalid hex chars, allow 'X'
	matched, _ := regexp.MatchString("[0-9-fA-FxX]+", pattern)
	if !matched {
		return nil, fmt.Errorf("patern contains invalid characters: %s", pattern)
	}

	bm := &BytePatternMatcher{
		Pattern:     pattern,
		Description: description,
	}

	// convert high nibble to mask
	if unicode.ToUpper(rune(pattern[0])) != 'X' {
		// if not x should be valid
		high, _ := hex.DecodeString(string(pattern[0]) + "0")
		mask := high[0]
		bm.HighMask = &mask
	}

	// convert high nibble to mask
	if unicode.ToUpper(rune(pattern[1])) != 'X' {
		// if not x should be valid
		low, _ := hex.DecodeString("0" + string(pattern[1]))
		mask := low[0]
		bm.LowMask = &mask
	}

	return bm, nil
}

func BytePatterns(patterns []string, description string) (result []ByteInterpreter, err error) {
	for _, pattern := range patterns {
		byteIntp, err := BytePattern(pattern, description)
		if err != nil {
			return nil, err
		}
		result = append(result, byteIntp)
	}
	return result, nil
}
