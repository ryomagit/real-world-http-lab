package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
)

func main() {
	jar, err := cookiejar.New(nil)	// クッキーを保存するcookiejar(クッキーの瓶）のインスタンスを作成
	if err != nil {
		panic(err)
	}
	client := http.Client{			// クッキーを保存可能なhttp.Clientを作成
		Jar: jar,
	}
	for i := 0; i< 2; i++ {			// クッキーは初回アクセスでCookieをセットし、2回目のアクセスでCookieを送信する
		resp, err := client.Get("http://localhost:18888/cookie")
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}
		log.Println(string(dump))
	}
}