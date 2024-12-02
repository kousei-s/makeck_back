package services

import (
	"net/http"
	"recipe/models"
)


type Recipe struct {
	RecipeName   string   `json:"recipeName"`
	RecipeImage  string   `json:"recipeImage"`
	Steps        []Step   `json:"steps"`
}

// Step represents each step in the recipe.
type Step struct {
	Name        string      `json:"name"`
	Time        string      `json:"time"`
	Concurrent  string      `json:"concurrent"`
	Ingredients []Ingredient `json:"ingredients"`
	Utensils    []Utensil    `json:"utensils"`
	Description string      `json:"description"`
}

// Ingredient represents an ingredient in a step.
type Ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Unit     string `json:"unit"`
}

// Utensil represents a utensil used in a step.
type Utensil struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Unit     string `json:"unit"`
}

// レシピのID とエラーを返す
func RegisterRecipe(args Recipe) (string,HttpResult) {
	// 新しいレシピを作成
	models.Recipe_Register(models.RecipeArgs{
		Name:     args.Name,
		Image:    "/images/recipe.png",
		Category: []models.Category{
			{
				Id: 1,
			},
		},
		Prosecc:  []models.Process{},
		LastSatate: models.Normal,
	})

	return "", HttpResult{Code: http.StatusOK, Msg: "success", Err: nil}
}