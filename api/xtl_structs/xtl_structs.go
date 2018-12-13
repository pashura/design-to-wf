package xtl_structs

type Xtl struct {
	Repo    string  `json:"repo"`
	Outfile string  `json:"outfile"`
	Branch  string  `json:"branch"`
	Output  XtlSide `json:"output"`
	Input   XtlSide `json:"input"`
	Infile  string  `json:"infile"`
}

type XtlSide struct {
	Atts     Atts           `json:"atts"`
	Name     string         `json:"name"`
	Children [1]DocumentDef `json:"children"`
}

type DocumentDef struct {
	Atts     Atts      `json:"atts"`
	Name     string    `json:"name"`
	Children []Element `json:"children"`
}

type Element struct {
	Atts     Atts      `json:"atts"`
	Name     string    `json:"name"`
	Children []Element `json:"children"`
}

type Atts struct {
	SourceFilter           string `json:"sourceFilter"`
	Enable                 string `json:"enable"`
	Name                   string `json:"name"`
	Min                    string `json:"min"`
	JavaName               string `json:"javaName"`
	IsRecord               string `json:"isRecord"`
	Persistent             string `json:"persistent"`
	Source                 string `json:"source"`
	NextRow                string `json:"nextRow"`
	Justification          string `json:"justification"`
	Max                    string `json:"max"`
	Print                  string `json:"print"`
	Exclude                string `json:"exclude"`
	IncludeInTestFile      string `json:"includeInTestFile"`
	Display                string `json:"display"`
	Present                string `json:"present"`
	Origin                 string `json:"origin"`
	LastModifiedBy         string `json:"lastModifiedBy"`
	Direction              string `json:"direction"`
	MaxSource              string `json:"maxSource"`
	FullyQualifiedJavaName string `json:"fullyQualifiedJavaName"`
	Keys                   string `json:"keys"`
	DesignDate             string `json:"designDate"`
	JavaPackageName        string `json:"javaPackageName"`
	XtencilType            string `json:"xtencilType"`
	SourceTypeDirection    string `json:"sourceTypeDirection"`
	Designerversion        string `json:"designerversion"`
	Owner                  string `json:"owner"`
	Displayer              string `json:"displayer"`
	SourceFiles            string `json:"sourceFiles"`
	Type                   string `json:"type"`
	Revision               string `json:"revision"`
	Mandatory              string `json:"mandatory"`
	Edi                    string `json:"edi"`
	DtdRequired            string `json:"dtdRequired"`
	Rounding               string `json:"rounding"`
	Precision              string `json:"precision"`
	DefaultValue           string `json:"defaultValue"`
	PercentWidth           string `json:"percentWidth"`
	PercentHeight          string `json:"percentHeight"`
	DataType               string `json:"dataType"`
	MinLength              string `json:"minLength"`
	Editable               string `json:"editable"`
	MaxLength              string `json:"maxLength"`
	SegmentTag             string `json:"segmentTag"`
	ReferenceNum           string `json:"referenceNum"`
	Templatable            string `json:"templatable"`
	Position               string `json:"position"`
	KeyType                string `json:"keyType"`
	SubPos                 string `json:"subPos"`
	Conditions             string `json:"conditions"`
	Choices                string `json:"choices"`
	Date                   string `json:"date"`
	FileType               string `json:"fileType"`
	Contents               string `json:"contents"`
}
