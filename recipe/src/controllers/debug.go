package controllers

import (
	"log"
	"net/http"
	"recipe/middlewares"

	"github.com/labstack/echo/v4"
)

func DebugRecipes(ctx echo.Context) error {
    // ユーザーを取得する
    user := ctx.Get("user").(middlewares.UserData)

    log.Print(user)

    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "recipes": "",
    })
}