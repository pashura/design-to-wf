package xtl_structs

type Xtl struct {
	Repo    string
	Outfile string
	Branch  string
	Output  XtlSide
	Input   XtlSide
	Infile  string
}

type XtlSide struct {
	Atts     Atts
	Name     string
	Children [1]DocumentDef
}

type DocumentDef struct {
	Atts     Atts
	Name     string
	Children []Element
}

type Element struct {
	Atts     Atts
	Children []Element
	Name     string
}

type Atts struct {
	SourceFilter           string
	Enable                 string
	Name                   string
	Min                    string
	JavaName               string
	IsRecord               string
	Persistent             string
	Source                 string
	NextRow                string
	Justification          string
	Max                    string
	Print                  string
	Exclude                string
	IncludeInTestFile      string
	Display                string
	Present                string
	Origin                 string
	LastModifiedBy         string
	Direction              string
	MaxSource              string
	FullyQualifiedJavaName string
	Keys                   string
	DesignDate             string
	JavaPackageName        string
	XtencilType            string
	SourceTypeDirection    string
	Designerversion        string
	Owner                  string
	Displayer              string
	SourceFiles            string
	Type                   string
	Revision               string
	Mandatory              string
	Edi                    string
	DtdRequired            string
	Rounding               string
	Precision              string
	DefaultValue           string
	PercentWidth           string
	PercentHeight          string
	DataType               string
	MinLength              string
	Editable               string
	MaxLength              string
	SegmentTag             string
	ReferenceNum           string
	Templatable            string
	Position               string
	KeyType                string
	SubPos                 string
	Conditions             string
	Choices                string
	Date                   string
	FileType               string
	Contents               string
}
