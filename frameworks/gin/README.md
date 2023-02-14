[gin-gonic/gin](https://github.com/gin-gonic/gin)

document:
- https://gin-gonic.com/ja/docs
- https://pkg.go.dev/github.com/gin-gonic/gin
examples: https://github.com/gin-gonic/examples

[ginを最速でマスターしよう - Qiita](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a)

Ginの特徴
```
高速
基数木ベースのルーティング、そして小さなメモリ使用量。reflection は使っていません。予測可能なAPIパフォーマンス。

ミドルウェアの支援
受け取った HTTP リクエストは、一連のミドルウェアと、最終的なアクションによって処理されます。 例：ログ出力、認証、GZIP 圧縮、そして最終的にDBにメッセージを投稿します。

クラッシュフリー
Gin は、HTTP リクエストの処理中に発生した panic を recover します。これにより、サーバーは常にユーザーからの応答を返すことができます。またこれにより例えば panic を Sentry に送ったりすることも可能です！

JSON バリデーション
Gin は JSON によるリクエストを解析してバリデーションすることができます。例えば必要な値の存在をチェックすることができます。

ルーティングのグループ化
ルーティングをもっとよく整理しましょう。認証が必要かどうか、異なるバージョンのAPIかどうか…加えて、パフォーマンスを低下させることなく、無制限にネストしたグループ化を行うことができます。

エラー管理
Ginは、HTTP リクエスト中に発生したすべてのエラーを収集するための便利な方法を提供します。最終的に、ミドルウェアはそれらをログファイル、データベースに書き込み、ネットワーク経由で送信することができます。

組み込みのレンダリング
Ginは、JSON、XML、およびHTMLでレンダリングするための使いやすいAPIを提供します。

拡張性
とても簡単に新しいミドルウェアを作成できます。サンプルコードをチェックしてみてください。
```


sample
```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    // rにルートを定義していく
    r.Get("/", func(c *gin.Context) {
        c.JSON(http.StatusOK,  gin.H{
            "message": "hello world",
        })
    })

    r.Run(":8080")
}
```

## Context
https://pkg.go.dev/github.com/gin-gonic/gin#Context

```go
type Context struct {
    Request *http.Request
    Writer  ResponseWriter

    Params Params

    pkgof each request.
    Keys map[string]any

    // Errors is a list of errors attached to all the handlers/middlewares who used this context.
    Errors errorMsgs

    // Accepted defines a list of manually accepted formats for content negotiation.
    Accepted []string
    // contains filtered or unexported fields
}
```
