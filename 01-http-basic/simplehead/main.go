package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		log.Panic(err)
	}
	log.Println("Status Code:", resp.StatusCode)
	log.Println("Headers:", resp.Header)
}