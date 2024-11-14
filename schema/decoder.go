package schema

import (
	"io"

	"gopkg.in/yaml.v3"
)

type SchemaDecoder interface {
	Decode() (*Schema, error)
}

func NewYamlSchemaDecoder(reader io.Reader) SchemaDecoder {
	result := yamlReader{
		reader: reader,
	}
	return &result
}

type yamlReader struct {
	reader io.Reader
}

func (yr *yamlReader) Decode() (*Schema, error) {
	decoder := yaml.NewDecoder(yr.reader)

	result := Schema{}
	err := decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
