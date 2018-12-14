package edi_info_service

import "testing"

func TestEdiInfo(t *testing.T) {
	testArtifacts := map[string][]string{
		"BSN01":    {"BSN", "01", ""},
		"RFF01-02": {"RFF", "01", "02"},
		"":         {"", "", ""},
	}

	for sendValue, expectedValues := range testArtifacts {
		segmentTag, position, subPos := EdiInfo(sendValue)
		if expectedValues[0] != segmentTag {
			t.Error(expectedValues, segmentTag)
		}
		if expectedValues[1] != position {
			t.Error(expectedValues, position)
		}
		if expectedValues[2] != subPos {
			t.Error(expectedValues, subPos)
		}
	}
}
