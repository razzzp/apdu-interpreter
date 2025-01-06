package interpreter

import "fmt"

type lengthInterpreter struct{}

func NewLengthInterpreter() *lengthInterpreter {
	return &lengthInterpreter{}
}

func (ii *lengthInterpreter) Matches(b byte) bool {
	// always match
	return true
}

func (ii *lengthInterpreter) Interpret(i Interpretations, b byte) error {
	// convert to int
	i.Add(fmt.Sprintf("Length: %d", uint(b)))
	return nil
}
