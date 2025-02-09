package services

// 材料を表す構造体
type Materials_list struct {
	Name string `json:"material_name"`
	Quantity     float32    `json:"quantity"`
	Unit         string `json:"unit"`
}

// レシピを表す構造体
type Recipes_list struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Materials []Materials_list `json:"materials"`
}


// SearchMaterials はレシピ情報を検索します
func SearchMaterials(recipes []Recipe) ([]Recipes_list, error) {
	var result []Recipes_list

	materialslists := []Materials_list{}

	// 各レシピをループして、新しい構造体に変換
	for _, recipe := range recipes {
		for _, material := range recipe.Process[0].Material {
			materialsList := Materials_list{
				Name: material.Name,
				Quantity:     material.Count,
				Unit:         material.Unit,
			}
			materialslists = append(materialslists, materialsList)
		}
		recipeslist := Recipes_list{
			ID:        recipe.Uid,
			Name:      recipe.Name,
			Materials: materialslists,
		}
		result = append(result, recipeslist)
	}

	return result,nil
}
