package td_service

import (
	"encoding/json"
	"fmt"
	"github.com/pashura/design-to-wf/api/design_structs"
	"io/ioutil"
	"net/http"
)


const url = "https://design-ui-api.dev.spsc.io/company_designs"

func DesignObject(orgId string, designName string, token string) design_structs.Design {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Companies/%s/Designs/%s.json", url, orgId, designName), nil)
	if err != nil{
		fmt.Println(err.Error())
	}

	req.Header.Add("Authorization", "bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println(err.Error())
	}

	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil{
		fmt.Println(err.Error())
	}

	rawDesign := design_structs.Design{}
	err = json.Unmarshal([]byte(data.(string)), &rawDesign)
	if err != nil{
		fmt.Println(err.Error())
	}

	return rawDesign
}

func RemoveNonVisible(design design_structs.Design) design_structs.Design {
	newDesign := design_structs.Design{}

	newDesign.Name = design.Name
	newDesign.HasEnum = design.HasEnum
	newDesign.ElementType = design.ElementType
	newDesign.DesignMeta = design.DesignMeta
	newDesign.Visible = design.Visible
	newDesign.MinOccurs = design.MinOccurs
	newDesign.Attributes = design.Attributes
	newDesign.Children = appendChildren(design.Children)

	return newDesign
}

func createChildren(childrenObject design_structs.Object) design_structs.Object {
	newChildrenObject := design_structs.Object{}

	newChildrenObject.Visible = childrenObject.Visible
	newChildrenObject.ElementType = childrenObject.ElementType
	newChildrenObject.HasEnum = childrenObject.HasEnum
	newChildrenObject.DesignMeta = childrenObject.DesignMeta
	newChildrenObject.Attributes = childrenObject.Attributes
	newChildrenObject.MinOccurs = childrenObject.MinOccurs
	newChildrenObject.Name = childrenObject.Name
	newChildrenObject.Base = childrenObject.Base
	newChildrenObject.DisplayName = childrenObject.DisplayName
	newChildrenObject.DropExtraRecords = childrenObject.DropExtraRecords
	newChildrenObject.EDIDataType = childrenObject.EDIDataType
	newChildrenObject.EDIid = childrenObject.EDIid
	newChildrenObject.ID = childrenObject.ID
	newChildrenObject.MinLength = childrenObject.MinLength
	newChildrenObject.MaxLength = childrenObject.MaxLength
	newChildrenObject.Attributes = childrenObject.Attributes
	newChildrenObject.QualifierConditions = childrenObject.QualifierConditions
	newChildrenObject.Qualifiers = childrenObject.Qualifiers
	newChildrenObject.Children = appendChildren(childrenObject.Children)

	fmt.Println(newChildrenObject)

	return newChildrenObject
}

func appendChildren(childrenObject []design_structs.Object) []design_structs.Object {

	var newChildrenObject []design_structs.Object
	for _, i := range childrenObject {
		if i.Visible == true {
			newChildrenObject = append(newChildrenObject, createChildren(i))
		}
	}
	return newChildrenObject
}
