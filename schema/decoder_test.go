package schema_test

import (
	"strings"
	"testing"

	"github.com/razzzp/apdu-interpreter/schema"
	"github.com/stretchr/testify/assert"
)

func TestDecode_AllFieldsPresent_ReturnSchema(t *testing.T) {
	// arrange
	input := `name: Test
group: Test Group
version: v0.1
description: Test Description
labels:
  - Label1
  - Label2
commandDefinitions:
  - name: Cmd
    description: Test
    cla:
      - bitPattern:
          description: Equals 1
          pattern: "00000001"
`
	reader := strings.NewReader(input)
	decoder := schema.NewYamlSchemaDecoder(reader)

	//act
	schema, err := decoder.Decode()
	if err != nil {
		t.Error(err)
	}

	//assert
	assert.Equal(t, "Test", schema.Name)
	assert.Equal(t, "Test Description", schema.Description)
	assert.Equal(t, "Test Group", schema.Group)
	assert.Len(t, schema.Labels, 2)
	assert.Len(t, schema.CommandDefinitions, 1)
	assert.Equal(t, "Cmd", schema.CommandDefinitions[0].Name)
	assert.Equal(t, "Test", schema.CommandDefinitions[0].Description)
	assert.Len(t, schema.CommandDefinitions[0].Cla, 1)
	assert.Equal(t, "Equals 1", schema.CommandDefinitions[0].Cla[0].BitPattern.Description)
	assert.Equal(t, "00000001", schema.CommandDefinitions[0].Cla[0].BitPattern.Pattern)

	t.Logf("result: %+v\n", schema)
}
