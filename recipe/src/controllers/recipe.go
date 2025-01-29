package controllers

import (
	"net/http"
	"recipe/services"

	"github.com/labstack/echo/v4"
)

func GetProcess(ctx echo.Context) error {
	processid := ctx.Param("id")

	// エラー処理
	if processid == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "processid is empty",
		})
	}

	// サービスを呼び出す
	process, err := services.GetProcess(processid)

	// エラー処理
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, process)
}