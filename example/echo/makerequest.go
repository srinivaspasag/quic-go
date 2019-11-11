package main

import (
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	MakeRequest()
}

func MakeRequest() {
	//resp, err := http.Get("https://httpbin.org/get")
	resp, err := http.Get("http://localhost:8000/en-US/services/collector/event")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
