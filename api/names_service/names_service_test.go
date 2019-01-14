package names_service

import "testing"

func TestCreateJavaName(t *testing.T) {
	testArtifacts := map[string]string {
		"Beginning Segment for Invoice": "beginningSegmentForInvoice",
		"Invoice Number": "invoiceNumber",
		"Baseline Item Data (Invoice)": "baselineItemDataInvoice",
		"ITEM _/<>+`*&^%$##!_()=^%#@ number - 15098765432321" : "itemNumber15098765432321",
		"Item Number": "itemNumber",
		"Item_ Number": "itemNumber1",
	}

	for sendValue, expectedValue := range testArtifacts {
		name := CreateJavaName(sendValue, "Segment")
		if expectedValue != name {
			t.Error(expectedValue, name)
		}
	}
}
