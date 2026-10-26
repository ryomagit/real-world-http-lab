# Real World HTTP Lab

『Real World HTTP 第3版』の内容を実際に動かしながら、HTTPの仕組みを理解するための学習用リポジトリ。

## 学習方針

1. 書籍のサンプルコードを実装・実行する
2. curlなどを使ってHTTP通信を確認する
3. 必要に応じてコードを変更し、動作を検証する
4. 学んだことや気づきを記録する

## ディレクトリ構成

リポジトリ内の番号は、書籍の章番号ではなく学習テーマの連番とする。

| ディレクトリ | 対応する書籍の章 | テーマ |
| --- | --- | --- |
| [`01-http-basic`](./01-http-basic/) | 第1〜4章 | ブラウザ、HTTP/1.0、Goによる基本的なHTTPクライアント |
| [`02-http11-syntax`](./02-http11-syntax/) | 第5章 | HTTP/1.1のシンタックス |
| [`03-http11-semantics`](./03-http11-semantics/) | 第6章 | HTTP/1.1時代のセマンティクス |
| [`04-go-http11-client`](./04-go-http11-client/) | 第7章 | GoによるHTTP/1.1クライアント |
| [`05-http2-http3-syntax`](./05-http2-http3-syntax/) | 第8章 | HTTP/2、HTTP/3のシンタックス |
| [`06-http2-use-cases`](./06-http2-use-cases/) | 第9章 | HTTP/2時代のユースケース |
| [`07-go-modern-http`](./07-go-modern-http/) | 第10章 | GoによるHTTP/2、HTTP/3、HTML 5プロトコルの実装 |
| [`08-restful-api`](./08-restful-api/) | 第11章 | クライアント視点のRESTful API |
| [`09-browser-dynamic-http`](./09-browser-dynamic-http/) | 第12章 | JavaScriptによる動的なHTTPリクエスト |
| [`10-web-application-basics`](./10-web-application-basics/) | 第13章 | ウェブアプリケーションの基礎 |
| [`11-cloud-http`](./11-cloud-http/) | 第14章 | クラウド時代のHTTP |
| [`12-web-platformization`](./12-web-platformization/) | 第15章 | ウェブのプラットフォーム化 |
| [`13-web-security`](./13-web-security/) | 第16章 | ブラウザを守るHTTPのセキュリティ機能 |
| [`14-http-reference`](./14-http-reference/) | 付録 | ステータスコードとフィールドのリファレンス |

すべてのGoサンプルは、原則としてリポジトリ直下の`go.mod`を共用する。

## 参考書籍

- 『Real World HTTP 第3版』（オライリー・ジャパン）
