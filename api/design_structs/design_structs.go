package design_structs

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
	DefaultValue         string             `json:"defaultValue"`
	DisplayName         string              `json:"displayName"`
	MinOccurs           string              `json:"minOccurs"`
	MaxOccurs           string              `json:"maxOccurs"`
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
