package rep_service

import (
	"github.com/pashura/design-to-wf/api/xtl_structs"
)

func AddReps(xtl xtl_structs.Xtl) {
	docDef := xtl.Input.Children[0]
	for i := range docDef.Children {
		docDef.Children[i] = addRep(docDef.Children[i])
	}
}

func addRep(group xtl_structs.Element) xtl_structs.Element {
	if group.IsRepeatableGroup() {
		group.Children = []xtl_structs.Element{createRep(group)}
	}

	for i := range group.Children {
		group.Children[i] = addRep(group.Children[i])
	}

	return group
}

func createRep(group xtl_structs.Element) xtl_structs.Element {
	rep := xtl_structs.Element{}
	rep.Name = "GROUPDEF"
	rep.Atts.Min = "1"
	rep.Atts.Max = "1"
	rep.Atts.JavaName = group.Atts.JavaName + "Rep"
	rep.Atts.Name = group.Atts.Name + " Rep"
	rep.Children = group.Children
	return rep
}
