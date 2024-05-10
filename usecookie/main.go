package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
)

func main() {
	// 1. クッキーを保存する cookiejar (クッキー瓶) のインスタンスを作成
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	// 2. クッキー保存可能な http.Client インスタンスを作成
	client := http.Client{
		Jar: jar,
	}

	// 3. 2回アクセスする (1回目…クッキーを受信、2回目…クッキーをサーバーに対して送信)
	for i := 0; i < 2; i++ {
		// 4. http.Get() の代わりに作成したクライアントの Get() メソッドを使ってアクセス
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
