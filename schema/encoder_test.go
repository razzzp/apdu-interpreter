package schema_test

import (
	"strings"
	"testing"

	"github.com/razzzp/apdu-interpreter/schema"
)

func TestYamlEncoder_Encode(t *testing.T) {
	// arrange
	inputDef := schema.SchemaDefinition{
		Name:        "Test",
		Group:       "Test Group",
		Version:     "v0.1",
		Description: "Test Description",
		Labels: []string{
			"Label1",
			"Label2",
		},
		Spec: []schema.CommandResponseDefinition{
			schema.CommandResponseDefinition{
				Name:        "Cmd",
				Description: "Test",
				Command: schema.CommandDefinition{
					Cla: []schema.ByteDefinition{
						{
							BitPattern: &schema.BitPatternDefinition{
								Description: "Equals 1",
								Pattern:     "00000001",
							},
						},
					},
				},
			},
		},
	}
	strBuilder := strings.Builder{}
	encoder := schema.NewYamlSchemaEncoder(&strBuilder)
	// execute
	encoder.Encode(&inputDef)
	result := strBuilder.String()
	//
	t.Log(result)
}
