package td_service_test

import (
	"os"
	"testing"
)

func TestGetDesign(t *testing.T) {
	if os.Getenv("DRONE") == "true" {
		t.Skip("Skipping integration test for CI")
	}
}
