package jackalope_service_test

import (
	"github.com/pashura/design-to-wf/api/jackalope_service"
	"os"
	"strings"
	"testing"
)

func TestDocumentation(t *testing.T) {
	var elementID = "143"
	var segmentRef = "ST"

	var expectedElementName = "Transaction Set Identifier Code "
	var expectedSegmentName = "Transaction Set Header"

	if os.Getenv("DRONE") == "true" {
		t.Skip("Skipping integration test for CI")
	}

	elementName := jackalope_service.Documentation("xml_parser_test_resources.xml", elementID)

	segmentName := jackalope_service.Documentation("xml_parser_test_resources.xml", segmentRef)

	if strings.TrimSpace(elementName) != expectedElementName {
		t.Errorf("got: %v, wanted: %v", elementName, expectedElementName)
	}

	if strings.TrimSpace(segmentName) != "Transaction Set Header" {
		t.Errorf("got: %v, wanted: %v", segmentName, expectedSegmentName)
	}
}
