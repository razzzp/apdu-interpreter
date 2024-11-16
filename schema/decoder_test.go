package schema_test

import (
	"testing"
)

func TestDecode_AllFieldsPresent_ReturnSchema(t *testing.T) {
	// arrange
	// 	input := `name: Test Schema
	// group: some group
	// description: Some desc
	// labels:
	//   - Label1
	//   - Label 2
	// commandDefinitions:
	//   - cla: A0
	//     ins: A4
	//     p1: 00
	//     p2: 01
	//     p3: 03
	//     data: 112233
	//     le: 00
	// `
	// 	reader := strings.NewReader(input)
	// 	decoder := schema.NewYamlSchemaDecoder(reader)

	// 	//act
	// 	schema, err := decoder.Decode()

	// 	//assert
	// 	assert.Nil(t, err)
	// 	assert.NotNil(t, schema)
	// 	assert.Equal(t, schema.Name, "Test Schema")
	// 	assert.Equal(t, schema.Description, "Some desc")
	// 	assert.Equal(t, schema.Group, "some group")
	// 	assert.Len(t, schema.Labels, 2)
	// 	assert.Len(t, schema.CommandDefinitions, 1)

	// t.Logf("result: %+v\n", schema)
}
