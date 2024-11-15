package interpreter

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type Matcher interface {
	Matches(bytes string) bool
	Interpret(bytes string) (string, error)
}

type ByteMatcher struct {
	pattern     string
	description string
}

// Interpret implements Matcher.
func (b *ByteMatcher) Interpret(byteStr string) (string, error) {
	panic("unimplemented")
}

// Matches implements Matcher.
func (b *ByteMatcher) Matches(byteStr string) bool {
	for pos, c := range byteStr {
		cur := unicode.ToUpper(rune(b.pattern[pos]))
		if unicode.ToUpper(c) != cur && cur != 'X' {
			return false
		}
	}
	return true
}

func NewByteMatcher(pattern string, description string) (Matcher, error) {
	pattern = strings.TrimSpace(pattern)
	if pattern == "" {
		return nil, nil
	}
	matched, _ := regexp.MatchString("[0-9-fA-FxX]+", pattern)
	if !matched {
		return nil, fmt.Errorf("patern contains invalid characters: %s", pattern)
	}

	return &ByteMatcher{
		pattern: pattern,
	}, nil
}
