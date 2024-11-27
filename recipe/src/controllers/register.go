package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegisterArgs struct {
	Name string `json:"name"` //料理名
}

func RegisterRecipe(ctx echo.Context) error {
	// jsonに変換
	

	return ctx.JSON(http.StatusOK,echo.Map{
		"result" : "success",
	})
}