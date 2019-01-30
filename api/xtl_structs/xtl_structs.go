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
	Children []Element `json:"children,omitempty"`
}

type Atts struct {
	SourceFilter           string `json:"sourceFilter,omitempty"`
	Enable                 string `json:"enable,omitempty"`
	Name                   string `json:"name,omitempty"`
	Min                    string `json:"min,omitempty"`
	JavaName               string `json:"javaName,omitempty"`
	IsRecord               string `json:"isRecord,omitempty"`
	Persistent             string `json:"persistent,omitempty"`
	Source                 string `json:"source,omitempty"`
	NextRow                string `json:"nextRow,omitempty"`
	Justification          string `json:"justification,omitempty"`
	Max                    string `json:"max,omitempty"`
	Print                  string `json:"print,omitempty"`
	Exclude                string `json:"exclude,omitempty"`
	IncludeInTestFile      string `json:"includeInTestFile,omitempty"`
	Display                string `json:"display,omitempty"`
	Present                string `json:"present,omitempty"`
	Origin                 string `json:"origin,omitempty"`
	LastModifiedBy         string `json:"lastModifiedBy,omitempty"`
	Direction              string `json:"direction,omitempty"`
	MaxSource              string `json:"maxSource,omitempty"`
	FullyQualifiedJavaName string `json:"fullyQualifiedJavaName,omitempty"`
	Keys                   string `json:"keys,omitempty"`
	DesignDate             string `json:"designDate,omitempty"`
	JavaPackageName        string `json:"javaPackageName,omitempty"`
	XtencilType            string `json:"xtencilType,omitempty"`
	SourceTypeDirection    string `json:"sourceTypeDirection,omitempty"`
	Designerversion        string `json:"designerversion,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	Displayer              string `json:"displayer,omitempty"`
	SourceFiles            string `json:"sourceFiles,omitempty"`
	Type                   string `json:"type,omitempty"`
	Revision               string `json:"revision,omitempty"`
	Mandatory              string `json:"mandatory,omitempty"`
	Edi                    string `json:"edi,omitempty"`
	DtdRequired            string `json:"dtdRequired,omitempty"`
	Rounding               string `json:"rounding,omitempty"`
	Precision              string `json:"precision,omitempty"`
	DefaultValue           string `json:"defaultValue,omitempty"`
	PercentWidth           string `json:"percentWidth,omitempty"`
	PercentHeight          string `json:"percentHeight,omitempty"`
	DataType               string `json:"dataType,omitempty"`
	MinLength              string `json:"minLength,omitempty"`
	Editable               string `json:"editable,omitempty"`
	MaxLength              string `json:"maxLength,omitempty"`
	SegmentTag             string `json:"segmentTag,omitempty"`
	ReferenceNum           string `json:"referenceNum,omitempty"`
	Templatable            string `json:"templatable,omitempty"`
	Position               string `json:"position,omitempty"`
	KeyType                string `json:"keyType,omitempty"`
	SubPos                 string `json:"subPos,omitempty"`
	Conditions             string `json:"conditions,omitempty"`
	Choices                string `json:"choices,omitempty"`
	Date                   string `json:"date,omitempty"`
	FileType               string `json:"fileType,omitempty"`
	Contents               string `json:"contents,omitempty"`
}
