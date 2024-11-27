package models

import (
	"recipe/utils"
)

var parallelt = true
var parallelf = false

var name1 = "キャベツ"
var name2 = "豚肉"
var name3 = "フライパン"
var name4 = "菜箸"
var name5 = "中火で炒める"
var name6 = "野菜炒め"
var name7 = "下準備"
var name8 = "全部煮る(nil)"
var image1 = "D:\\Users\\Pictures\\IMG_mattuu.pwg"

var name11 = "レタス"
var name12 = "牛肉"
var name13 = "フライパン"
var name14 = "菜箸"
var name15 = "強火で炒める"
var name16 = "野菜ファイヤー"
var name17 = "調理"
var name18 = "冷蔵庫で冷やす"
var image11 = "D:\\Users\\Pictures\\IMG_mattuu.pwg"

var name21 = "ブロッコリー"
var name22 = "鶏肉"
var name23 = "鍋"
var name24 = "菜箸"
var name25 = "中火で蒸す"
var name26 = "ブロッコリー炒め"
var name27 = "下準備"
var name28 = "全部焼く"
var image21 = "D:\\Users\\Pictures\\IMG_broccoli.pwg"

var name31 = "ニンジン"
var name32 = "魚"
var name33 = "フライパン"
var name34 = "菜箸"
var name35 = "強火で焼く"
var name36 = "ニンジンのソテー"
var name37 = "調理"
var name38 = "冷凍庫で冷やす"
var image31 = "D:\\Users\\Pictures\\IMG_carrot.pwg"

func RunDebug() {
	// ここにデバック用のコードを書く
	registration(name1, name2, name3, name4, name5, name6, name7, image1, parallelf, parallelt)
	registration(name11, name12, name13, name14, name15, name16, name17, image11, parallelf, parallelt)
	registration(name21, name22, name23, name24, name25, name26, name27, image21, parallelt, parallelf)
	registration(name31, name32, name33, name34, name35, name36, name37, image31, parallelf, parallelt)

	search()
}

// 検索
func search() {
	utils.Println("材料検索")
	material, err := Material_Search(name2)
	if err != nil {
		utils.Println(err)
	}
	utils.Println(material)

	utils.Println("名前検索")
	name, err := Name_Search("野菜")
	if err != nil {
		utils.Println(err)
	}
	utils.Println(name)

	utils.Println("カテゴリから探す")
	recipies, err := Category_Search(2)
	// エラー処理
	if err != nil {
		utils.Println(err)
	}
	utils.Println(recipies)

}

// 登録の一連の流れ
func registration(name1 string, name2 string, name3 string, name4 string, name5 string, name6 string, name7 string, image string, parallel1 bool, parallel2 bool) {

	utils.Println("材料登録１")
	material1, err := Material_Register(MaterialArgs{
		name:  name1,
		count: 1,
		unit:  "個",
	})
	if err != nil {
		utils.Println(err)
	}
	utils.Println(material1)

	utils.Println("材料登録２")
	material2, err := Material_Register(MaterialArgs{
		name:  name2,
		count: 200,
		unit:  "g",
	})
	if err != nil {
		utils.Println(err)
	}
	utils.Println(material2)

	utils.Println("器具登録１")
	tool1, err := Tools_Register(ToolArgs{
		name:  name3,
		count: 1,
		unit:  "個",
	})
	if err != nil {
		utils.Println(err)
	}
	utils.Println(tool1)

	utils.Println("器具登録２")
	tool2, err := Tools_Register(ToolArgs{
		name:  name4,
		count: 1,
		unit:  "膳",
	})
	if err != nil {
		utils.Println(err)
	}
	utils.Println(tool2)

	utils.Println("手順登録１")
	process1, err := Process_Register(ProcessArgs{
		name:        name7,
		description: name5,
		parallel:    parallel1,
		time:        10,
		tools:       []Tools{tool1, tool2},
		materials:   []Material{material1, material2},
	})
	if err != nil {
		utils.Println(err)
	}
	utils.Println(process1)

	utils.Println("手順登録２")
	process2, err := Process_Register(ProcessArgs{
		name:        name7,
		description: name5,
		parallel:    parallel2,
		time:        15,
		tools:       []Tools{tool1, tool2},
		materials:   []Material{material1, material2},
	})
	if err != nil {
		utils.Println(err)
	}
	utils.Println(process2)

	utils.Println("レシピ登録")
	recipe, err := Recipe_Register(RecipeArgs{
		Name:  name6,
		Image: image,
		//カテゴリー情報を入れる
		Category: []Category{
			{Id: 1},
			{Id: 2},
		},
		Prosecc:    []Process{process1},
		LastSatate: Hot,
	})

	if err != nil {
		utils.Println(err)
	}
	utils.Println(recipe)

}
