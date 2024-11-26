package models

import (
	"log"
)

// var name1 = "キャベツ"
// var name2 = "豚肉"
// var name3 = "フライパン"
// var name4 = "菜箸"
// var name5 = "中火で炒める"
// var name6 = "野菜炒め"

var name1 = "レタス"
var name2 = "牛肉"
var name3 = "フライパン"
var name4 = "菜箸"
var name5 = "強火で炒める"
var name6 = "野菜ファイヤー"

var image = "C:\\Users\\2230010\\Pictures\\IMG_6091.png"


func RunDebug() {
	// ここにデバック用のコードを書く
	registration()

	search()
}


//検索
func search() {
	log.Print("材料検索")
	material,err := Material_Search(name2)
	if err != nil {log.Print(err)}
	log.Print(material)

	log.Print("名前検索")
	name,err := Name_Search("野菜")
	if err != nil {log.Print(err)}
	log.Print(name)
}

//登録の一連の流れ
func registration() {

	log.Print("材料登録１")
	material1,err := Material_Register(MaterialArgs{
		name:       name1,
		count:      1,
		unit:       "個",
	})
	if err != nil {log.Print(err)}
	log.Print(material1)

	log.Print("材料登録２")
	material2,err := Material_Register(MaterialArgs{
		name:       name2,
		count:      200,
		unit:       "g",
	})
	if err != nil {log.Print(err)}
	log.Print(material2)


	log.Print("器具登録１")
	tool1,err := Tools_Register(ToolArgs{
		name:  name3,
		count: 1,
		unit:  "個",
	})
	if err != nil {log.Print(err)}
	log.Print(tool1)

	log.Print("器具登録２")
	tool2,err := Tools_Register(ToolArgs{
		name:  name4,
		count: 1,
		unit:  "膳",
	})
	if err != nil {log.Print(err)}
	log.Print(tool2)


	log.Print("手順登録１")
	process1,err := Process_Register(ProcessArgs{
		name:      name5,
		parallel:  false,
		time:      10,
		tools: []Tools{tool1,tool2},
		materials: []Material{material1,material2},
	})

	if err != nil {log.Print(err)}
	log.Print(process1)


	log.Print("レシピ登録")
	recipe,err := Recipe_Register(RecipeArgs{
		Name:     name6,
		Image:    image,
		Category: "1,2",
		Prosecc:  []Process{process1},
	})

	if err != nil {log.Print(err)}
	log.Print(recipe)
}

