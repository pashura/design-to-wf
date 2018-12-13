package xd_service

import (
	"encoding/json"
	"fmt"
	"github.com/pashura/design-to-wf/api/design_to_xtl_service"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "https://xd.spsdev.in/xd/map"

const token = ""

type Payload struct {
	Message    string                        `json:"message"`
	Output     design_to_xtl_service.XtlSide `json:"output"`
	Input      string                        `json:"input"`
	Infile     string                        `json:"infile"`
	Outfile    string                        `json:"outfile"`
	InfileNew  string                        `json:"infile_new"`
	OutfileNew string                        `json:"outfile_new"`
}

func XDService(xtl design_to_xtl_service.XtlSide, repo string, branch string, filename string) {

	var url = fmt.Sprintf("%s/%s/%s", url, repo, branch)

	payload := &Payload{
		Message:   "Generated from design",
		Input:     "",
		Output:    xtl,
		Infile:    "",
		InfileNew: "",
		Outfile:   filename,
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
