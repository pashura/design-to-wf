package schema_enum_service_test

import (
	"os"
	"testing"
)

func TestGetSchemaEnums(t *testing.T) {
	if os.Getenv("DRONE") == "true" {
		t.Skip("Skipping integration test for CI")
	}
}