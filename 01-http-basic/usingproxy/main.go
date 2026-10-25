package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxyURL, err := url.Parse("http://localhost:18888") // プロキシサーバーのURLを解析
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL), // プロキシサーバーを設定
		},
	}
	resp, err := client.Get("http://github.com") // プロキシサーバー経由でリクエストを送信
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true) 
	if err != nil {
		panic(err)
	}
	log.Println(string(dump)) // レスポンスをログに出力
}
