package models

import "errors"

//カテゴリーからレシピを検索する関数
func Category_Search(name string)([]string,error) {
	// カテゴリを定義
	category := &Category{
        Name: name,
    } 

    // 返す配列を定義
	recipe_id := []string{}

	// カテゴリ名からidをLIKE検索で取得
	category_len := dbconn.Where(category).Find(&category).RowsAffected
	
	// エラーの時ってなにをどう返すん？
	//取得件数0の時
	if category_len == 0 {
		return []string{}, errors.New("category not found")
	}

    recipe_id = append(recipe_id,category.Name )

	// // 取得したカテゴリからIDを抽出
	// for _, id := range category.Name {
	// 	recipe_id = append(recipe_id) // IDの型に応じて適切に変更
	// }

	return recipe_id, nil
}