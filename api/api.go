package api

import (
	"fmt"
	"github.com/pashura/design-to-wf/api/design_to_xtl_service"
	"github.com/pashura/design-to-wf/api/td_service"
)


var orgId = "320092678795032242107614176381310194095"
var designName = "testyTestTest_ik_X12_4010_Transaction-850_to_RSX_7.7_Orders"

func Run() {

	fmt.Println("Fetching design...")
	rawDesign := td_service.DesignObject(orgId, designName, token)

	fmt.Println("Removing non visible elements...")
	design := td_service.RemoveNonVisible(rawDesign)
	design_to_xtl_service.ConvertDesignToXtl(design)
}
