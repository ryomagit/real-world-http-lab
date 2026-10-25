package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"query": {"hellow world"},
	}
	resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
	log.Println("Status Code:", resp.StatusCode)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println("Fields:", resp.Header)
	log.Println(string(body))
}