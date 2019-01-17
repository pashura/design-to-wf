package structure_levels_service

import (
	"encoding/json"
	"fmt"
	"github.com/pashura/design-to-wf/api/design_structs"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func testDesignObject() design_structs.Design {
	jsonFile, err := os.Open("structure_levels_service_test_resources.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	design := design_structs.Design{}
	json.Unmarshal(byteValue, &design)
	return design
}

func TestStructureLevelsFromDesign(t *testing.T) {
	StructureLevelsFromDesign(testDesignObject())

	expectedValue := map[string]string{"Header": "BIG",
		"LineItem": "IT1",
		"Summary":  "TDS"}

	eq := reflect.DeepEqual(expectedValue, structureLevels)

	if !eq {
		t.Error(expectedValue, structureLevels)
	}
}

func TestStructureLevelByItsFirstSegmentTag(t *testing.T) {
	StructureLevelsFromDesign(testDesignObject())

	segmentTags := [3]string{"BIG", "IT1", "TDS"}

	for _, segmentTag := range segmentTags {
		_, ok := StructureLevelByItsFirstSegmentTag(segmentTag)
		if !ok {
			t.Error(segmentTag, structureLevels)
		}
	}
}
