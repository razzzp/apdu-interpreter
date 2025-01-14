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
spec:
  - name: Cmd
    description: Test 
    command:  
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
		t.Fatal(err)
	}

	//assert
	assert.Equal(t, "Test", schema.Name)
	assert.Equal(t, "Test Description", schema.Description)
	assert.Equal(t, "Test Group", schema.Group)
	assert.Len(t, schema.Labels, 2)
	assert.Len(t, schema.Spec, 1)
	assert.Equal(t, "Cmd", schema.Spec[0].Name)
	assert.Equal(t, "Test", schema.Spec[0].Description)
	assert.Len(t, schema.Spec[0].Command.Cla, 1)
	assert.Equal(t, "Equals 1", schema.Spec[0].Command.Cla[0].BitPattern.Description)
	assert.Equal(t, "00000001", schema.Spec[0].Command.Cla[0].BitPattern.Pattern)

	t.Logf("result: %+v\n", schema)
}

func TestDecode_DataFields_ReturnSchema(t *testing.T) {
	// arrange
	input := `name: Test
group: Test Group
version: v0.1
description: Test Description
labels:
  - Label1
  - Label2
spec:
  - name: Cmd
    description: Test 
    command:  
      data:
        - when:
            label: First when
            criteria:
              p1:
                - bytePattern:
                    pattern: "0X"
            interpreters:
              - lengthValue:
                  label: Length Value Label
`
	reader := strings.NewReader(input)
	decoder := schema.NewYamlSchemaDecoder(reader)

	//act
	schema, err := decoder.Decode()
	if err != nil {
		t.Fatal(err)
	}

	//assert
	assert.Len(t, schema.Spec[0].Command.Data, 1)
	assert.NotNil(t, schema.Spec[0].Command.Data[0].When)
	assert.Equal(t, "First when", schema.Spec[0].Command.Data[0].When.Label)
	assert.Len(t, schema.Spec[0].Command.Data[0].When.Criteria.P1, 1)
	assert.Equal(t, "0X", schema.Spec[0].Command.Data[0].When.Criteria.P1[0].BytePattern.Pattern)
	assert.Len(t, schema.Spec[0].Command.Data[0].When.Interpreters, 1)
	assert.Equal(t, "Length Value Label", schema.Spec[0].Command.Data[0].When.Interpreters[0].LengthValue.Label)

	t.Logf("result: %+v\n", schema)
}
