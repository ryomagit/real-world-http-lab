package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	var buffer bytes.Buffer // multipart/form-dataのリクエストボディを作成するためのバッファ
	writer := multipart.NewWriter(&buffer) // multipart/form-dataのリクエストボディを作成するためのライター
	writer.WriteField("name", "test") // ファイル以外のフォームフィールド

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`) // ファイルのMIMEヘッダーを作成する
	fileWriter, err := writer.CreatePart(part) // ファイルのフォームフィールドを作成する
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open("photo.jpg") // ファイルを開く
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