package structure_levels_service

import (
	"github.com/pashura/design-to-wf/api/design_structs"
	"github.com/pashura/design-to-wf/api/xtl_structs"
	"strings"
)

var structureLevels = make(map[string]string)

func DesignRootStructureLevelsFromSources(design design_structs.Design) {
	for i := range design.Children {
		child := design.Children[i]
		structureLevelsFromElementsSources(child)
	}
}

func structureLevelsFromElementsSources(object design_structs.Object) {
	if len(object.Children) > 0 {
		for i := range object.Children {
			child := object.Children[i]
			structureLevelsFromElementsSources(child)
		}
	} else {
		structureLevelFromElementSource(object)
	}
}

func structureLevelFromElementSource(designObject design_structs.Object) {
	source := designObject.Sourcing.Location

	if len(source) > 0 {
		level := strings.Split(source, "/")[1]
		segmentTag := designObject.Name[:len(designObject.Name)-2]
		_, present := structureLevels[level]
		if !present {
			structureLevels[level] = segmentTag
		}
	}
}

func structureLevelByItsFirstSegmentTag(segmentTag string) (string, bool) {
	for levelName, firstSegmentTag := range structureLevels {
		if segmentTag == firstSegmentTag {
			return levelName, true
		}
	}
	return "", false
}

func StructureLevelByHL(group xtl_structs.Element) (string, bool) {
	for i := range group.Children {
		child := group.Children[i]
		if child.Atts.SegmentTag == "HL" && child.Atts.Position == "03" {
			choices := child.Atts.Choices
			descriptionStartIndex := strings.Index(choices, ":") + 1
			return choices[descriptionStartIndex:], true
		}
	}
	return "", false
}

func StructureLevel(group xtl_structs.Element) (string, bool) {
	structureLevel, ok := StructureLevelByHL(group)
	if ok {
		return structureLevel, true
	}
	structureLevel, ok = structureLevelByItsFirstSegmentTag(group.Children[0].Atts.SegmentTag)
	if ok {
		return structureLevel, true
	}
	return "", false
}
