package controllers

import (
	"log"
	"net/http"
	"recipe/middlewares"
	"recipe/services"
	"slices"

	"github.com/labstack/echo/v4"
)

func DebugRecipes(ctx echo.Context) error {
    // ユーザーを取得する
    user := ctx.Get("user").(middlewares.UserData)

    log.Print(user)
    // admin じゃない場合
    if slices.Contains(user.Labels, "admin") == false {
        return ctx.JSON(http.StatusForbidden, map[string]interface{}{
            "error": "forbidden",
        })
    }

    // レシピを取得する
    recipies,err := services.DebugRecipes()

    // エラー処理
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, echo.Map{
            "error": err.Error(),
        })
    }

    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "recipes": recipies,
    })
}

func DebugDeleteRecipe(ctx echo.Context) error {
    // ユーザーを取得する
    user := ctx.Get("user").(middlewares.UserData)

    log.Print(user)
    // admin じゃない場合
    if slices.Contains(user.Labels, "admin") == false {
        return ctx.JSON(http.StatusForbidden, map[string]interface{}{
            "error": "forbidden",
        })
    }

    // レシピIDを取得する
    recipeid := ctx.Request().Header.Get("recipeid")

    // レシピを削除する
    err := services.DebugDeleteRecipe(recipeid)

    // エラー処理
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, echo.Map{
            "error": err.Error(),
        })
    }

    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "result": "success",
    })
}