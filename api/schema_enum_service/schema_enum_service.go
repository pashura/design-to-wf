package schema_enum_service

import (
	"encoding/json"
	"fmt"
	"github.com/pashura/design-to-wf/api/properties"
	"io/ioutil"
	"net/http"
)

const url = "https://design-ui-api.dev.spsc.io"

type SchemaEnums []struct {
	Enum          string `json:"enum"`
	Documentation string `json:"documentation"`
}

func GetSchemaEnums(groupName string, elementName string, qualifier string) string {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/schema_enums/XSD/%s/%s/%s/%s/%s", url, properties.Format, properties.Version, properties.Document, groupName, elementName), nil)

	if err != nil {
		panic(err.Error())
	}

	req.Header.Add("Authorization", "bearer "+properties.Token)
	req.Header.Set("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
	}

	schemaEnums := SchemaEnums{}

	if fmt.Sprintf("%T", data) == "string" {
		err = json.Unmarshal([]byte(data.(string)), &schemaEnums)
		if err != nil {
			fmt.Println(err.Error())
		}
		desc := getDescription(schemaEnums, qualifier)
		fmt.Println(desc)
		return desc
	}
	return ""
}

func getDescription(schemaEnums SchemaEnums, qualifier string) string {
	for i := 0; i < len(schemaEnums); i++ {
		if schemaEnums[i].Enum == qualifier {
			return schemaEnums[i].Documentation
		}
	}
	return ""
}
