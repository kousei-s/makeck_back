package models

import (
	"errors"
)

// レシピ名を部分検索する関数
func Name_Search(name string) ([]Recipe, error) {

	// 返す配列を定義
	recipes := []Recipe{}

	// カテゴリ名のLIKE検索用のパターンを設定
	cname := "%" + name + "%" // nameを引数から取得

	// カテゴリ名からidをLIKE検索で取得
	category_len := dbconn.Where("name LIKE ?", cname).Find(&recipes).RowsAffected

	//取得件数0の時
	if category_len == 0 {
		return []Recipe{}, errors.New("Recipe not found")
	}
	
	return recipes, nil
}
