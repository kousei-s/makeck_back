package services

import (
	"net/http"
	"recipe/models"
	"recipe/utils"
)

type Recipe struct {
	RecipeName  string `json:"recipeName"`
	RecipeImage string `json:"recipeImage"`
	Steps       []Step `json:"steps"`
	RecipeCategory string `json:"recipeCategory"`
	FinalState	string `json:"finalState"`
}

// Step represents each step in the recipe.
type Step struct {
	Name        string       `json:"name"`
	Time        int          `json:"time"`
	Concurrent  string       `json:"concurrent"`
	Ingredients []Ingredient `json:"ingredients"`
	Utensils    []Utensil    `json:"utensils"`
	Type        string       `json:"type"`
	Description string       `json:"description"`
}

// Ingredient represents an ingredient in a step.
type Ingredient struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

// Utensil represents a utensil used in a step.
type Utensil struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

// レシピのID とエラーを返す
func RegisterRecipe(args Recipe) (string, HttpResult) {
	procescc := []models.Process{}

	// step を手順に変換する
	for _, val := range args.Steps {
		// レシピを登録する
		process, err := ConvertToProcess(val)
		if err != nil {
			return "", HttpResult{Code: http.StatusBadRequest, Msg: err.Error(), Err: err}
		}

		procescc = append(procescc, process)
	}

	// カテゴリを変換する
	categoryId,err := ConvertToCategory(args.RecipeCategory)

	// エラー処理
	if err != nil {
		return "",HttpResult{Code: http.StatusBadRequest,Msg: err.Error(),Err: err}
	}

	// 新しいレシピを作成
	_,err = models.Recipe_Register(models.RecipeArgs{
		Name:  args.RecipeName,
		Image: "/images/recipe.png",
		Category: []models.Category{
			{
				Id: categoryId,
			},
		},
		Prosecc:    procescc,
		LastSatate: models.LastSatate(args.FinalState),
	})

	// エラー処理
	if err != nil {
		return "",HttpResult{Code: http.StatusInternalServerError,Err: err,Msg: err.Error()}
	}

	return "", HttpResult{Code: http.StatusOK, Msg: "success", Err: nil}
}

func ConvertToCategory(category string) (int,error) {
	// カテゴリ取得
	result,err := models.GetCategoryByName(category)

	// エラー処理
	if err != nil {
		return -1,err
	}

	return result.Id,nil
}

func ConvertToProcess(args Step) (models.Process, error) {
	uid, err := utils.Genid()

	// エラー処理
	if err != nil {
		return models.Process{}, err
	}

	// 器具を変換する
	tools, err := ConvertAllToTools(args.Utensils)
	if err != nil {
		return models.Process{}, err
	}

	// 材料を変換する
	materials, err := ConvertAllToMaterials(args.Ingredients)
	if err != nil {
		return models.Process{}, err
	}

	return models.Process{
		Uid:      uid,
		Name:     args.Name,
		Parallel: args.Concurrent == "可",
		Time:     args.Time,
		Tools:    tools,
		Material: materials,
		Recipeid: "",
		Description: args.Description,
	}, nil
}

func ConvertAllToTools(args []Utensil) ([]models.Tools, error) {
	var tools []models.Tools
	for _, val := range args {
		tool, err := ConvertToTool(val)
		if err != nil {
			return nil, err
		}
		tools = append(tools, tool)
	}

	return tools, nil
}

func ConvertToTool(args Utensil) (models.Tools, error) {
	uid, err := utils.Genid()

	// エラー処理
	if err != nil {
		return models.Tools{}, err
	}

	return models.Tools{
		Uid:   uid,
		Name:  args.Name,
		Count: args.Quantity,
		Unit:  args.Unit,
	}, nil
}

func ConvertAllToMaterials(args []Ingredient) ([]models.Material, error) {
	var materials []models.Material
	for _, val := range args {
		material, err := ConvertToMaterial(val)
		if err != nil {
			return nil, err
		}
		materials = append(materials, material)
	}
	return materials, nil
}

func ConvertToMaterial(args Ingredient) (models.Material, error) {
	uid, err := utils.Genid()

	// エラー処理
	if err != nil {
		return models.Material{}, err
	}

	return models.Material{
		Uid:   uid,
		Name:  args.Name,
		Count: args.Quantity,
		Unit:  args.Unit,
	}, nil
}
