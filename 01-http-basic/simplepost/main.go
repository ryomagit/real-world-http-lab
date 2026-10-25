package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// 実行場所にある main.go を送信する
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status Code:", resp.StatusCode)
}