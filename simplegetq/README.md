## simplegetq
GETメソッドでクエリを送信するコード。以下の curl コマンドと等価
```shell
curl -G --data-urlencode "query=hello world" http://localhost:18888
```
- `--data-urlencode` の代わりに `--data` や `-d` も使える
- `-G` は `--get` の短縮形

### 実装
- [simpleget](../simpleget/README.md) のコードに以下の変更を加える
  - クエリー文字列を `url.Values` 型を用いて宣言し、`values.Encode()` で RFC と互換性の高い変換方式を使って文字列にする

### 実行例
```go
go run simplegetq/main.go
// 2024/05/08 19:40:22 <html><body>hello</body></html>
```
```md
GET /?query=hello%2C+world HTTP/1.1
Host: localhost:18888
Accept-Encoding: gzip
User-Agent: Go-http-client/1.1
```
