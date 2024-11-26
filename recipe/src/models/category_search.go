package models

// カテゴリーからレシピを検索する関数
func Category_Search(id int) ([]Recipe, error) {
	// カテゴリ取得
	category,err := GetCategory(id)

	// エラー処理
	if err != nil {
		return []Recipe{},err
	}

	// 返すデータ
	recipies := []Recipe{}

	// データベースから取得
	err = dbconn.Model(&category).Association("Recipes").Find(&recipies)

	// エラー処理
	if err != nil {
		return []Recipe{},err
	}

	return recipies, nil
}
