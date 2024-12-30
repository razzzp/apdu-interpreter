package schema

import (
	"io"

	"gopkg.in/yaml.v3"
)

type SchemaEncoder interface {
	Encode(*SchemaDefinition) error
}

func NewYamlSchemaEncoder(writer io.Writer) SchemaEncoder {
	result := yamlEncoder{
		writer: writer,
	}
	return &result
}

type yamlEncoder struct {
	writer io.Writer
}

func (ye *yamlEncoder) Encode(schema *SchemaDefinition) error {
	encoder := yaml.NewEncoder(ye.writer)

	err := encoder.Encode(schema)
	return err
}
