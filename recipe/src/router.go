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
	// server.Use(middleware.Recover())

	server.POST("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!")
	},middlewares.PocketAuth())

	// レシピ名を登録するエンドポイント
	server.POST("/register_recipe",controllers.RegisterRecipe)
	
	// レシプピを検索するエンドポイント
	server.POST("/search",controllers.SearchByName)
	
	return server
}