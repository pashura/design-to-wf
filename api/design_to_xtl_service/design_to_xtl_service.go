package design_to_xtl_service

import (
	"fmt"
	"github.com/pashura/design-to-wf/api/design_structs"
	"github.com/pashura/design-to-wf/api/design_to_xtl_service/edi_info_service"
	"github.com/pashura/design-to-wf/api/design_to_xtl_service/structure_levels_service"
	"github.com/pashura/design-to-wf/api/jackalope_service"
	"github.com/pashura/design-to-wf/api/names_service"
	"github.com/pashura/design-to-wf/api/properties"
	"github.com/pashura/design-to-wf/api/schema_enum_service"
	"github.com/pashura/design-to-wf/api/xtl_structs"
	"strings"
	"time"
)

var JAVA_PACKAGE_NAME string
var SYS_DATE string
var DATA_TYPES map[string]string

var currentGroup design_structs.Object

func ConvertDesignToXtl(design design_structs.Design, javaPackageName string) xtl_structs.Xtl {
	JAVA_PACKAGE_NAME = javaPackageName
	setupConstants()
	return createXtl(design)
}

func setupConstants() {
	SYS_DATE = time.Now().Format("01/02/2006")
	DATA_TYPES = map[string]string{
		"String":    "JString",
		"StringSet": "JMappedSet",
		"Time":      "JTime",
		"Date":      "JDate",
		"Integer":   "JInteger",
		"Decimal":   "JDouble",
	}
}

func createXtl(design design_structs.Design) xtl_structs.Xtl {
	baseXtl := xtl_structs.Xtl{}
	baseXtl.Repo = fmt.Sprintf("%v.web", JAVA_PACKAGE_NAME)
	baseXtl.Branch = "new_form"
	baseXtl.Input = createInput(design)
	baseXtl.Infile = ""
	return baseXtl
}

func createInput(design design_structs.Design) xtl_structs.XtlSide {
	baseRoot := xtl_structs.XtlSide{}
	baseRoot.Atts = createRootAtts()
	baseRoot.Name = "SPSFILE"
	baseRoot.Children = createDocumentDef(design)
	return baseRoot
}

func createDocumentDef(design design_structs.Design) [1]xtl_structs.DocumentDef {
	docDef := xtl_structs.DocumentDef{}
	docDef.Atts = createDocumentDefAtts(design)
	docDef.Name = "DOCUMENTDEF"
	docDef.Children = createStructure(design)
	return [1]xtl_structs.DocumentDef{docDef}
}

func createDocumentDefAtts(design design_structs.Design) xtl_structs.Atts {
	atts := xtl_structs.Atts{}
	atts.Origin = "design-to-webforms"
	if design.DesignMeta.ViewedSchema.Source {
		atts.Direction = "O"
		atts.MaxSource = "1"
		atts.Type = design.DesignMeta.ViewedSchema.Document
	} else {
		atts.Direction = "I"
		atts.MaxSource = "-1"
		atts.Type = design.DesignMeta.HiddenSchema.Document
	}
	atts.Revision = properties.Version
	atts.DesignDate = SYS_DATE
	atts.JavaPackageName = JAVA_PACKAGE_NAME
	atts.XtencilType = "FEDS"
	atts.Designerversion = "2.8.4e"
	atts.Owner = design.DesignMeta.HiddenSchema.OrgName
	atts.Displayer = "TabPanel"
	atts.Name = jackalope_service.Documentation(properties.Document)
	atts.JavaName = names_service.CreateJavaName(atts.Name, "")
	return atts
}

func createRootAtts() xtl_structs.Atts {
	atts := xtl_structs.Atts{}
	atts.Date = SYS_DATE
	atts.FileType = "FORM"
	atts.Name = "SPS Commerce Xtencil"
	atts.Contents = "NORM"
	return atts
}

func createStructure(design design_structs.Design) []xtl_structs.Element {
	var elements []xtl_structs.Element
	for i := range design.Children {
		child := design.Children[i]
		group := createGroup(child)
		elements = createStructureLevels(elements, group)
	}
	return elements
}

func createStructureLevels(elements []xtl_structs.Element, group xtl_structs.Element) []xtl_structs.Element {
	structureLevel, ok := structure_levels_service.GetStructureLevelByItsFirstSegmentTag(group.Children[0].Atts.SegmentTag)
	if ok {
		group.Atts.Name = structureLevel
		group.Atts.JavaName = names_service.CreateJavaName(structureLevel, "")
		elements = append(elements, group)
	} else {
		lastElementIndex := len(elements)-1
		elements[lastElementIndex].Children = append(elements[lastElementIndex].Children, group)
	}
	return elements
}

func createGroup(designObject design_structs.Object) xtl_structs.Element {
	if len(designObject.Children) > 0 {
		newGroup := xtl_structs.Element{}
		newGroup.Atts = createGroupAtts(designObject)
		newGroup.Name = "GROUPDEF"
		for i := range designObject.Children {
			currentGroup = designObject
			child := designObject.Children[i]
			if designObject.GetSegmentName() != child.GetSegmentName() {
				newGroup.Children = append(newGroup.Children, createGroup(child))
			} else {
				for k := range child.Children {
					miniChild := child.Children[k]
					newGroup.Children = append(newGroup.Children, createGroup(miniChild))
				}
			}
		}
		return newGroup
	} else {
		return createElement(designObject)
	}
}

func createGroupAtts(designObject design_structs.Object) xtl_structs.Atts {
	atts := xtl_structs.Atts{}
	atts.Enable = "Y"
	atts.Min = designObject.MinOccurs
	atts.Name = names_service.CreateName(designObject.Name)
	atts.JavaName = names_service.CreateJavaName(atts.Name, designObject.Name)
	atts.Justification = "Left"
	if len(designObject.MaxOccurs) > 0 {
		atts.Max = designObject.MaxOccurs
	} else {
		atts.Min = "1"
		atts.Max = "1"
	}
	atts.Display = "Y"
	return atts
}

func createElement(designObject design_structs.Object) xtl_structs.Element {
	newElement := xtl_structs.Element{}
	newElement.Atts = createElementAtts(designObject)
	newElement.Name = "FIELDDEF"
	return newElement
}

func createElementAtts(designObject design_structs.Object) xtl_structs.Atts {
	atts := xtl_structs.Atts{}
	if designObject.MinOccurs == "1" {
		atts.Mandatory = "Y"
	} else {
		atts.Mandatory = "N"
	}
	atts.Edi = "Y"
	atts.Enable = "Y"
	atts.Editable = "Y"
	atts.Display = "Y"
	atts.DefaultValue = designObject.DefaultValue
	if len(designObject.Attributes) > 0 {
		designObjectAttributes := designObject.Attributes[len(designObject.Attributes)-1]
		atts.Name = names_service.CreateName(designObjectAttributes.EDIid)
		atts.JavaName = names_service.CreateJavaName(atts.Name, currentGroup.Name)
		atts.ReferenceNum = designObjectAttributes.EDIid
		atts.DataType = DATA_TYPES[designObjectAttributes.DisplayName]
		if designObjectAttributes.HasEnum {
			atts.Choices = qualifiers(designObject.Name, designObject.Qualifiers)
		}
		atts.MinLength = designObjectAttributes.MinLength
		atts.MaxLength = designObjectAttributes.MaxLength
	}
	atts.SegmentTag, atts.Position, atts.SubPos = edi_info_service.EdiInfo(designObject.Name)
	return atts
}

func qualifiers(elementName, qualifiers string) string {
	groupName := fmt.Sprintf("Segment-%v", elementName[:len(elementName)-2])

	result := make([]string, 0)
	qualifierList := strings.Split(qualifiers, ",")
	for i := 0; i < len(qualifierList); i++ {
		qual := strings.TrimSpace(string(qualifierList[i]))
		description := schema_enum_service.GetSchemaEnums(groupName, elementName, qual)
		result = append(result, fmt.Sprintf("%v: %v", qual, description))
	}
	fmt.Println()

	return strings.Join(result, ", ")
}
