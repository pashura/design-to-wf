package structure_levels_service

import (
	"github.com/pashura/design-to-wf/api/design_structs"
	"strings"
)

var structureLevels = make(map[string]string)

func GetStructureLevelsFromDesign(design design_structs.Design) {
	for i := range design.Children {
		child := design.Children[i]
		getStructureLevelsFromElementsSources(child)
	}
}

func getStructureLevelsFromElementsSources(object design_structs.Object) {
	if len(object.Children) > 0 {
		for i := range object.Children {
			child := object.Children[i]
			getStructureLevelsFromElementsSources(child)
		}
	} else {
		getStructureLevelFromElementSource(object)
	}
}

func getStructureLevelFromElementSource(designObject design_structs.Object) {
	source := designObject.Sourcing.Location

	if len(source) > 0 {
		level := strings.Split(source,"/")[1]
		segmentTag := designObject.Name[:len(designObject.Name)-2]
		_, present := structureLevels[level]
		if !present {
			structureLevels[level] = segmentTag
		}
	}
}

func GetStructureLevelByItsFirstSegmentTag(segmentTag string) (string, bool) {
	for levelName, firstSegmentTag := range structureLevels {
		if segmentTag == firstSegmentTag {
			return levelName, true
		}
	}
	return "", false
}




