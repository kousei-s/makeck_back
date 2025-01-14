package models

import "gorm.io/gorm/clause"

type LastSatate string

const (
	Hot    = LastSatate("hot")
	Cool   = LastSatate("cool")
	Normal = LastSatate("normal")
	reheat = LastSatate("reheat")
)

// レシピテーブルの構造体宣言
type Recipe struct {
	Uid       string      `gorm:"primaryKey"` //レシピID
	Name      string      //料理名
	Image     string      //画像パス
	Category  []*Category `gorm:"many2many:recipe_category;foreignKey:uid"` //カテゴリー
	Process   []Process   `gorm:"foreignKey:recipeid"`                      //手順
	LastState LastSatate  //最終状態
}

// レシピを作成するための引数宣言
type RecipeArgs struct {
	Uid        string
	Name       string
	Image      string
	Category   []Category
	Prosecc    []Process
	LastSatate LastSatate
}

// データベースにレシピを登録する処理
func Recipe_Register(args RecipeArgs) (string, error) {
	// 新しいレシピを作成
	newRecipe := Recipe{
		Uid:       args.Uid,
		Name:      args.Name,
		Image:     args.Image,
		Process:   args.Prosecc,
		LastState: args.LastSatate,
	}

	result := dbconn.Save(&newRecipe)

	if result.Error != nil {
		return "", result.Error
	}

	// 追加予定のリスト
	append_list := []Category{}

	// カテゴリーを検証
	for _, val := range args.Category {
		// カテゴリーを検証
		category, err := GetCategory(val.Id)

		// エラー処理
		if err != nil {
			continue
		}

		// 存在するとき追加する
		append_list = append(append_list, category)
	}

	// カテゴリーを追加する処理
	err := dbconn.Model(&newRecipe).Association("Category").Append(append_list)

	return args.Uid, err
}

func (recipe *Recipe) CheckCategory(categoryid int) bool {
	Categories := []Category{}

	// データベースから取得
	err := dbconn.Model(recipe).Association("Category").Find(&Categories)

	// エラー処理
	if err != nil {
		return false
	}

	for _, category := range Categories {
		if category.Id == categoryid {
			return true
		}
	}

	return false
}

func GetRecipe(uid string) (Recipe, error) {
	// データを格納する変数
	data := Recipe{}

	// 取得
	result := dbconn.Where(&Recipe{
		Uid: uid,
	}).Preload(clause.Associations).First(&data)

	return data, result.Error
}

func Recipe_Update(recipe Recipe) error {
	return dbconn.Save(&recipe).Error
}