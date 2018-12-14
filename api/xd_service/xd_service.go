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
	Output     map[string]int      `json:"output"`
	Input      xtl_structs.XtlSide `json:"input"`
	Infile     string              `json:"infile"`
	Outfile    string              `json:"outfile"`
	InfileNew  string              `json:"infile_new"`
	InfileRm   bool                `json:"infile_rm"`
	OutfileNew string              `json:"outfile_new"`
	OutfileRm  bool                `json:"outfile_rm"`
	Overwrite  bool                `json:"overwrite"`
	Files      []string            `json:"files,omitempty"`
}

func XDService(xtl xtl_structs.Xtl, repo string, branch string, token string) {

	var url = fmt.Sprintf("%s/%s/%s", url, repo, branch)

	payload := &Payload{
		Message:    "Generated from design",
		Input:      xtl.Input,
		Output:     make(map[string]int, 1),
		Infile:     "invoiceTestIN.xtl",
		InfileNew:  "invoiceTestIN.xtl",
		InfileRm:   false,
		OutfileRm:  false,
		Overwrite:  false,
		Outfile:    "",
		OutfileNew: "",
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
