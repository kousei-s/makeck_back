package services

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	_ "image/jpeg"
	_ "image/gif"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"recipe/models"
	"recipe/utils"

	"golang.org/x/image/draw"
)

type Recipe struct {
	RecipeName     string                `json:"recipeName"`
	Steps          []Step                `json:"steps"`
	RecipeCategory string                `json:"recipeCategory"`
	FinalState     string                `json:"finalState"`
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
	newWidth := 800
	newHeight := int(float64(newWidth) * float64(img.Bounds().Dy()) / float64(img.Bounds().Dx()))
	resizedImg := resizeImage(img, newWidth, newHeight)

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

	return models.Process{
		Uid:         uid,
		Name:        args.Name,
		Parallel:    args.Concurrent == "可",
		Time:        args.Time,
		Tools:       tools,
		Material:    materials,
		Recipeid:    "",
		Description: args.Description,
	}, nil
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

	recipe.Image = "https://makeck.tail6cf7b.ts.net:8030/recipe/images/" + uid + ".jpg"

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
