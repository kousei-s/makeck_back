package services

import (
	"encoding/json"
	"fmt"
)

// 材料を表す構造体
type SearchMaterials struct {
	MaterialName string `json:"material_name"`
	Quantity     int    `json:"quantity"`
	Unit         string `json:"unit"`
}

// レシピを表す構造体
type SearchRecipes struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Materials []SearchMaterials `json:"materials"`
}

func main() {
	// サンプルデータ
	recipes := []SearchRecipes{
		{
			ID:   "f1ac3c05-d055-4695-8083-ec2434dbafa1",
			Name: "トマトスパゲッティ",
			Materials: []SearchMaterials{
				{"スパゲッティ", 100, "g"},
				{"トマト", 2, "個"},
				{"オリーブオイル", 2, "大さじ"},
				{"塩", 1, "小さじ"},
				{"バジル", 5, "枚"},
			},
		},
		{
			ID:   "f36c60e5-1345-4bc1-9359-7daee17121d5",
			Name: "スパゲッティ",
			Materials: []SearchMaterials{
				{"スパゲッティ", 100, "g"},
				{"トマト", 2, "個"},
				{"オリーブオイル", 2, "大さじ"},
				{"塩", 1, "小さじ"},
				{"バジル", 5, "枚"},
			},
		},
	}

	// JSONにエンコードして表示
	jsonData, err := json.MarshalIndent(recipes, "", "  ")
	if err != nil {
		fmt.Println("エンコードエラー:", err)
		return
	}
	fmt.Println(string(jsonData))
}
