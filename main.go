package main

import (
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
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
}

func getInput(name string, defaultValues ...string) string {
	if value := os.Getenv(fmt.Sprintf("INPUT_%s", strings.ToUpper(name))); value != "" {
		return value
	}

	if len(defaultValues) == 1 {
		return defaultValues[0]
	}

	log.Fatalf("Error: Input '%s' is required!\n", name)

	return ""
}

func getInputEnum(name string, enum []string, defaultValue ...string) string {
	value := getInput(name, defaultValue...)

	for _, e := range enum {
		if value == e {
			return value
		}
	}

	log.Fatalf("Error: Input '%s' has to be one of [%s]!\n", name, strings.Join(enum, ", "))

	return ""
}
