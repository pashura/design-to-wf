package schema_enum_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SchemaEnums interface {

}


func GetSchemaEnums(fileType string, version string, fileName string, groupName string, elementName string, token string) SchemaEnums {

	baseUrl := "https://design-ui-api.dev.spsc.io"

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/schema_enums/XSD/%s/%s/%s/%s/%s", baseUrl, fileType, version, fileName, groupName, elementName), nil)

	if err != nil {
		panic(err.Error())
	}

	req.Header.Add("Authorization", "bearer " + token)
	req.Header.Set("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var data interface {}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Qualifiers for element %s : %v\n", elementName, data)

	return data
}
