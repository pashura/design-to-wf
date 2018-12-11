package api

import (
	"fmt"
	"github.com/pashura/design-to-wf/api/schema_enum_service"
	"github.com/pashura/design-to-wf/api/td_service"
)

var token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6IjE2MDE4NjA3NTQwNDM1MzEwMTg2NTEwNjcxNTU3NzgzNTYyNDcwNiIsImVtYWlsIjoibXBhc2h1cmFAc3BzY29tbWVyY2UuY29tIiwiZmlyc3RfbmFtZSI6Ik5pa2l0YSIsImxhc3RfbmFtZSI6IlBhc2h1cmEiLCJhdmF0YXJfaW1hZ2VfdXJsIjoiIiwidXNlcl9pZCI6IjEwODQ3MTA3NjczNjA0NjM3MTExNDM0NzQ3MTk4NDI5MjgzNzk0NyIsIm9yZ19pZCI6IjIwOTA5ODgwODMwNTU2MjQ0NjgxMTI2NTM0MTAzMzQ1MzAxNTcwMyIsIm9yZ19uYW1lIjoiU1BTIENvbW1lcmNlIiwiZXhwIjoxNTQ0MDQ1MzA0LCJ2ZXIiOiIxIiwiZW52IjoiZGV2IiwidXJpIjoiaHR0cHM6Ly9kZXYuaWQuc3BzYy5pbyIsImlhdCI6MTU0Mzk1ODkwNH0.RwdwQEzKnMCn5OYGKGYx-5KU7s70R78GmPzVKoK8tWxVroKqL0fQF8W3geVGTOme0uPrHB7YAK9BrPrmSS-irxSNTnOAlmNii3iH4DffVv174037yzm4QUoKXm0_2RncJ_dNc4J0Fs1uQytQ7hqfa85aVEs2Rej7GP-QIZZmXFBBPWsNYGSpaZo3tqa3YRsA4XKs-eryAdAFX-EeIgnLcPVNfJBe52l-JXuEOF7WAbquzVUA48YHIBznLor-lMKl0uH0VElqcuKS1uaM93qZvmbY-n6HpaP3bQphIe94cUWY8ZK1PbQQnZQYFeptdeJ-U-r0HCrQvKrP7Uq_w_KALg"

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


func Run()  {
	fmt.Println("Hello World")
	td_service.GetDesign(orgId, designName, token)
	schema_enum_service.GetSchemaEnums(fileType, version, fileName, groupName, elementName, token)
	schema_enum_service.GetSchemaEnums(fileType, version, fileName, groupSAC, elementSAC01, token)
	schema_enum_service.GetSchemaEnums(fileType, version, fileName, groupBEG, elementBEG02, token)
}
