package models

import (
	"log"
	"recipe/utils"
	"time"
)

type Category struct {
	Id      int       `gorm:"primaryKey"` //カテゴリーID
	Name    string    //カテゴリー名
	Recipes []*Recipe `gorm:"many2many:recipe_category"`
}

type RecipeCategory struct {
	CategoryID int    `gorm:"primaryKey"`
	RecipeID   string `gorm:"primaryKey"`
	CreatedAt  time.Time
}

func initcategory() {
	defer func() {
		if rec := recover(); rec != nil {
			log.Printf("Recovered from: %w", rec)
		}
	}()

	//以下の要素を作成
	categories := []Category{
		{Id: 1, Name: "主食"},
		{Id: 2, Name: "主菜"},
		{Id: 3, Name: "副菜"},
		{Id: 4, Name: "汁物"},
	}

	for index := range categories {
		//カテゴリー作成
		result := dbconn.Create(&categories[index])

		if result.Error != nil {
			utils.Println(result)
		}
	}

}

func GetCategory(id int) (Category, error) {
	// データを格納する変数
	data := Category{}

	// 取得
	result := dbconn.Where(&Category{
		Id: id,
	}).First(&data)

	return data, result.Error
}


func GetCategoryByName(name string) (Category, error) {
	// データを格納する変数
	data := Category{}

	// 取得
	result := dbconn.Where(&Category{
		Name: name,
	}).First(&data)

	return data, result.Error
}
