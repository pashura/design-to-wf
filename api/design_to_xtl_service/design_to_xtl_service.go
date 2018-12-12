package design_to_xtl_service

type Xtl struct{
	repo string
	outfile string
	branch string
	output XtlSide
	input XtlSide
	infile string
}

type XtlSide struct {
	atts Atts
	name string
	children [1]RootElement
}

type RootElement struct {
	atts Atts
	name string
	children []Element
}

type Element struct {
	atts Atts
	children []Element
	name string
}

type Atts struct {
	sourceFilter string
	enable string
	name string
	min string
	javaName string
	isRecord string
	persistent string
	source string
	nextRow string
	justification string
	max string
	print string
	exclude string
	includeInTestFile string
	display string
	present string
	origin string
	lastModifiedBy string
	direction string
	maxSource string
	fullyQualifiedJavaName string
	keys string
	designDate string
	javaPackageName string
	xtencilType string
	sourceTypeDirection string
	designerversion string
	owner string
	displayer string
	sourceFiles string
	type_ string
	revision string
	mandatory string
	edi string
	dtdRequired string
	rounding string
	precision string
	defaultValue string
	percentWidth string
	percentHeight string
	dataType string
	minLength string
	editable string
	maxLength string
	segmentTag string
	referenceNum string
	templatable string
	position string
	keyType string
	subPos string
	conditions string
	choices string
	date string
	fileType string
	contents string
}
