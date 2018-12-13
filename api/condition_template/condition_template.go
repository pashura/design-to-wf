package condition_template

import "fmt"

type RootCode struct {
	code string
}

type BodyCode struct {
	code   string
	events []string
}

func InitialRoots() {
	optionalGroupMandatoryElementsRoot := RootCode{"public void OptionalGroupMandatoryElements(JEntity ent,String conditionalElements,JElement... bar){ \n\tString x = \"N\";\n\n\tfor(int i=0;i<ent.getChildrenCount();i++){\n\t\tif(ent.getJChild(i) instanceof JElement){\n\t\t\tJElement element = (JElement)ent.getJChild(i);\n\t\t\tString javaName = (String)element.getProperty(\"javaName\");\n\n\t\tif( element.hasData() && (javaName.matches(conditionalElements) || conditionalElements.length()==0)){\n\t\t\t\tx = \"Y\";\n\t\t\t\tbreak;\n\t\t\t}\n\t\t}\n\t}\n\n\tfor (JElement baz : bar){\n\t\tbaz.setProperty(\"mandatory\",x);\n\t}"}
	fmt.Println(optionalGroupMandatoryElementsRoot)
	conditionCRoot := RootCode{"public void conditionC(JEntity ent, String list, String AllorOne, JElement... bar){\n\tboolean enabled = false;\n\tint hasData = 0;\n\tint elCount = 0;\n\tString enElList = \" [\";\n\tString elList = \" [\";\n\tString[] x = list.split(\",\");\n\n\tfor(int i=0;i<x.length;i++){\n\t\tJElement t = (JElement)ent.getJChild(x[i]);\n\t\tenElList += t.getProperty(\"name\")+\"], [\";\n\t\tif( t.hasData()) enabled=true;\n\t}\n\tfor (JElement baz : bar){\n\t\telCount++;\n\t\telList += baz.getProperty(\"name\")+\"], [\";\n\t\tif(baz.hasData()) hasData++;\n\t}\n\tif(AllorOne.equals(\"All\") && enabled && hasData<elCount) ent.reportError(ent,\"if \" + ( x.length>1?\"at least one of \":\"\" )  + enElList.substring(0,enElList.length()-3) + \" is present then \" + elList.substring(0,elList.length()-3)+ ( elCount>1?\" are\":\" is\" ) +\" required !\" );\n\tif(AllorOne.equals(\"One\") && enabled && hasData==0)     ent.reportError(ent,\"if \" + ( x.length>1?\"at least one of \":\"\" )  + enElList.substring(0,enElList.length()-3) + \" is present then \" + ( elCount>1?\" at least one of\":\"\" ) + elList.substring(0,elList.length()-3)+ \" is required !\" );\n}"}
	fmt.Println(conditionCRoot)
}

func InitialBodies() {
	ifOneThenAllBody := BodyCode{"root.OptionalGroupMandatoryElements(me,\"{{.JavaNames}}\",{{.JavaNames}});\nrefresh(me);",
		[]string{"documentCheck", "entityEdited"}}
	fmt.Println(ifOneThenAllBody)
	ifThenBody := BodyCode{"root.OptionalGroupMandatoryElements(me,\"{{.JavaNames}}\",{{.JavaNames}});\nrefresh(me);",
		[]string{"documentCheck", "entityEdited"}}
	fmt.Println(ifThenBody)
	atLeastOneOfBody := BodyCode{"root.OptionalGroupMandatoryElements(me,\"{{.JavaNames}}\",{{.JavaNames}});\nrefresh(me);",
		[]string{"documentCheck"}}
	fmt.Println(atLeastOneOfBody)
	onlyOneOfBody := BodyCode{"root.OptionalGroupMandatoryElements(me,\"{{.JavaNames}}\",{{.JavaNames}});\nrefresh(me);",
		[]string{"documentCheck"}}
	fmt.Println(onlyOneOfBody)
}
