## simplehead
HEADメソッドを送信するコード。以下の curl コマンドと等価
```shell
curl --head http://localhost:18888
```

### 実装
- [simpleget](../simpleget/README.md) の `http.Get()` を `http.Head()` にするだけ

### 実行例
```shell
% go run simplehead/main.go
2024/05/08 22:06:03
2024/05/08 22:06:03 Status: 200 OK
2024/05/08 22:06:03 StatusCode: 200
2024/05/08 22:06:03 Fields: map[Content-Length:[32] Content-Type:[text/html; charset=utf-8] Date:[Wed, 08 May 2024 13:06:03 GMT]]
2024/05/08 22:06:03 Content-Length: 32
```
```shell
HEAD / HTTP/1.1
Host: localhost:18888
User-Agent: Go-http-client/1.1
```
