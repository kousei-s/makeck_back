package models

import (
	"errors"
)

//レシピ名を部分検索する関数
func Name_Search(name string) ([]string, error) {
	// カテゴリを定義
	recipe := []Recipe{} 
	// 返す配列を定義
	recipe_id := []string{}

	// カテゴリ名のLIKE検索用のパターンを設定
	cname := "%" + name + "%" // nameを引数から取得

	// カテゴリ名からidをLIKE検索で取得
	category_len := dbconn.Where("name LIKE ?", cname).Find(&recipe).RowsAffected
	
	// エラーの時ってなにをどう返すん？
	//取得件数0の時
	if category_len == 0 {
		return recipe_id, errors.New("category not found")
	}

	// 取得したカテゴリからIDを抽出
	for _, category := range recipe {
		recipe_id = append(recipe_id, category.Uid) // IDの型に応じて適切に変更
	}

	return recipe_id, nil
}
