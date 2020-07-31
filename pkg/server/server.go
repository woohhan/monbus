package server

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"log"
	"net/http"
)

type SkillRequest struct {
	Action SkillBodyAction `json:"action"`
}

type SkillBodyAction struct {
	Name         string       `json:"name"`
	DetailParams DetailParams `json:"detailParams"`
}

type DetailParams struct {
	IsWeekend isWeekend `json:"IsWeekend"`
}

type isWeekend struct {
	Value bool `json:"value"`
}

type SkillResponse struct {
	Version  string   `json:"version"`
	Template Template `json:"template"`
}

type Template struct {
	Outputs []SkillOutput `json:"outputs"`
}

type SkillOutput struct {
	SimpleText SimpleText `json:"simpleText"`
}

type SimpleText struct {
	Text string `json:"text"`
}

func Test() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hi"))
	})
	http.HandleFunc("/bustime", busTimeHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func busTimeHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("start busTimeHandler with %v", r)

	if r.Method != "POST" {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
	}

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}
	glog.Infof("--- body ---\n%s\n------------", body)

	// parse body
	sbody := &SkillRequest{}
	json.Unmarshal(body, &sbody)
	fmt.Printf("sbody\n%v", sbody)

	resp := &SkillResponse{
		Version: "2.0",
		Template: Template{
			Outputs: []SkillOutput{
				{
					SimpleText: SimpleText{
						Text: "간단한 내용",
					},
				},
			},
		},
	}

	j, _ := json.Marshal(resp)
	fmt.Println(string(j))

	w.Write(j)
}
