package services

import (
	"recipe/models"
)

type MatchRecipies struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Image string `json:"image"`
}

func SearchByName(name string,category string) ([]MatchRecipies,error) {
	// カテゴリ取得する
	targetCategory,err := models.GetCategoryByName(category)

	// エラー処理
	if err != nil {
		return []MatchRecipies{},err
	}

	
	result := []MatchRecipies{}
	if name == "" {
		// レシピを検索する
		recipies,err := models.Category_Search(targetCategory.Id)

		// エラー処理
		if err != nil {
			return []MatchRecipies{},err
		}

		for _, recipie := range recipies {
			// 追加
			result = append(result, MatchRecipies{
				ID:    recipie.Uid,
				Name:  recipie.Name,
				Image: recipie.Image,
			})
		}

		return result,nil
	}

	// レシピを検索する
	recipies,err := models.Name_Search(name)

	// エラー処理
	if err != nil {
		return []MatchRecipies{},err
	}

	for _, recipie := range recipies {
		// カテゴリで弾く
		if !recipie.CheckCategory(targetCategory.Id) {
			continue
		}

		// 追加
		result = append(result, MatchRecipies{
			ID:    recipie.Uid,
			Name:  recipie.Name,
			Image: recipie.Image,
		})
	}

	return result,nil

	// {
	// 	"id": "ce9c3514d8434f92b0675562466b0284",
	// 	"name": "田舎風トリのから揚げ",
	// 	"image": "https://makeck.mattuu.com/images/ce9c3514d8434f92b0675562466b0284.jpg"
	// },
}
