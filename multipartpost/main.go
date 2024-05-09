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
	// 1. マルチパート部を組み立てた後のバイト列を格納するバッファ
	var buffer bytes.Buffer

	// 2. マルチパート部を組み立てるライター
	writer := multipart.NewWriter(&buffer)

	// 3. ファイル以外のフォームフィールドの値は WriteField メソッドで登録
	writer.WriteField("name", "Michael Jackson")

	// ** 4 から 6 までファイルを読み込む操作 ** //
	// 4. 個別のファイル書き込みの io.Writer を作成
	fileWriter, err := writer.CreateFormFile("thumbnail", "150x150.png")
	if err != nil {
		panic(err)
	}

	// 5. ファイルを開く
	readFile, err := os.Open("multipartpost/150x150.png")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	// 6. ファイルの全コンテンツをファイル書き込み用の io.Writer にコピー
	io.Copy(fileWriter, readFile)

	// 7. マルチパートの io.Writer をクローズし、バッファにすべてを書き込む
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
