package services

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // Recipe はレシピ情報を表します
// type Recipes struct {
// 	ID   string `json:"id"`   // レシピID
// 	Name string `json:"name"` // レシピ名
// }

// // Task はタスク情報を表します
// type Task struct {
// 	Name      string           `json:"name"`      // タスク名
// 	Tejuns    map[string]Tejun `json:"tejuns"`    // 手順情報
// 	StartTime int              `json:"startTime"` // 開始時間
// }

// // Tejun は手順情報を表します
// type Tejun struct {
// 	ID   string `json:"id,omitempty"`   // 手順ID (オプション)
// 	Name string `json:"name,omitempty"` // 手順名 (オプション)
// 	Time int    `json:"time,omitempty"` // 所要時間 (オプション)
// }

// // RecipeCollection はレシピとタスクのコレクションを表します
// type RecipeCollection struct {
// 	Recipes []ShortRecipe `json:"recipies"` // レシピのリスト
// 	Tasks   []Task        `json:"tasks"`    // タスクのリスト
// }

// // ShortRecipe は簡略化されたレシピ情報を表します
// type ShortRecipe struct {
// 	Uid  string // レシピのユニークID
// 	Name string // レシピ名
// }

// func chart_Register(recipes []Recipe) error {
// 	shortRecipes := make([]ShortRecipe, len(recipes))
// 	for i, recipe := range recipes {
// 		shortRecipes[i] = ShortRecipe{
// 			Uid:  recipe.Uid,
// 			Name: recipe.Name,
// 		}
// 	}

// 	tasks := []Task{}

// 	for _, recipe := range recipes {
// 		tejuns := make(map[string]Tejun) // 手順のマップを初期化

// 		// 手順がない場合でも空の構造体を追加するためのフラグ
// 		hasTejun := false

// 		for _, process := range recipe.Process {
// 		fmt.Print(process.Parallel)
// 			if process.Time > 0 {
// 				// 手順の情報を設定
// 				tejuns[process.Uid] = Tejun{
// 					// ID:   process.Uid,
// 					// Name: process.Name,
// 					// Time: process.Time,
// 				}
// 				hasTejun = true // 手順が存在することを示す
// 			} else {
// 				// 手順がない場合は空の構造体を追加
// 				tejuns[recipe.Uid] = Tejun{}
// 			}

// 			// 手順がない場合でも空の構造体を追加
// 			if !hasTejun {
// 				tejuns[recipe.Uid] = Tejun{} // 空の手順を追加
// 			}
// 		}

// 		// タスクを追加
// 		tasks = append(tasks, Task{
// 			Name:     "task1", // タスク名はレシピのUidを使用
// 			Tejuns:    tejuns,
// 			StartTime: 0, // 一旦開始時間を0に設定
// 		})
// 	}

// 	// レシピコレクションを作成
// 	recipeCollection := RecipeCollection{
// 		Recipes: shortRecipes,
// 		Tasks:   tasks,
// 	}

// 	// JSON形式で出力
// 	result, err := json.MarshalIndent(recipeCollection, "", "  ")
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println(string(result))
// 	return nil
// }
