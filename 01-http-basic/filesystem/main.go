package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	transport := &http.Transport{} // デフォルトのTransportを作成
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("."))) // fileスキームのURLを処理するためのTransportを登録
	client := http.Client{
		Transport: transport, // Transportを設定
	}
	resp, err := client.Get("file://./main.go") // fileスキームのURLを指定してリクエストを送信
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true) // レスポンスをダンプ
	if err != nil {
		panic(err)
	}
	log.Println(string(dump)) 
}