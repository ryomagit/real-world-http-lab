package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"test": {"values"},
	}
	
	rest, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Status Code:", rest.StatusCode)
}