package models

import (
	"log"
)

type Category struct {
	Id       int    `gorm:"primaryKey"` //カテゴリーID
	Name     string 					//カテゴリー名
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
			log.Println(result)
		}
	}

}
