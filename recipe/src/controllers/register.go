package controllers

import (
	"net/http"
	"recipe/middlewares"
	"recipe/services"
	"recipe/utils"
	"slices"

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

	// サービスを呼び出す
	uid, herr := services.RegisterRecipe(val)

	// エラー処理
	if herr.Err != nil {
		utils.Println("failed to register recipe : " + herr.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": uid,
	})
}

func UpdateRecipe(ctx echo.Context) error {
	// ユーザーを取得する
    user := ctx.Get("user").(middlewares.UserData)

    // admin じゃない場合
    if slices.Contains(user.Labels, "admin") == false {
        return ctx.JSON(http.StatusForbidden, map[string]interface{}{
            "error": "forbidden",
        })
    }

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

	// uid を取得
	targetUid := ctx.Request().Header.Get("recipeId")

	// サービスを呼び出す
	uid, herr := services.UpdateRecipe(targetUid,val)

	// エラー処理
	if herr.Err != nil {
		utils.Println("failed to register recipe : " + herr.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": uid,
	})
}

func RestoreRecipe(ctx echo.Context) error {
	// uid を取得
	uid := ctx.Request().Header.Get("uid")

	// サービスを呼び出す
	recipie, herr := services.RestoreRecipe(uid)

	// エラー処理
	if herr.Err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": herr.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": recipie,
	})
}

type SearchParam struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

func SearchByName(ctx echo.Context) error {
	// body 取得
	var param SearchParam
	if err := ctx.Bind(&param); err != nil {
		// error handling
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// 検索する
	recipies,err := services.SearchByName(param.Name,param.Category)

	// エラー処理
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": recipies,
	})
}

type SearchCategoryParam struct {
	Category string `json:"category"`
}
func SearchByCategory(ctx echo.Context) error {
	// body 取得
	var param SearchCategoryParam
	if err := ctx.Bind(&param); err != nil {
		// error handling
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// 検索する
	recipies,err := services.SearchByCategory(param.Category)

	// エラー処理
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": recipies,
	})
}

func UploadImage(ctx echo.Context) error {
	// body 取得
	uid := ctx.Request().Header.Get("uid")

	// 画像取得
	file, err := ctx.FormFile("image")
	if err != nil {
		// error handling
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// 検索する
	err = services.UploadImage(uid,file)

	// エラー処理
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"result": "success",
	})
}