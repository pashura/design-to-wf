package api

import (
	"fmt"
	"github.com/pashura/design-to-wf/api/design_to_xtl_service"
	"github.com/pashura/design-to-wf/api/jackalope_service"
	"github.com/pashura/design-to-wf/api/td_service"
	"github.com/pashura/design-to-wf/api/xd_service"
)

var key = "txn/jackalope/edispec/4010_X12/4010X12_850.xsd"

var orgId = "320092678795032242107614176381310194095"
var designName = "testyTestTest_ik_X12_4010_Transaction-850_to_RSX_7.7_Orders"

var repo = "testDesignToWf.web"
var branch = "awesome_branch"

func Run() {

	fmt.Println("Fetching design...")
	rawDesign := td_service.DesignObject(orgId, designName, token)

	jackalope_service.S3Service(s3bucket, key)

	fmt.Println("Removing non visible elements...")
	design := td_service.RemoveNonVisible(rawDesign)
	des := design_to_xtl_service.ConvertDesignToXtl(design)

	xd_service.XDService(des, repo, branch, token)
}
