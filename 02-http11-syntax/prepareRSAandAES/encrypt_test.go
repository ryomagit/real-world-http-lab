package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"io"
	"testing"
)

// 共通鍵（AES）と公開鍵（RSA）の暗号化・復号化を比較するベンチマーク。
//
// 実行方法（このファイルがあるディレクトリで実行）:
//
// 	go test -bench .
//
// 「.」は、すべてのBenchmarkXxx関数を実行する指定。
//
// 出力例:
//
// 	BenchmarkRSAEncryption-8    36368    32827 ns/op
//
// 出力の見方:
//   - BenchmarkRSAEncryption-8: ベンチマーク名。末尾の8はGOMAXPROCSの値
//   - 36368: Goが調整して決めた処理の繰り返し回数（b.N）
//   - 32827 ns/op: 1回あたりの平均実行時間。小さいほど高速
//
// PASSはベンチマークが完走したという意味であり、暗号化・復号化結果の
// 正しさを検証したという意味ではない。実行時間はCPUやGoのバージョン、
// 実行時の負荷によって変わるため、RSAとAESの相対的な差に注目する。

func prepareRSA() (sourceData, label []byte, privateKey *rsa.PrivateKey) {
	sourceData = make([]byte, 128)                     // 128バイトのランダムなデータを生成
	label = []byte("")                                 // ラベルは空のバイト列
	io.ReadFull(rand.Reader, sourceData)               // ランダムなデータを生成
	privateKey, _ = rsa.GenerateKey(rand.Reader, 2048) // RSA鍵ペアを生成
	return
}

func BenchmarkRSAEncryption(b *testing.B) {
	sourceData, label, privateKey := prepareRSA() // RSA暗号化の準備
	publicKey := &privateKey.PublicKey            // 公開鍵を取得
	md5Hash := md5.New()                          // MD5ハッシュ関数を作成
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rsa.EncryptOAEP(md5Hash, rand.Reader, publicKey, sourceData, label)
	}
}

func BenchmarkRSADecryption(b *testing.B) {
	sourceData, label, privateKey := prepareRSA()                                           // RSA暗号化の準備
	publicKey := &privateKey.PublicKey                                                      // 公開鍵を取得
	md5Hash := md5.New()                                                                    // MD5ハッシュ関数を作成
	encryptedData, _ := rsa.EncryptOAEP(md5Hash, rand.Reader, publicKey, sourceData, label) // データを暗号化
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rsa.DecryptOAEP(md5Hash, rand.Reader, privateKey, encryptedData, label)
	}
}

func prepareAES() (sourceData, nonce []byte, gcm cipher.AEAD) {
	sourceData = make([]byte, 128)       // 128バイトのランダムなデータを生成
	io.ReadFull(rand.Reader, sourceData) // ランダムなデータを生成
	key := make([]byte, 32)              // 32バイトのランダムな鍵を生成
	io.ReadFull(rand.Reader, key)
	nonce = make([]byte, 12) // 12バイトのランダムなnonceを生成
	io.ReadFull(rand.Reader, nonce)
	block, _ := aes.NewCipher(key) // AESブロック暗号を作成
	gcm, _ = cipher.NewGCM(block)  // GCMモードを作成
	return
}

func BenchmarkAESEncryption(b *testing.B) {
	sourceData, nonce, gcm := prepareAES() // AES暗号化の準備
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gcm.Seal(nil, nonce, sourceData, nil) // データを暗号化
	}
}

func BenchmarkAESDecryption(b *testing.B) {
	sourceData, nonce, gcm := prepareAES()                 // AES暗号化の準備
	encryptedData := gcm.Seal(nil, nonce, sourceData, nil) // データを暗号化
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gcm.Open(nil, nonce, encryptedData, nil) // データを復号化
	}
}
