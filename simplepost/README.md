## simplepost
`x-www-form-urlencoded` 形式でPOSTフォーム送信するコード。以下の curl コマンドと等価
```shell
curl --data-urlencode "query=hello world" http://localhost:18888
```

### 実装
- [simplegetq](../simplegetq/README.md) で出てきた `url.Values` 型の変数を `http.PostForm()` に渡す
  - Go は RFC 3986 に準拠した変換を行う (と書かれているが、スペースが `%20` ではなく `+` に変換されるのはなぜ？)

### 実行例
```shell
% go run simplepost/main.go
2024/05/09 07:33:44 Status: 200 OK
```
```shell
POST / HTTP/1.1
Host: localhost:18888
Accept-Encoding: gzip
Content-Length: 10
Content-Type: application/x-www-form-urlencoded
User-Agent: Go-http-client/1.1

query=hello%2C+world
```
