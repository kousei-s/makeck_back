package models

import (
	"errors"
)

// 材料からレシピを検索する関数
func Material_Search(name string) ([]Recipe, error) {
	// 材料を定義
	material := &Material{
		Name: name,
	}

	//材料構造体定義
	materials := []Material{}

	//手順構造体定義
	processes := []Process{}

	//レシピ構造体定義
	recipes := []Recipe{}

	//材料から取得
	material_len := dbconn.Where(material).Find(&materials).RowsAffected

	//材料が見つからなかった場合
	if material_len == 0 {
		return []Recipe{}, errors.New("Material not found")
	}

	for index := range materials {

		//手順から取得
		process := &Process{
			Uid: materials[index].Processid,
		}

		process_len := dbconn.Where(process).Find(&processes).RowsAffected

		//手順が見つからなかった場合
		if process_len == 0 {
			return []Recipe{}, errors.New("Process not found")
		}

		recipe := &Recipe{
			Uid: process.Recipeid,
		}


		//レシピから取得
		recipe_len := dbconn.Where(recipe).Find(&recipes).RowsAffected

		if recipe_len == 0 {
			return []Recipe{}, errors.New("Process not found")
		}
	}

	return recipes, nil
}
