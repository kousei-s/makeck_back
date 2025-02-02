package services

import "log"

type ReturnMaterial struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

type ReturnRecipe struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Materials []ReturnMaterial `json:"materials"`
}

type MaterialArgs struct {
	Recipes []string
}

func GetMaterial(args RecipeArgs) ([]ReturnRecipe, error) {
	convertedRecipes, err := converter(args)
	if err != nil {
		return []ReturnRecipe{}, err
	}

	// 返すデータ
	returnData := []ReturnRecipe{}

	// レシピを回す
	for _, recipe := range convertedRecipes {
		// 材料を格納するマップ
		materials := map[string]*ReturnMaterial{}

		// 手順を回す
		for _, process := range recipe.Process {
			log.Println(process.Material)

			// 手順の材料を回す
			for _, material := range process.Material {
				if _, ok := materials[material.Name]; ok {
					// 材料の名前が存在するとき
					// 取得する
					materials[material.Name].Quantity += material.Count
				} else {
					// 材料の名前が存在しないとき
					materials[material.Name] = &ReturnMaterial{
						Name:     material.Name,
						Quantity: material.Count,
						Unit:     material.Unit,
					}
				}
			}
		}

		// マップを配列に変換する
		materialsList := []ReturnMaterial{}
		for _, material := range materials {
			materialsList = append(materialsList, *material)
		}

		// 返すデータに追加する
		returnData = append(returnData, ReturnRecipe{
			ID:        recipe.Uid,
			Name:      recipe.Name,
			Materials: materialsList,
		})
	}

	return returnData, nil
}

func SerMaterials(args RecipeArgs) ([]Recipes_list, error) {
	convertedRecipes, err := converter(args)
	if err != nil {
		return []Recipes_list{}, err
	}

	return SearchMaterials(convertedRecipes)
}
