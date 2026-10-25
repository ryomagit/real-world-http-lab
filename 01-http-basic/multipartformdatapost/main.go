package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var buffer bytes.Buffer // multipart/form-dataのリクエストボディを作成するためのバッファ
	writer := multipart.NewWriter(&buffer) // multipart/form-dataのリクエストボディを作成するためのライター
	writer.WriteField("name", "test") // ファイル以外のフォームフィールド

	// 個別のファイル書き込みio.Writerを作成する
	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	// ファイルの内容をfileWriterにコピーする
	io.Copy(fileWriter, readFile)
	writer.Close() // バッファに書き込み

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status Code:", resp.StatusCode)
}