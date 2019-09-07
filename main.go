package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var methods = []string{"GET", "POST"}

func main() {
	url := getInput("url")
	method := getInputEnum("method", methods, "POST")

	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Panicln(err)
	}

	if data := getInput("data", ""); data != "" {
		if isJSON(data) {
			req.Header.Set("Content-Type", "application/json")
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(data)))
	}

	if auth := getInput("authorization", ""); auth != "" {
		req.Header.Set("Authorization", auth)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

func getInput(name string, defaultValues ...string) string {
	if value := os.Getenv(fmt.Sprintf("INPUT_%s", strings.ToUpper(name))); value != "" {
		return value
	}

	if len(defaultValues) == 1 {
		return defaultValues[0]
	}

	log.Panicf("Error: Input '%s' is required!\n", name)

	return ""
}

func getInputEnum(name string, enum []string, defaultValue ...string) string {
	value := getInput(name, defaultValue...)

	for _, e := range enum {
		if value == e {
			return value
		}
	}

	log.Panicf("Error: Input '%s' has to be one of [%s]!\n", name, strings.Join(enum, ", "))

	return ""
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
