package schema_test

import (
	"apdu-interpreter/schema"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode_AllFieldsPresent_ReturnSchema(t *testing.T) {
	input := `name: Test Schema
description: Some desc
labels:
  - Label1
  - Label 2
commandDefinitions:
  - cla: A0
    ins: A4
    p1: 00
    p2: 01
    p3: 03
    data: 112233
    le: 00
`
	reader := strings.NewReader(input)
	decoder := schema.NewYamlSchemaDecoder(reader)

	schema, err := decoder.Decode()
	assert.Nil(t, err)
	assert.NotNil(t, schema)

	t.Logf("result: %+v\n", schema)
}
