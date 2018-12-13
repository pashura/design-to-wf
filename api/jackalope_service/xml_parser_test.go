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
	var transactionRef = "Transaction-850"

	var expectedElementName = "Transaction Set Identifier Code"
	var expectedSegmentName = "Transaction Set Header"
	var expectedTransactionName = "Purchase Order"

	if os.Getenv("DRONE") == "true" {
		t.Skip("Skipping integration test for CI")
	}

	elementName := jackalope_service.Documentation("xml_parser_test_resources.xml", elementID)

	segmentName := jackalope_service.Documentation("xml_parser_test_resources.xml", segmentRef)

	transactionName := jackalope_service.Documentation("xml_parser_test_resources.xml", transactionRef)

	if strings.TrimSpace(elementName) != expectedElementName {
		t.Errorf("got: %v, wanted: %v", elementName, expectedElementName)
	}

	if strings.TrimSpace(segmentName) != expectedSegmentName {
		t.Errorf("got: %v, wanted: %v", segmentName, expectedSegmentName)
	}

	if strings.TrimSpace(transactionName) != expectedTransactionName {
		t.Errorf("got: %v, wanted: %v", transactionName, expectedTransactionName)
	}
}
