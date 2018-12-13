package td_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Design struct {
	HasEnum     bool         `json:"hasEnum"`
	Name        string       `json:"name"`
	ElementType string       `json:"elementType"`
	Visible     bool         `json:"visible"`
	MinOccurs   string       `json:"minOccurs"`
	DesignMeta  DesignMeta   `json:"designMeta"`
	Attributes  []Object     `json:"attributes"`
	Validation  []Validation `json:"validation"`
	Children    []Object     `json:"children"`
}

type Object struct {
	HasEnum             bool                `json:"hasEnum"`
	Base                string              `json:"base"`
	Name                string              `json:"name"`
	ElementType         string              `json:"elementType"`
	DisplayName         string              `json:"displayName"`
	MinOccurs           string              `json:"minOccurs"`
	MinLength           string              `json:"minLength"`
	ID                  int64               `json:"id"`
	Visible             bool                `json:"visible"`
	Qualifiers          string              `json:"qualifiers"`
	EDIid               string              `json:"ediId"`
	Validation          []Validation        `json:"validation"`
	DesignMeta          DesignMeta          `json:"designMeta"`
	QualifierConditions QualifierConditions `json:"qualifierConditions"`
	MaxLength           string              `json:"maxLength"`
	EDIDataType         string              `json:"EDIDataType"`
	DropExtraRecords    bool                `json:"dropExtraRecords"`
	Attributes          []Object            `json:"attributes"`
	Children            []Object            `json:"children"`
}
type DesignMeta struct {
	Tag          string     `json:"tag"`
	HiddenSchema Schematype `json:"hiddenSchema"`
	ViewedSchema Schematype `json:"viewedSchema"`
}

type Schematype struct {
	Source   bool   `json:"source"`
	Version  string `json:"version"`
	Document string `json:"document"`
	Format   string `json:"format"`
	OrgName  string `json:"orgName"`
}

type Validation struct {
	Rules      []Conditions `json:"rules"`
	Type       string       `json:"type"`
	Conditions []Conditions `json:"conditions"`
	Results    []Conditions `json:"results"`
}

type Conditions struct {
	Conjunction            string                   `json:"conjunction"`
	ConditionsInConditions []ConditionsInConditions `json:"conditions"`
}

type ConditionsInConditions struct {
	Condition string `json:"condition"`
	Value     string `json:"value"`
	Element   string `json:"element"`
}

type QualifierConditions struct {
	Converts  string `json:"converts"`
	Qualifier string `json:"qualifier"`
	MinOccurs string `json:"minOccurs"`
}

const url = "https://design-ui-api.dev.spsc.io/company_designs"

func DesignObject(orgId string, designName string, token string) Design {

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

	rawDesign := Design{}
	err = json.Unmarshal([]byte(data.(string)), &rawDesign)
	if err != nil{
		fmt.Println(err.Error())
	}

	return rawDesign
}

func RemoveNonVisible(design Design) Design {
	newDesign := Design{}

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

func createChildren(childrenObject Object) Object {
	newChildrenObject := Object{}

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

func appendChildren(childrenObject []Object) []Object {

	newChildrenObject := []Object{}
	for _, i := range childrenObject {
		if i.Visible == true {
			newChildrenObject = append(newChildrenObject, createChildren(i))
		}
	}
	return newChildrenObject
}
