package controllers

import (
	"net/http"
	"recipe/services"

	"github.com/labstack/echo/v4"
)

type ExtractParam struct {
    URL string `json:"url"`
}

func Extract(ctx echo.Context) error {
    // body 取得
    var param ExtractParam
    if err := ctx.Bind(&param); err != nil {
        // error handling
        return ctx.JSON(http.StatusBadRequest, echo.Map{
            "error": err.Error(),
        })
    }

    // 情報を抽出する
    response,err := services.Extract(param.URL)

    // エラー処理
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, echo.Map{
            "error": err.Error(),
        })
    }

    return ctx.JSON(http.StatusOK, echo.Map{
        "result": response,
    })
}