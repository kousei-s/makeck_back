package main

import (
	// "recipe/controllers"
	"recipe/middlewares"
	"net/http"

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
	
	return server
}