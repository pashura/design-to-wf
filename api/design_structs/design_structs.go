package design_structs

import "strings"

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
	HasEnum             bool                 `json:"hasEnum"`
	Base                string               `json:"base"`
	Name                string               `json:"name"`
	ElementType         string               `json:"elementType"`
	DefaultValue        string               `json:"defaultValue"`
	DisplayName         string               `json:"displayName"`
	MinOccurs           string               `json:"minOccurs"`
	MaxOccurs           string               `json:"maxOccurs"`
	MinLength           string               `json:"minLength"`
	ID                  int64                `json:"id"`
	Visible             bool                 `json:"visible"`
	Sourcing 			Sourcing 			 `json:"sourcing"`
	Qualifiers          string               `json:"qualifiers"`
	EDIid               string               `json:"ediId"`
	Validation          []Validation         `json:"validation"`
	DesignMeta          DesignMeta           `json:"designMeta"`
	QualifierConditions []QualifierCondition `json:"qualifierConditions"`
	MaxLength           string               `json:"maxLength"`
	EDIDataType         string               `json:"EDIDataType"`
	DropExtraRecords    bool                 `json:"dropExtraRecords"`
	Attributes          []Object             `json:"attributes"`
	Children            []Object             `json:"children"`
}

type Sourcing struct {
	Location		string		`json:"location"`
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

type QualifierCondition struct {
	Converts  string `json:"converts"`
	Qualifier string `json:"qualifier"`
	MinOccurs int    `json:"minOccurs"`
}

func (o Object) GetSegmentName() string {
	if strings.Contains(o.Name, "-"){
		 return o.Name[strings.Index(o.Name,"-")+1:]
	}
	return ""
}