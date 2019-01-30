package design_structs

import "strings"

func (object Object) SegmentName() string {
	if strings.Contains(object.Name, "-") {
		return object.Name[strings.Index(object.Name, "-")+1:]
	}
	return ""
}

func (object Object) IsLoop() bool {
	if len(object.Children) > 0 {
		if object.SegmentName() == object.Children[0].SegmentName() {
			return true
		}
	}
	return false
}

func (object Object) RestrictionAttributes() Object {
	for _, attr := range object.Attributes {
		if attr.ElementType == "restriction" {
			return attr
		}
	}
	return Object{}
}
