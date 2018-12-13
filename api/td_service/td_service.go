package td_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Design interface {

}

const url = "https://design-ui-api.dev.spsc.io/company_designs"

func GetDesign(orgId string, designName string, token string) Design {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Companies/%s/Designs/%s.json", url, orgId, designName), nil)

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
