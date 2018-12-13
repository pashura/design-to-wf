package jackalope_service

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

type Schema struct {
	XMLName xml.Name `xml:"schema"`
	Element   []Element   `xml:"element"`
}

type Element struct {
	XMLName xml.Name `xml:"element"`
	Name    string   `xml:"name,attr"`
	Annotation   []Annotation   `xml:"annotation"`
}

type Annotation struct {
	XMLName xml.Name `xml:"annotation"`
	Documentation   []string   `xml:"documentation"`
}


func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}


func Documentation(filename string, ref string) string {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var schema Schema

	xml.Unmarshal(byteValue, &schema)

	var searchedName string

	if isInt(ref){
		searchedName = fmt.Sprintf("Element-%s", ref)
	} else {
		searchedName = fmt.Sprintf("Segment-%s", ref)
	}


	for i := 0; i < len(schema.Element); i++ {
		currentElement := schema.Element[i]


		if currentElement.Name == searchedName {
			fmt.Printf("Name for reference element ID %v is: %v \n", ref, currentElement.Annotation[0].Documentation[0])
			return currentElement.Annotation[0].Documentation[0]
		}
	}
	fmt.Printf("Couldn't find name for reference: %v \n", ref)
	return ""
}
