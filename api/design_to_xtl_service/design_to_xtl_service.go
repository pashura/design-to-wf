package design_to_xtl_service

import (
	"fmt"
	"github.com/pashura/design-to-wf/api/design_structs"
	"github.com/pashura/design-to-wf/api/names_service"
	"github.com/pashura/design-to-wf/api/edi_info_service"
	"github.com/pashura/design-to-wf/api/xtl_structs"
	"strings"
	"time"
)

var JAVA_PACKAGE_NAME string
var SYS_DATE string

func ConvertDesignToXtl(design design_structs.Design) xtl_structs.Xtl {
	setupConstants(design)
	return createXtl(design)
}

func setupConstants(design design_structs.Design) {
	JAVA_PACKAGE_NAME = buildJavaPackageName(design.DesignMeta.HiddenSchema.OrgName)
	SYS_DATE = time.Now().Format("01/02/2006")
}
func buildJavaPackageName(orgName string) string {
	companyNameChunks := strings.Split(orgName, "_")
	for i, chunk := range companyNameChunks {
		companyNameChunks[i] = strings.Title(strings.ToLower(chunk))
	}
	return strings.Join(companyNameChunks, "")
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
		atts.Revision = design.DesignMeta.ViewedSchema.Version
	} else {
		atts.Direction = "I"
		atts.MaxSource = "-1"
		atts.Type = design.DesignMeta.HiddenSchema.Document
		atts.Revision = design.DesignMeta.HiddenSchema.Version
	}
	atts.DesignDate = SYS_DATE
	atts.JavaPackageName = JAVA_PACKAGE_NAME
	atts.XtencilType = "FEDS"
	atts.Designerversion = "2.8.4e"
	atts.Owner = design.DesignMeta.HiddenSchema.OrgName
	atts.Displayer = "TabPanel"
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
		elements = append(elements, group)
	}
	return elements
}

func createGroup(designObject design_structs.Object) xtl_structs.Element {
	if len(designObject.Children) > 0 {
		newGroup := xtl_structs.Element{}
		newGroup.Atts = createGroupAtts(designObject)
		newGroup.Name = "GROUPDEF"
		for i := range designObject.Children {
			child := designObject.Children[i]
			newGroup.Children = append(newGroup.Children, createGroup(child))
		}
		return newGroup
	} else {
		return createElement(designObject)
	}
}

func createElement(designObject design_structs.Object) xtl_structs.Element {
	newElement := xtl_structs.Element{}
	newElement.Atts = createElementAtts(designObject)
	newElement.Name = "FIELDDEF"
	return newElement
}

func createGroupAtts(designObject design_structs.Object) xtl_structs.Atts {
	atts := xtl_structs.Atts{}
	atts.Enable = "Y"
	atts.Min = designObject.MinOccurs
	atts.Name = names_service.CreateName(designObject.Name)
	atts.JavaName = names_service.CreateJavaName(designObject.Name)
	atts.Justification = "Left"
	if len(designObject.MaxOccurs) > 0 {
		atts.Max = designObject.MaxOccurs
	} else {
		atts.Max = designObject.MinOccurs
	}
	atts.Display = "Y"
	return atts
}

func createElementAtts(designObject design_structs.Object) xtl_structs.Atts {
	atts := xtl_structs.Atts{}
	//atts.Mandatory = designObject.Mandatory
	atts.Edi = "Y"
	atts.Name = names_service.CreateName(designObject.Attributes[0].EDIid)
	atts.JavaName = names_service.CreateJavaName(designObject.Attributes[0].EDIid)
	atts.Enable = "Y"
	atts.MinLength = designObject.MinLength
	atts.Editable = "Y"
	atts.MaxLength = designObject.MaxLength
	atts.Display = "Y"
	atts.SegmentTag, atts.Position, atts.SubPos = edi_info_service.EdiInfo(designObject.Name)
	return atts
}
