package main

import (
	"apdu-interpreter/interpreter"
	"apdu-interpreter/schema"
	"strings"
)

func main() {
	_ = schema.NewYamlSchemaDecoder(strings.NewReader("test"))
	_ = interpreter.InterpreterBuilder{}
}
