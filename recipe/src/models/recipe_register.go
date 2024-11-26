package models

import (
	"errors"
	"recipe/utils"
)

//レシピテーブルの構造体宣言
type Recipe struct {
	Uid         string   `gorm:"primaryKey"`			  //レシピID
	Name        string 									  //料理名
	Image		string 									  //画像パス
	Category  	string                                    //カテゴリー
	Process   []Process  `gorm:"foreignKey:recipeid"`	  //手順
}

//レシピを作成するための引数せんげ
type RecipeArgs struct {
	Name  		string
	Image 		string
	Category 	string
	Prosecc 	[]Process
}

// データベースにレシピを登録する処理
func Recipe_Register(args RecipeArgs) (string,error) {
	// レシピIDを生成
	uid, err := utils.Genid()
	if err != nil {
		return "",errors.New("uuid_create_error")
	}
	
	// 新しいレシピを作成
	newRecipe := Recipe{
		Uid:   uid,
		Name:  args.Name,
		Image: args.Image,
		Category: args.Category,
		Process: args.Prosecc,
	}

	result := dbconn.Create(&newRecipe)
	
	if result.Error != nil {
		return "",result.Error
	}

	return uid,err
}
