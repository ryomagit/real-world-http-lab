package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:18888", nil) // DELETEメソッドのリクエストを作成
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req) // リクエストを送信
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}