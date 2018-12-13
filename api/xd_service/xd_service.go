package xd_service

import (
	"encoding/json"
	"fmt"
	"github.com/pashura/design-to-wf/api/xtl_structs"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "https://xd.spsdev.in/xd/map"

type Payload struct {
	Message    string              `json:"message"`
	Output     xtl_structs.XtlSide  `json:"output,omitempty"`
	Input      xtl_structs.XtlSide `json:"input"`
	Infile     string              `json:"infile"`
	Outfile    string              `json:"outfile"`
	InfileNew  string              `json:"infile_new"`
	InfileRm   bool				   `json:"infile_rm"`
	OutfileNew string             `json:"outfile_new"`
	OutfileRm  bool				   `json:"outfile_rm"`
	Overwrite  bool				   `json:"overwrite"`
	Files      []string			   `json:"files"`
	//ForkedBranch string				`json:"forked_branch"`
}

func XDService(xtl xtl_structs.Xtl, token string) {

	var url = fmt.Sprintf("%s/%s/%s", url, "testDesignToWf.web", "my_awsome_branch")

	payload := &Payload{
		Message:   "Generated from design",
		Input:     xtl.Input,
		Output:    xtl.Input,
		Infile:    "invoiceTest.xtl",
		InfileNew: "invoiceTest.xtl",
		InfileRm:  false,
		OutfileRm: false,
		Overwrite: false,
		Outfile: "invoiceTestIN.xtl",
		OutfileNew: "invoiceTestIN.xtl",
		Files: []string{},
		//ForkedBranch: "master",
	}
	e, err := json.Marshal(payload)

	data := strings.NewReader(string(e))

	req, err := http.NewRequest("POST", url, data)
	req.Header.Add("X-Authorization-Token", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
