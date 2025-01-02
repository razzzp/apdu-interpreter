package app_test

import (
	"testing"

	"github.com/razzzp/apdu-interpreter/app"
)

func TestApp_Case1(t *testing.T) {
	// arrange
	app := app.NewApduInterpreterApp(app.ApduInterpreterConfiguration{
		InputFile:  "../test_inputs/case_1_apdu_log.txt",
		SchemaPath: "../test_inputs/case_1_schema.yaml",
	})

	// execute
	app.Run()
}
