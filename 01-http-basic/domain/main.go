package main

import (
	"fmt"

	"golang.org/x/net/idna"
)

func main() {
	src := "握力王"
	ascii, err := idna.ToASCII(src) // Unicode文字列をASCIIに変換
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s -> %s/\n", src, ascii) // 変換結果を出力
}