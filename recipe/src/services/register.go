package services

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"

	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"recipe/models"
	"recipe/utils"

	"golang.org/x/image/draw"
)

type Recipe struct {
	RecipeName     string `json:"recipeName"`
	Steps          []Step `json:"steps"`
	RecipeCategory string `json:"recipeCategory"`
	FinalState     string `json:"finalState"`
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
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
	Unit     string  `json:"unit"`
}

// Utensil represents a utensil used in a step.
type Utensil struct {
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
	Unit     string  `json:"unit"`
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
	categoryId, err := ConvertToCategory(args.RecipeCategory)

	// エラー処理
	if err != nil {
		return "", HttpResult{Code: http.StatusBadRequest, Msg: err.Error(), Err: err}
	}

	// レシピIDを生成
	uid, err := utils.Genid()
	if err != nil {
		return "", HttpResult{Code: http.StatusInternalServerError, Err: err, Msg: err.Error()}
	}

	utils.Println(args)

	// 新しいレシピを作成
	_, err = models.Recipe_Register(models.RecipeArgs{
		Uid:   uid,
		Name:  args.RecipeName,
		Image: "NoImage",
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
		return "", HttpResult{Code: http.StatusInternalServerError, Err: err, Msg: err.Error()}
	}

	return uid, HttpResult{Code: http.StatusOK, Msg: "success", Err: nil}
}

// レシピを更新する
func UpdateRecipe(targetUid string,args Recipe) (string, HttpResult) {
	// レシピを削除する
	err := models.DeleteRecipe(targetUid)
	if err != nil {
		return "", HttpResult{Code: http.StatusBadRequest, Msg: err.Error(), Err: err}
	}

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
	categoryId, err := ConvertToCategory(args.RecipeCategory)

	// エラー処理
	if err != nil {
		return "", HttpResult{Code: http.StatusBadRequest, Msg: err.Error(), Err: err}
	}

	// レシピIDを生成
	uid, err := utils.Genid()
	if err != nil {
		return "", HttpResult{Code: http.StatusInternalServerError, Err: err, Msg: err.Error()}
	}

	utils.Println(args)

	// 新しいレシピを作成
	_, err = models.Recipe_Register(models.RecipeArgs{
		Uid:   uid,
		Name:  args.RecipeName,
		Image: "NoImage",
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
		return "", HttpResult{Code: http.StatusInternalServerError, Err: err, Msg: err.Error()}
	}

	return uid, HttpResult{Code: http.StatusOK, Msg: "success", Err: nil}
}

func RestoreRecipe(recipeId string) (Recipe, HttpResult) {
	// データベースからレシピを取得する
	recipe, err := models.GetRecipe(recipeId)

	// エラー処理
	if err != nil {
		return Recipe{}, HttpResult{Code: http.StatusInternalServerError, Err: err, Msg: err.Error()}
	}

	// 手順リスト
	steps := []Step{}


	// 手順を変換する
	for _, val := range recipe.Process {
		utils.Println(val)

		// 材料を取得
		gmaterials, err := val.GetMaterials()
		if err != nil {
			return Recipe{}, HttpResult{Code: http.StatusInternalServerError, Err: err, Msg: err.Error()}
		}

		// 材料を変換する
		materials := []Ingredient{}
		for _, val := range gmaterials {
			// 材料を変換
			material := Ingredient{
				Name:     val.Name,
				Quantity: val.Count,
				Unit:     val.Unit,
			}
			materials = append(materials, material)
		}

		// 器具を取得
		gtools, err := val.GetTools()
		if err != nil {
			return Recipe{}, HttpResult{Code: http.StatusInternalServerError, Err: err, Msg: err.Error()}
		}

		// 器具を変換
		tools := []Utensil{}
		for _, val := range gtools {
			// 器具を変換
			tool := Utensil{
				Name:     val.Name,
				Quantity: val.Count,
				Unit:     val.Unit,
			}
			tools = append(tools, tool)
		}

		// 手順を変換
		step := Step{
			Name:        val.Name,
			Time:        int(val.Time),
			Concurrent:  modelConcurrentToRecipeConcurrent(val.Parallel),
			Ingredients: materials,
			Utensils:    tools,
			Type:        modelTypeToRecipeType(int(val.Type)),
			Description: val.Description,
		}

		steps = append(steps, step)
	}

	// レシピを返す
	return_recipe := Recipe{
		RecipeName:     recipe.Name,
		Steps:          steps,
		RecipeCategory: recipe.Category[0].Name,
		FinalState:     modelLastStateToRecipeLastState(recipe.LastState),
	}

	return return_recipe, HttpResult{Code: http.StatusOK, Msg: "success", Err: nil}
}

func modelLastStateToRecipeLastState(lastState models.LastSatate) string {
	if lastState == models.Hot {
		return "Hot"
	} else if lastState == models.Cool {
		return "Cool"
	} else {
		return "Normal"
	}
}

func modelConcurrentToRecipeConcurrent(isConcurrent bool) string {
	if isConcurrent {
		return "可"
	} else {
		return "不可"
	}
}

func modelTypeToRecipeType(recipeType int) string {
	if recipeType == int(models.CookProcess) {
		return "調理"
	} else if recipeType == int(models.PrepareProcess) {
		return "下準備"
	} else if recipeType == int(models.FinishProcess) {
		return "仕上げ"
	} else {
		return "不明"
	}
}

func resizeImage(src image.Image, width, height int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(dst, dst.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.CatmullRom.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

// 画像を保存する関数
func saveImage(fileHeader *multipart.FileHeader, filename string) error {
	// ファイルを開く
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// ファイルの内容を読み込む
	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	// 画像をデコード
	img, _, err := image.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		return err
	}

	// 画像をリサイズする（幅800、高さはアスペクト比を維持）
	resizedImg := resizeImage(img, 400, 400)

	// リサイズした画像を保存
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// JPEG形式で保存
	if err := jpeg.Encode(outFile, resizedImg, nil); err != nil {
		return err
	}

	return nil
}

func ConvertToCategory(category string) (int, error) {
	// カテゴリ取得
	result, err := models.GetCategoryByName(category)

	// エラー処理
	if err != nil {
		return -1, err
	}

	return result.Id, nil
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

	// 手順の種類を変換する
	Type, err := ConvertTypeToInt(args.Type)
	if err != nil {
		return models.Process{}, err
	}

	return models.Process{
		Uid:         uid,
		Name:        args.Name,
		Parallel:    args.Concurrent == "可",
		Time:        args.Time,
		Tools:       tools,
		Material:    materials,
		Recipeid:    "",
		Description: args.Description,
		Type:        Type,
	}, nil
}

func ConvertTypeToInt(args string) (int, error) {
	utils.Println(args)

	switch args {
	case "調理":
		return int(models.CookProcess), nil
	case "下準備":
		return int(models.PrepareProcess), nil
	case "仕上げ":
		return int(models.FinishProcess), nil
	default:
		return -1, errors.New("invalid type")
	}
}

func UploadImage(uid string, file *multipart.FileHeader) error {
	// レシピ取得
	recipe, err := models.GetRecipe(uid)
	if err != nil {
		return err
	}

	// noimage 以外の場合
	if recipe.Image != "NoImage" {
		return errors.New("Image already exists")
	}

	recipe.Image = "https://dev-makeck.mattuu.com//recipe/images/" + uid + ".jpg"

	// レシピを更新する
	if err := models.Recipe_Update(recipe); err != nil {
		return err
	}

	// 画像を保存する
	filename := fmt.Sprintf("./images/%s.jpg", uid)
	if err := saveImage(file, filename); err != nil {
		return err
	}

	return nil
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
