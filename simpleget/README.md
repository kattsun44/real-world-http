## simpleget
一番簡単なGETメソッドを送信するコード。`curl http://localhost:18888` と等価

> [!tip]
> defer はリソース生成と同じスコープ内で近くに終了処理を書くテクニック。
> 言語を問わず利用され、Java の try-with-resources, Ruby の File.open, Python の with 構文、C# や JavaScript の using などが該当する。

### 実装
`resp` 変数に入っている要素が `http.Response` 型のオブジェクト。サーバーから返ってきた情報をすべて格納している。

- `resp.Status`, `resp.StatusCode`: ステータスが格納されている (e.g. `200`, `"200 OK"`)
- `resp.Header` ヘッダーフィールドの一覧が格納されている

`http.Get` は高水準な API。未指定で10回までリダイレクトを追いかけるなど、ブラウザの挙動を意識せずに使用できる

### 実行例
```shell
go run simpleget/main.go
```
```md
2024/05/08 19:40:22 <html><body>hello</body></html>

2024/05/08 19:40:22 Status: 200 OK
2024/05/08 19:40:22 StatusCode: 200
2024/05/08 19:40:22 Fields: map[Content-Length:[32] Content-Type:[text/html; charset=utf-8] Date:[Wed, 08 May 2024 10:40:22 GMT]]
2024/05/08 19:40:22 Content-Length: 32
```
```md
GET / HTTP/1.1
Host: localhost:18888
Accept-Encoding: gzip
User-Agent: Go-http-client/1.1
```
