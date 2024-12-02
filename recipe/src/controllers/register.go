package controllers

import (
	"log"
	"net/http"
	"recipe/utils"

	"github.com/labstack/echo/v4"
)

type RegisterArgs struct {
	Name string `json:"name"` //料理名
}

func RegisterRecipe(ctx echo.Context) error {
	// jsonに変換
	val := new(RegisterArgs)
	if err := ctx.Bind(val); err != nil {
		//エラー処理
		utils.Println("failed to bind json : " + err.Error())
		return ctx.JSON(http.StatusBadRequest,echo.Map{
			"result" : "failed to bind json",
		})
	}

	log.Println(val.Name)

	return ctx.JSON(http.StatusOK,echo.Map{
		"result" : "success",
	})
}