package app

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/razzzp/apdu-interpreter/apdu"
	"github.com/razzzp/apdu-interpreter/interpreter"
	"github.com/razzzp/apdu-interpreter/schema"
)

type ApduInterpreterConfiguration struct {
	InputFile  string
	SchemaPath string
}

type ApduInterpreterApp struct {
	Config ApduInterpreterConfiguration
}

func NewApduInterpreter(config ApduInterpreterConfiguration) *ApduInterpreterApp {
	return &ApduInterpreterApp{}
}

func (a *ApduInterpreterApp) BuildSchema() (*schema.SchemaDefinition, error) {
	stat, err := os.Stat(a.Config.SchemaPath)
	if err != nil {
		return nil, fmt.Errorf("error building schema: %v", err)
	}

	if stat.IsDir() {
		//TODO
		return nil, fmt.Errorf("not immplemented")
	} else {
		file, err := os.Open(a.Config.SchemaPath)
		if err != nil {
			return nil, fmt.Errorf("error opening schema file: %v", err)
		}

		schemaBuilder := schema.NewYamlSchemaDecoder(file)
		schema, err := schemaBuilder.Decode()
		if err != nil {
			return nil, fmt.Errorf("error building schema: %v", err)
		}
		return schema, nil
	}
}

func (a *ApduInterpreterApp) BuildApduParser() (apdu.ApduParser, error) {
	stat, err := os.Stat(a.Config.InputFile)
	if err != nil {
		return nil, fmt.Errorf("error building schema: %v", err)
	}

	if stat.IsDir() {
		//TODO
		return nil, fmt.Errorf("input must be a file")
	}

	file, err := os.Open(a.Config.SchemaPath)
	if err != nil {
		return nil, fmt.Errorf("error opening schema file: %v", err)
	}
	parser := apdu.NewApduLogParser(bufio.NewReader(file))
	return &parser, nil
}

func (a *ApduInterpreterApp) Run() {
	// parse schema
	schema, err := a.BuildSchema()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// build interpreter
	builder := interpreter.InterpreterEngineBuilder{}
	engine := builder.BuildSchema(schema)

	// prepare apdu parser
	parser, err := a.BuildApduParser()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	engine.Parser = parser
}