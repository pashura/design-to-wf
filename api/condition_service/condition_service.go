package condition_service

import (
	"fmt"
	"github.com/pashura/design-to-wf/api/design_structs"
	"github.com/pashura/design-to-wf/api/xtl_structs"
	"strings"
)

var VALIDATIONS map[string][]design_structs.Validation

func ProcessValidationsDesign(design design_structs.Design) map[string][]design_structs.Validation {
	VALIDATIONS = make(map[string][]design_structs.Validation)
	findValidations(design.Children, "")
	return VALIDATIONS
}

func findValidations(children []design_structs.Object, ediPath string) {
	for i := range children {
		child := children[i]
		if len(child.Children) > 0 {
			ediElement := getEdiPathDesign(child.Name)
			if checkOnDuplicates(ediElement, ediPath) {
				ediPath += "/" + ediElement
			}
			if len(child.Validation) > 0 {
				VALIDATIONS[ediPath] = child.Validation
			}
			findValidations(child.Children, ediPath)
		}
	}
}

func getEdiPathDesign(name string) string {
	if len(strings.Split(name, "-")) > 1 {
		return strings.Split(name, "-")[1]
	} else {
		return name
	}
}

func checkOnDuplicates(name string, ediPath string) bool {
	splitEdiPath := strings.Split(ediPath, "/")
	if splitEdiPath[len(splitEdiPath)-1] == name {
		return false
	}
	return true
}

func PrintValidations(validations map[string][]design_structs.Validation) {
	for i := range validations {
		fmt.Printf("Edi Path: %v\n", i)
		fmt.Printf("Validations: %v\n", validations[i])
		fmt.Printf("Count of Validations: %v\n", len(validations[i]))
	}
}

func ProcessValidationsXtl(xtl xtl_structs.Xtl) {
	documentDef := xtl.Input.Children[0]
	searchElementsForValidations(documentDef.Children, "")
}

func searchElementsForValidations(children []xtl_structs.Element, ediPath string) {
	for i := range children {
		child := children[i]
		if child.Atts.SegmentTag != "" {

		}
	}
}
