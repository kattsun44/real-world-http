## anyfilepost
POST メソッドを使って任意のコンテンツを送信するコード。以下の curl コマンドと等価
```shell
curl --data-binary @anyfilepost/main.go -H "Content-Type: text/plain" http://localhost:18888
```

### 実装
- Go の `http.Post()` を使う
  - `Content-Type` ヘッダーフィールドの内容は第2引数に指定する
  - 送信内容は `io.Reader` の形式で第3引数に指定する (`os.File` は `io.Reader` インタフェースを満たす)
    - ファイルではなく文字列を `http.Post` にわたす場合は、`bytes.Buffer` や `strings.Reader` を使って文字列を `io.Reader` インタフェース化する

### 実行例
```shell
% go run anyfilepost/main.go
2024/05/09 22:12:13 Status: 200 OK
```
```shell
POST / HTTP/1.1
Host: localhost:18888
Transfer-Encoding: chunked
Accept-Encoding: gzip
Content-Type: text/plain
User-Agent: Go-http-client/1.1

11d
package main

import (
        "log"
        "net/http"
        "os"
)

func main() {
        file, err := os.Open("anyfilepost/main.go")
        if err != nil {
                panic(err)
        }
        resp, err := http.Post("http://localhost:18888", "text/plain", file)
        if err != nil {
                panic(err)
        }
        log.Println("Status:", resp.Status)
}

0

```
(`Content-Length` が表示されないのと、途中に `11d`、最後に `0` が表示されるのはなんでなんだろう…。curl だとこんな表示はされない)
