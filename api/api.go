package api

import (
	"fmt"
	"github.com/pashura/design-to-wf/api/design_to_xtl_service"
	"github.com/pashura/design-to-wf/api/design_to_xtl_service/structure_levels_service"
	"github.com/pashura/design-to-wf/api/properties"
	"github.com/pashura/design-to-wf/api/td_service"
	"github.com/pashura/design-to-wf/api/xd_service"
	"github.com/pashura/design-to-wf/api/jackalope_service"
)

var orgId = "320092678795032242107614176381310194095"
var designName = "testyTestTest_test_nikita_RSX_7.7_Invoices_to_X12_4010_Transaction-810"

var repo = "testDesignToWf.web"
var branch = "awesome_branch"

func Run(javaPackageName string) {

	fmt.Println("Fetching design...")
	rawDesign := td_service.DesignObject(orgId, designName, properties.Token)

	properties.Format = rawDesign.DesignMeta.ViewedSchema.Format
	properties.Version = rawDesign.DesignMeta.ViewedSchema.Version
	properties.Document = rawDesign.DesignMeta.ViewedSchema.Document

	properties.SchemaKey = fmt.Sprintf("txn/jackalope/edispec/%v_%v/%v%v_%v.xsd",
		properties.Version, properties.Format, properties.Version, properties.Format,
		properties.Document[len(properties.Document)-3:])

	properties.EnumKey = fmt.Sprintf("XSD/%v/%v/%v.enums",
		properties.Format, properties.Version, properties.Document)

	jackalope_service.S3Service(properties.SchemaKey, properties.EnumKey)

	fmt.Println("Removing non visible elements...")
	design := td_service.RemoveNonVisible(rawDesign)
	structure_levels_service.GetStructureLevelsFromDesign(design)
	des := design_to_xtl_service.ConvertDesignToXtl(design, javaPackageName)

	fmt.Printf("Publishing xtl to branch: %v...", branch)
	xd_service.XDService(des, repo, branch, properties.Token)

}
