package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	c "go-render-api/controller"
	"net/http"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/list", c.List)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}

// ハンドラーを定義
