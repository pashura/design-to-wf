package xd_service_test

import (
	"os"
	"testing"
)

func TestXDService(t *testing.T) {
	if os.Getenv("DRONE") == "true" {
		t.Skip("Skipping integration test for CI")
	}
}
