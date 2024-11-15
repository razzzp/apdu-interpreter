package interpreter

import (
	"apdu-interpreter/apdu"
	"fmt"
	"regexp"
	"unicode"
)

type Interpreter interface {
	Matches(apdu apdu.ApduCommand) bool
	Interpret(apdu apdu.ApduCommand) (Interpretation, error)
}

type InterpreterEngine struct {
	Schema              Schema
	CommandInterpreters []ApduCommandInterpreter
}

type Matcher interface {
	Matches(bytes string) bool
	Interpret(bytes string) (string, error)
}

type ApduCommandInterpreter struct {
	ClaMatcher      Matcher
	InsMatcher      Matcher
	P1Matcher       Matcher
	P2Matcher       Matcher
	P3Matcher       Matcher
	DataInterpreter any
	LeMatcher       Matcher
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
	matched, _ := regexp.MatchString(pattern, "[0123456789abcdefABCDEFxX]+")
	if !matched {
		return nil, fmt.Errorf("patern contains invalid characters: %s", pattern)
	}

	return &ByteMatcher{
		pattern: pattern,
	}, nil
}

type Interpretation struct {
}
