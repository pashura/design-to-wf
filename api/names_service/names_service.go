package names_service

import (
	"github.com/pashura/design-to-wf/api/jackalope_service"
	"regexp"
	"strconv"
	"unicode"
)
import "strings"

var elements = make(map[string]int)
var currentSegment string


func CreateName(elementName string) string{
	return jackalope_service.Documentation(normalizeElementName(elementName) )
}

func CreateJavaName(elementName string, currentSegment string) string {
	data := jackalope_service.Documentation(normalizeElementName(elementName) )
	javaName := removeNonLiteralSymbols(data)
	javaName = lowerAllLetters(javaName)
	javaName = upperFirstLetters(javaName)
	javaName = removeSpaces(javaName)
	javaName = lowerInitial(javaName)
	javaName = addUniqueIdItoNameIfNeeded(javaName, currentSegment)
	return javaName
}

func normalizeElementName(name string) string{
	name = strings.Replace(name,"Segment-","", -1)
	name = strings.Replace(name,"Composite-","", -1)
	name = strings.Replace(name,"Loop-","", -1)
	return name
}

func removeNonLiteralSymbols(str string) string{
	re := regexp.MustCompile("[ 0-9a-zA-Z]+")
	return strings.Join(re.FindAllString(str, -1), " ")
}


func lowerInitial(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func lowerAllLetters(str string) string {
	return strings.ToLower(str)
}

func upperFirstLetters(str string) string {
	return strings.Title(str)
}

func removeSpaces(str string) string {
	return strings.Replace(str, " ", "", -1)
}

func addUniqueIdItoNameIfNeeded(elementName string, segmentName string) string{
	if currentSegment != segmentName{
		elements = make(map[string]int)
		currentSegment = segmentName
	}

	if id, ok := elements[elementName]; ok {
		elements[elementName] = id+1
		elementName = elementName+strconv.Itoa(id+1)
	} else {
		elements[elementName] = 0
	}
	return elementName
}
