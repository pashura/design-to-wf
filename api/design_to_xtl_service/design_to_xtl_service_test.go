package design_to_xtl_service

import (
	"github.com/pashura/design-to-wf/api/design_structs"
	"github.com/pashura/design-to-wf/api/design_to_xtl_service/structure_levels_service"
	"github.com/pashura/design-to-wf/api/jackalope_service"
	"github.com/pashura/design-to-wf/api/names_service"
	"github.com/pashura/design-to-wf/api/xtl_structs"
	"reflect"
	"regexp"
	"testing"
)

func init() {
	jackalope_service.TestingMode = true
}

func MockDocumentation(key string) string {
	return "mocked documentation"
}

func MockQualifierDescription(key, qual string) string {
	return "mocked qualifier description"
}

func TestConvertDesignToXtlBaseChecks(t *testing.T) {

	names_service.Documentation = MockDocumentation
	QualifierDescription = MockQualifierDescription
	Documentation = MockDocumentation
	hiddenMeta := design_structs.Schematype{OrgName: "TEST_ORG_NAME"}
	designMeta := design_structs.DesignMeta{HiddenSchema: hiddenMeta}
	testDesign := design_structs.Design{DesignMeta: designMeta}
	testDesign.Children = []design_structs.Object{}
	resultXtl := ConvertDesignToXtl(testDesign, "testOrgName")

	if resultXtl.Infile != "" {
		t.Error(resultXtl.Infile)
	}
	if resultXtl.Repo != "testOrgName.web" {
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
	names_service.Documentation = MockDocumentation
	QualifierDescription = MockQualifierDescription
	Documentation = MockDocumentation
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
	expField.Atts.Mandatory = "N"
	expField.Atts.JavaName = "mockedDocumentation1"
	expField.Atts.Name = "mocked documentation"
	expField.Atts.SegmentTag = "BIG"
	expField.Atts.Position = "01"
	expGroup.Atts.JavaName = "header"
	expGroup.Atts.Name = "Header"
	expField.Name = "FIELDDEF"
	expGroup.Children = []xtl_structs.Element{expField}
	testDesign := design_structs.Design{}
	testDesignElement := design_structs.Object{}
	testDesignElement.Attributes = []design_structs.Object{{}}
	testDesignElement.Attributes[0].ElementType = "restriction"
	testDesignElement.Attributes[0].MinLength = "1"
	testDesignElement.Attributes[0].MaxLength = "8"
	testDesignElement.Sourcing = design_structs.Sourcing{}
	testDesignElement.Sourcing.Location = "Invoice/Header/InvoiceHeader/PurchaseOrderDate"
	testDesignElement.Name = "BIG01"
	testDesignGroup := design_structs.Object{}
	testDesignGroup.MinOccurs = "1"
	testDesignGroup.MaxOccurs = "2"
	testDesignGroup.Name = "Segment-SEG"
	testDesignGroup.Children = []design_structs.Object{testDesignElement}
	testDesign.Children = []design_structs.Object{testDesignGroup}

	structure_levels_service.StructureLevelsFromDesign(testDesign)
	resultXtl := ConvertDesignToXtl(testDesign, "name")

	if !reflect.DeepEqual(resultXtl.Input.Children[0].Children[0], expGroup) {
		t.Errorf("\nactual: %v\n  expected: %v", resultXtl.Input.Children[0].Children[0], expGroup)
	}
}
