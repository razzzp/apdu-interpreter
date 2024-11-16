package main

import (
	"strings"

	"github.com/razzzp/apdu-interpreter/schema"

	"github.com/razzzp/apdu-interpreter/interpreter"
)

func main() {
	_ = schema.NewYamlSchemaDecoder(strings.NewReader("test"))
	_ = interpreter.InterpreterBuilder{}
}
