package controllers

import (
	"chart/services"
	"chart/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Chart struct {
	RecipeIds []string `json:"recipe_ids"`
}

func GenChart(ctx echo.Context) error {
	// bind
	var chart Chart
	if err := ctx.Bind(&chart); err != nil {
		return err
	}

	// 長さが４以外の時
	if len(chart.RecipeIds) != 4 {
		utils.Println("長さが４以外")
		return ctx.String(http.StatusBadRequest, "Bad Request")
	}

	// 重複チェック
	for i, id := range chart.RecipeIds {
		for j := i + 1; j < len(chart.RecipeIds); j++ {
			if id == chart.RecipeIds[j] {
				utils.Println("重複あり")
				return ctx.String(http.StatusBadRequest, "Bad Request")
			}
		}
	}

	// ガントチャート作成
	chartData,err := services.GenChart(services.RecipeArgs{
		Recipes: chart.RecipeIds,
	})

	if err != nil {
		utils.Println(err)
		return ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.JSON(http.StatusOK, chartData)
}