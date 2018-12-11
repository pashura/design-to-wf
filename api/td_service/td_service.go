package td_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Design interface {

}

func GetDesign(orgId string, designName string, token string) Design {

	baseUrl := "https://design-ui-api.dev.spsc.io/company_designs"

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Companies/%s/Designs/%s.json", baseUrl, orgId, designName), nil)

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
	fmt.Printf("Design for %s: %v\n", designName, data)

	return data
}


//func MarshalOnlyFields(structa interface{},
//	includeFields map[string]bool) (jsona []byte, status error) {
//	value := reflect.ValueOf(structa)
//	typa := reflect.TypeOf(structa)
//	size := value.NumField()
//	jsona = append(jsona, '{')
//	for i := 0; i < size; i++ {
//		structValue := value.Field(i)
//		var fieldName string = typa.Field(i).Name
//		if marshalledField, marshalStatus := json.Marshal((structValue).Interface()); marshalStatus != nil {
//			return []byte{}, marshalStatus
//		} else {
//			if includeFields[fieldName] {
//				jsona = append(jsona, '"')
//				jsona = append(jsona, []byte(fieldName)...)
//				jsona = append(jsona, '"')
//				jsona = append(jsona, ':')
//				jsona = append(jsona, (marshalledField)...)
//				if i+1 != len(includeFields) {
//					jsona = append(jsona, ',')
//				}
//			}
//		}
//	}
//	jsona = append(jsona, '}')
//	return
//}