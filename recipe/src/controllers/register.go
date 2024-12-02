package controllers

import (
	"net/http"
	"recipe/services"
	"recipe/utils"

	"github.com/labstack/echo/v4"
)


func RegisterRecipe(ctx echo.Context) error {
	// json を受け取る
	var val services.Recipe
	err := ctx.Bind(&val)

	// エラー処理
	if err != nil {
		utils.Println(err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"result": "failed",
		})
	}

	// バリデーション
	

	// サービスを呼び出す
	_, herr := services.RegisterRecipe(val)

	// エラー処理
	if herr.Err != nil {
		utils.Println("failed to register recipe : " + herr.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{
		
	})}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "success",
	})
}
