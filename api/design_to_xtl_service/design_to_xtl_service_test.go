package design_to_xtl_service

import (
	"github.com/pashura/design-to-wf/api/design_structs"
	"github.com/pashura/design-to-wf/api/xtl_structs"
	"reflect"
	"regexp"
	"testing"
)


func TestConvertDesignToXtlBaseChecks(t *testing.T) {

	hiddenMeta := design_structs.Schematype{OrgName: "TEST_ORG_NAME"}
	designMeta := design_structs.DesignMeta{HiddenSchema: hiddenMeta}
	testDesign := design_structs.Design{DesignMeta: designMeta}
	testDesign.Children = []design_structs.Object{}
	resultXtl := ConvertDesignToXtl(testDesign)

	if resultXtl.Infile != "" {
		t.Error(resultXtl.Infile)
	}
	if resultXtl.Repo != "TestOrgName.web" {
		t.Error(resultXtl.Repo)
	}
	if resultXtl.Branch != "new_form" {
		t.Error(resultXtl.Repo)
	}

	if resultXtl.Input.Name != "SPSFILE" {
		t.Error(resultXtl.Input.Name)
	}

	if match, _ := regexp.MatchString("\\d{2}//\\d{2}//\\d{4}", resultXtl.Input.Atts.Date); match {
		t.Error(resultXtl.Input.Atts.Date)
	}
	if resultXtl.Input.Atts.FileType != "FORM" {
		t.Error(resultXtl.Input.Atts.FileType)
	}
	if resultXtl.Input.Atts.Name != "SPS Commerce Xtencil" {
		t.Error(resultXtl.Input.Atts.Name)
	}
	if resultXtl.Input.Atts.Contents != "NORM" {
		t.Error(resultXtl.Input.Atts.Contents)
	}

}

func TestConvertDesignToXtlCreatesCorrectXtlStructure(t *testing.T) {

	expGroup := xtl_structs.Element{}
	expField := xtl_structs.Element{}

	expGroup.Atts.Display = "Y"
	expGroup.Atts.Enable = "Y"
	expGroup.Atts.Justification = "Left"
	expGroup.Name = "GROUPDEF"
	expGroup.Atts.Max = "2"
	expGroup.Atts.Min = "1"
	expField.Atts.Edi = "Y"
	expField.Atts.Enable = "Y"
	expField.Atts.Editable = "Y"
	expField.Atts.Display = "Y"
	expField.Atts.MinLength = "1"
	expField.Atts.MaxLength = "8"
	expField.Name = "FIELDDEF"
	expGroup.Children = []xtl_structs.Element{expField}
	testDesign := design_structs.Design{}
	testDesignElement := design_structs.Object{}
	testDesignElement.MinLength = "1"
	testDesignElement.MaxLength = "8"
	testDesignGroup := design_structs.Object{}
	testDesignGroup.MinOccurs = "1"
	testDesignGroup.MaxOccurs = "2"
	testDesignGroup.Children = []design_structs.Object{testDesignElement}
	testDesign.Children = []design_structs.Object{testDesignGroup}
	resultXtl := ConvertDesignToXtl(testDesign)

	if !reflect.DeepEqual(resultXtl.Input.Children[0].Children[0], expGroup) {
		t.Error(resultXtl.Input.Children[0].Children[0], expGroup)
	}
}
