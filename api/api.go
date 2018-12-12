package api

import (
	"fmt"
	"github.com/pashura/design-to-wf/api/jackalope_service"
	"github.com/pashura/design-to-wf/api/schema_enum_service"
	"github.com/pashura/design-to-wf/api/td_service"
)

var token = ""

var orgId = "320092678795032242107614176381310194095"
var designName = "testyTestTest_test_nikita_RSX_7.7_Invoices_to_X12_4010_Transaction-810"

var fileType = "X12"
var version = "4010"
var fileName = "Transaction-850"
var groupName = "Segment-TD5"
var elementName = "TD501"

var groupSAC = "Segment-SAC"
var elementSAC01 = "SAC01"

var groupBEG = "Segment-BEG"
var elementBEG02 = "BEG02"


var key = "txn/jackalope/edispec/4010_X12/4010X12_850.xsd"
var filename = "api/jackalope_service/4010X12_850.xsd"
var elementID = "355"
var segmentName = "REF"


func Run()  {
	fmt.Println("Hello World")
	td_service.GetDesign(orgId, designName, token)
	schema_enum_service.GetSchemaEnums(fileType, version, fileName, groupName, elementName, token)
	schema_enum_service.GetSchemaEnums(fileType, version, fileName, groupSAC, elementSAC01, token)
	schema_enum_service.GetSchemaEnums(fileType, version, fileName, groupBEG, elementBEG02, token)

	jackalope_service.S3Service(filename, s3bucket, key)
	jackalope_service.Documentation(filename, elementID)

	jackalope_service.Documentation(filename, segmentName)

}
