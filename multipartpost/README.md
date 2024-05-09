## multipartpost
`multipart/form-data` を使ったファイル送信。HTTP/1.0 の HTML で一番複雑な送信処理。
以下の curl コマンドと等価
```shell
curl -F "name=Michael Jackson" -F "thumbnail=@multipartpost/150x150.png" http://localhost:18888
```

### 実装
- 最初に `bytes.Buffer` にマルチパートフォームの送信コンテンツを作成
- コンテンツ作成に `multipart.Writer` オブジェクトを使用
  - オブジェクトを通じて `bytes.Buffer` に内容が書き込まれる
- `multipart.Writer` は `Content-Type` に渡すバウンダリー文字列を乱数で生成する。この文字列は `Boundary()` メソッドで取得できるため、以下のコードは同等
  - `writer.FormDataContentType()`
  - `"multipart/form-data; boundary=" + writer.Boundery()`
- `multipart.Writer.WriteField()` と `multipart.Writer.CreateFormFile()` はマルチパート構成要素のパートそのものに触れずにコンテンツを生成できる高度な API。実装ではファイルの `Content-Type` フィールドを `application/octet-stream` ではなく任意の MIME タイプ (e.g. `image/png`) に設定するため、この API 内で行われている詳細処理の一部を取り出し (`CreateFormFile()` から `CreatePart()` に変更し) て指定している

### 実行例
```shell
% go run multipartpost/main.go
2024/05/09 23:07:12 Status: 200 OK
```
(任意の MIME タイプ追加前)
```shell
POST / HTTP/1.1
Host: localhost:18888
Accept-Encoding: gzip
Content-Length: 2690
Content-Type: multipart/form-data; boundary=0170fe2b901e51d6073fba7939e0557a533b7f8b7f2876a67295f4092c5d
User-Agent: Go-http-client/1.1

--0170fe2b901e51d6073fba7939e0557a533b7f8b7f2876a67295f4092c5d
Content-Disposition: form-data; name="name"

Michael Jackson
--0170fe2b901e51d6073fba7939e0557a533b7f8b7f2876a67295f4092c5d
Content-Disposition: form-data; name="thumbnail"; filename="150x150.png"
Content-Type: application/octet-stream

PNG

IHD<q   pHYsaa?IDATxiWZ$Ie

# (省略)

--0170fe2b901e51d6073fba7939e0557a533b7f8b7f2876a67295f4092c5d--
```
