package rep_service

import (
	"github.com/pashura/design-to-wf/api/xtl_structs"
	"reflect"
	"testing"
)

func TestAddReps(t *testing.T) {
	xtl := xtl_structs.Xtl{}
	xtl.Input.Name = "SPSFILE"

	expXtl := xtl_structs.Xtl{}
	expXtl.Input.Name = "SPSFILE"

	docDef := [1]xtl_structs.DocumentDef{}
	docDef[0].Name = "DOCUMENTDEF"

	header := createGroup("1", "1", "Header", "header")
	invoiceNumber := createElement("1", "22", "Invoice Number", "invoiceNumber")
	address := createGroup("0", "100", "Address", "address")
	addressRep := createGroup("1", "1", "Address Rep", "addressRep")
	addressTypeCode := createElement("2", "4", "Address Type Code", "addressTypeCode")
	contact := createGroup("0", "100", "Contact", "contact")
	contactRep := createGroup("1", "1", "Contact Rep", "contactRep")
	contactTypeCode := createElement("2", "4", "Contact Type Code", "contactTypeCode")
	lineItem := createGroup("1", "100", "Line Item", "lineItem")
	lineItemRep := createGroup("1", "1", "Line Item Rep", "lineItemRep")
	lineItemNumber := createElement("1", "20", "Line Item Number", "lineItemNumber")
	physicalDetails := createGroup("1", "1", "Physical Details", "physicalDetails")
	quantity := createElement("1", "10", "Quantity", "quantity")
	summary := createGroup("1", "1", "Summary", "summary")
	totalAmount := createElement("1", "15", "Total Amount", "totalAmount")

	summary.Children = []xtl_structs.Element{totalAmount}
	physicalDetails.Children = []xtl_structs.Element{quantity}
	lineItem.Children = []xtl_structs.Element{lineItemNumber, physicalDetails}
	contact.Children = []xtl_structs.Element{contactTypeCode}
	address.Children = []xtl_structs.Element{addressTypeCode, contact}
	header.Children = []xtl_structs.Element{invoiceNumber, address}
	docDef[0].Children = []xtl_structs.Element{header, lineItem, summary}
	xtl.Input.Children = docDef

	AddReps(xtl)

	summary.Children = []xtl_structs.Element{totalAmount}
	physicalDetails.Children = []xtl_structs.Element{quantity}
	lineItemRep.Children = []xtl_structs.Element{lineItemNumber, physicalDetails}
	lineItem.Children = []xtl_structs.Element{lineItemRep}
	contactRep.Children = []xtl_structs.Element{contactTypeCode}
	contact.Children = []xtl_structs.Element{contactRep}
	addressRep.Children = []xtl_structs.Element{addressTypeCode, contact}
	address.Children = []xtl_structs.Element{addressRep}
	header.Children = []xtl_structs.Element{invoiceNumber, address}
	docDef[0].Children = []xtl_structs.Element{header, lineItem, summary}
	expXtl.Input.Children = docDef

	if !reflect.DeepEqual(xtl, expXtl) {
		t.Errorf("\nactual: %v\n  expected: %v", xtl, expXtl)
	}
}

func createGroup(min, max, name, javaName string) xtl_structs.Element {
	group := xtl_structs.Element{}
	group.Name = "GROUPDEF"
	group.Atts.Min = min
	group.Atts.Max = max
	group.Atts.Name = name
	group.Atts.JavaName = javaName
	return group
}

func createElement(minLength, maxLength, name, javaName string) xtl_structs.Element {
	element := xtl_structs.Element{}
	element.Name = "FIELDDEF"
	element.Atts.MinLength = minLength
	element.Atts.MaxLength = maxLength
	element.Atts.Name = name
	element.Atts.JavaName = javaName
	return element
}
