package services

import (
	"encoding/json"
	"recipe/utils"
)

const (
	SampleJson = `{"recipeCategory":"副菜","recipeName":"かぼちゃのリゾット","recipeImage":"","finalState":"Reheat","steps":[{"name":"test1","time":5,"type":"下準備","concurrent":"可","ingredients":[{"name":"かぼちゃ","quantity":1,"unit":"個"}],"utensils":[{"name":"包丁","quantity":1,"unit":"本"}],"description":"test1"},{"name":"test2","time":5,"type":"下準備","concurrent":"可","ingredients":[{"name":"水","quantity":1,"unit":"L"}],"utensils":[{"name":"包丁","quantity":1,"unit":"丁"}],"description":"test2"}]}`
)

func Test() {
	utils.Println("test Register Recipe")

	// json を受け取る
	var val Recipe
	err := json.Unmarshal([]byte(SampleJson), &val)

	// エラー処理
	if err != nil {
		utils.Println(err)
		return
	}

	// サービスを呼び出す
	_, herr := RegisterRecipe(val)

	// エラー処理
	if herr.Err != nil {
		utils.Println("failed to register recipe : " + herr.Error())
		return
	}
}