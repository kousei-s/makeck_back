package services

import (
	"recipe/models"
	"strings"
)

type ReturnRecipe struct {
	Uid         string	`json:"id"`
	Name        string	`json:"name"`
	Category    string	`json:"type"`
	Image       string	`json:"image"`
	LastState   string	`json:"status"`
}

func DebugRecipes() ([]ReturnRecipe, error) {
	returnRecipe := []ReturnRecipe{}

	//	全レシピ取得
	getRecipes, err := models.GetAllRecipes()	

	// エラー処理
	if err != nil {
		return returnRecipe, err
	}

	for _, val := range getRecipes {
		returnRecipe = append(returnRecipe, ReturnRecipe{
			Uid:         val.Uid,
			Name:        val.Name,
			Category:    recipeCategorysToString(val.Category),
			Image:       strings.Replace(val.Image,"makeck.tail6cf7b.ts.net:8030","dev-makeck.mattuu.com",1),
			LastState:   string(val.LastState),
		})
	}

	return returnRecipe, nil
}

func recipeCategorysToString(categorys []*models.Category) string {
	var result string
	for _, val := range categorys {
		result += val.Name + " "
	}
	return result
}

func DebugDeleteRecipe(recipeid string) error {
	return models.DeleteRecipe(recipeid)
}