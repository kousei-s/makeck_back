package main

import (
	// "recipe/controllers"
	"net/http"
	"recipe/controllers"
	"recipe/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer() *echo.Echo {
	// サーバー作成
	server := echo.New()

	// ミドルウェア
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	server.POST("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!")
	},middlewares.PocketAuth())

	// レシピ名を登録するエンドポイント
	server.POST("/register_recipe",controllers.RegisterRecipe,middlewares.PocketAuth())
	// 画像をアップロードするエンドポイント
	server.POST("/upload_image",controllers.UploadImage,middlewares.PocketAuth())

	// 画像ディレクトリ公開
	server.Static("/images", "./images")
	
	// レシプピを検索するエンドポイント
	server.POST("/search",controllers.SearchByName)

	// カテゴリからレシピを検索するエンドポイント
	server.POST("/search_category",controllers.SearchByCategory)

	// データを抽出するエンドポイント
	server.POST("/extract",controllers.Extract,middlewares.PocketAuth())

	// 手順の詳細を取得するエンドポイント
	server.GET("/:id",controllers.GetProcess)
	
	return server
}