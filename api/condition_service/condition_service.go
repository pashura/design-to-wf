package condition_service

import (
	"fmt"
	"github.com/pashura/design-to-wf/api/design_structs"
	"github.com/pashura/design-to-wf/api/xtl_structs"
	"strings"
)

func DesignValidations(design design_structs.Design) map[string][]design_structs.Validation {
	validations := make(map[string][]design_structs.Validation)
	findValidations(&validations, design.Children, "")
	return validations
}

func ExtraRecords(design design_structs.Design) []string {
	extraRecords := make([]string, 0)
	findExtraRecords(&extraRecords, design.Children, "")
	return extraRecords
}

func findExtraRecords(extraRecords *[]string, children []design_structs.Object, ediPath string) {
	for _, child := range children {
		if len(child.Children) > 0 {
			ediPath = buildEdiPath(child, ediPath)
			if child.DropExtraRecords && !ifElementInSlice(ediPath, extraRecords) {
				*extraRecords = append(*extraRecords, ediPath)
			}
			findExtraRecords(extraRecords, child.Children, ediPath)
		}
	}
}

func ifElementInSlice(element string, testSlice *[]string) bool {
	for _, testElement := range *testSlice {
		if testElement == element {
			return true
		}
	}
	return false
}

func findValidations(validations *map[string][]design_structs.Validation, children []design_structs.Object, ediPath string) {
	for _, child := range children {
		if len(child.Children) > 0 {
			ediPath = buildEdiPath(child, ediPath)
			if len(child.Validation) > 0 {
				(*validations)[ediPath] = child.Validation
			}
			findValidations(validations, child.Children, ediPath)
		}
	}
}

func buildEdiPath(child design_structs.Object, ediPath string) string {
	ediElement := getEdiPathDesign(child.Name)
	if checkOnDuplicates(ediElement, ediPath) {
		ediPath += "/" + ediElement
	}
	return ediPath
}

func getEdiPathDesign(name string) string {
	ediPathChunks := strings.Split(name, "-")
	if len(ediPathChunks) > 1 {
		return ediPathChunks[1]
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
