package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/josa42/gh-actions-toolkit/core"
)

var methods = []string{"GET", "POST", "PUT", "DELETE"}

func main() {
	url := core.GetInput("url", core.Options{Required: true})
	method := core.GetInput("method", core.Options{Enum: methods, Default: "POST"})

	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Panicln(err)
	}

	if data := core.GetInput("data", core.Options{Required: true}); data != "" {
		if isJSON(data) {
			req.Header.Set("Content-Type", "application/json")
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(data)))
	}

	if auth := core.GetInput("authorization", core.Options{}); auth != "" {
		req.Header.Set("Authorization", auth)
	}

	if contentType := core.GetInput("content_type", core.Options{}); contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	if accept := core.GetInput("accept", core.Options{}); accept != "" {
		req.Header.Set("Accept", accept)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

func isJSON(s string) bool {
	var arr []interface{}
	var obj map[string]interface{}
	var str string

	check := func(d interface{}) bool {
		return json.Unmarshal([]byte(s), d) == nil
	}

	return check(&obj) || check(&str) || check(&arr)
}
