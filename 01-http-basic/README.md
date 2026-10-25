f## 概要

## ディレクトリ構成

- server.go
  テストサーバー。
  `go run server.go`
- test.txt
  サンプルファイル。

## curlコマンド

### オプション

- `-d` Content-Type: application/x-www-form-urlencodedを設定
- `-F` Content-Type: multipart/form-dataを設定

### コマンド例

- `curl --http1.0 -F title="the art of community" -F author="Jono Bacon" -F attachment-file=@Test.txt http://localhost:18888`
  `Content-Type` = multipart/form-data
  フォームを使ったファイル送信を再現
  ```
  <form action="POST" enctype="multipart/form-data'>
  </form>
  ```
