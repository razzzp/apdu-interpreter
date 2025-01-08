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

func TestApp_TS102221_V18(t *testing.T) {
	// arrange
	app := app.NewApduInterpreterApp(app.ApduInterpreterConfiguration{
		InputFile:  "../test_inputs/etsi_102_221_v18.1.0_log.txt",
		SchemaPath: "../specifications/etsi_102_221_v18.1.0.yaml",
	})

	// execute
	app.Run()
}

func TestApp_GP_CardSpec_V2_3_1(t *testing.T) {
	// arrange
	app := app.NewApduInterpreterApp(app.ApduInterpreterConfiguration{
		InputFile:  "../test_inputs/gp_cardspec_v2.3.1_log.txt",
		SchemaPath: "../specifications/gp_cardspec_v2.3.1.yaml",
	})

	// execute
	app.Run()
}
