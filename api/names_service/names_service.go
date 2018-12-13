package names_service

import (
	"github.com/pashura/design-to-wf/api/jackalope_service"
	"unicode"
)
import "strings"

func CreateName(elementName string) string {
	return jackalope_service.Documentation(normalizeElementName(elementName))
}

func CreateJavaName(elementName string) string {
	data := jackalope_service.Documentation(normalizeElementName(elementName))
	javaName := lowerAllLetters(data)
	javaName = upperFirstLetters(javaName)
	javaName = removeSpaces(javaName)
	javaName = lowerInitial(javaName)
	return javaName
}

func normalizeElementName(name string) string {
	name = strings.Replace(name, "Segment-", "", -1)
	name = strings.Replace(name, "Composite-", "", -1)
	name = strings.Replace(name, "Loop-", "", -1)
	return name
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
