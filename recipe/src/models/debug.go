package models

import (
	"log"
)

func RunDebug() {
	// ここにデバック用のコードを書く
	name1 := "温野菜のサラダ"
	// name2 := "野菜のソテー"
	// cname := "主菜"
	image := "C:\\Users\\2230010\\Pictures\\IMG_6091.png"

	log.Print("レシピ登録")
	id1, err := Recipe_Register(RecipeArgs{
		Name:     name1,
		Image:    image,
		Category: "1,2,3,4",
		Prosecc:  []Process{},
	})
	if err != nil {
		log.Print(err)		
	}
	log.Print(id1)




}