## usecookie
クッキーの送受信

### 実装
- クッキーの場合はブラウザ内部に状態を保つ必要があるため、http.Client 構造体を利用する
- [net/http/cookiejar](https://pkg.go.dev/net/http/cookiejar) は組み込みライブラリで実装されているクッキー機能
