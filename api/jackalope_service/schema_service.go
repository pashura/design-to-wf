package jackalope_service

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

var TestingMode bool

type xmlSchema struct {
	XMLName xml.Name     `xml:"schema"`
	Element []xmlElement `xml:"element"`
}

type xmlElement struct {
	XMLName    xml.Name     `xml:"element"`
	Name       string       `xml:"name,attr"`
	Annotation []annotation `xml:"annotation"`
}

type annotation struct {
	XMLName       xml.Name `xml:"annotation"`
	Documentation []string `xml:"documentation"`
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func Documentation(ref string) string {
	var filename string
	if TestingMode == true {
		filename = "xml_parser_test_resources.xml"
	} else {
		filename = "api/jackalope_service/schema.xsd"
	}

	xmlFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var schema xmlSchema

	if err := xml.Unmarshal(byteValue, &schema); err != nil {
		panic(err)
	}

	var searchedName string

	if isInt(ref) {
		searchedName = fmt.Sprintf("Element-%s", ref)
	} else if strings.HasPrefix(ref, "Transaction") {
		searchedName = ref
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

type Enums []struct {
	Enum          string `json:"enum"`
	Documentation string `json:"documentation"`
}

func QualifierDescription(name, qualifier string) string {

	var filename string
	if TestingMode == true {
		filename = "enums_test_resources.json"
	} else {
		filename = "api/jackalope_service/enums.json"
	}

	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var dat map[string]interface{}
	if err := json.Unmarshal(byteValue, &dat); err != nil {
		panic(err)
	}

	a := fmt.Sprintf("Segment-%v", name[:len(name)-2])
	if segment, ok := dat[a].(map[string]interface{}); ok {
		var el = segment[name]

		// convert map to json
		jsonString, _ := json.Marshal(el)

		// convert json to struct
		s := Enums{}
		if err := json.Unmarshal(jsonString, &s); err != nil {
			panic(err)
		}

		for _, enum := range s {
			if enum.Enum == qualifier {
				fmt.Println(name, qualifier, "-", enum.Documentation)
				return enum.Documentation
			}
		}

	}

	return ""
}
