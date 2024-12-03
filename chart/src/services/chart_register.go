package services

import (
	"chart/utils"
	"encoding/json"
	"fmt"
)

// Recipe はレシピ情報を表します
type Recipes struct {
	Uid     string    `json:"uid"`     // レシピのユニークID
	Name    string    `json:"name"`    // レシピ名
	Process []Process `json:"process"` // プロセスのリスト
}

// Process はレシピのプロセス情報を表します
type Processes struct {
	Uid         string `json:"uid"`         // プロセスのユニークID
	Name        string `json:"name"`        // プロセス名
	Description string `json:"description"` // 説明
	Parallel    bool   `json:"parallel"`    // 並行処理かどうか
	Time        int    `json:"time"`        // 所要時間
}

// Task はタスク情報を表します
type Task struct {
	Name      string           `json:"name"`      // タスク名
	Tejuns    map[string]Tejun `json:"tejuns"`    // 手順情報
	StartTime int              `json:"startTime"` // 開始時間
}

// Tejun は手順情報を表します
type Tejun struct {
	Id   string `json:"id,omitempty"`   // 手順ID (オプション)
	Name string `json:"name,omitempty"` // 手順名 (オプション)
	Time int    `json:"time,omitempty"` // 所要時間 (オプション)
}

// RecipeCollection はレシピとタスクのコレクションを表します
type RecipeCollection struct {
	Recipes []ShortRecipe `json:"recipies"` // レシピのリスト
	Tasks   []TaskDivide  `json:"tasks"`    // タスクのリスト
}

// ShortRecipe は簡略化されたレシピ情報を表します
type ShortRecipe struct {
	Uid    string // レシピのユニークID
	Name   string // レシピ名
	LastSatate LastSatate
	Divide []TaskDivide
}

// 判別に必要な構造体
type TaskDivide struct {
	Uid      string
	Time     int
	Parallel bool
}

func chart_Register(recipes []Recipe) ([]Task, error) {

	//必要な値のみに整理する
	simple_recipe,err := chart_Extraction(recipes)
	if err != nil {
		return []Task{}, err
	}
	_= simple_recipe[0].LastSatate 



	return []Task{}, nil
}


//必要な値のみに整理する
func chart_Extraction(recipes []Recipe)  ([]ShortRecipe, error) {
	shortRecipes := make([]ShortRecipe, len(recipes))
	for i, recipe := range recipes {
		tasks := make([]TaskDivide, len(recipe.Process))
		for j := range len(recipe.Process) {
			uuid,err := utils.Genid()
			if err != nil {
				return []ShortRecipe{}, err
			}
			tasks[j] = TaskDivide{
				Uid:      uuid,
				Time:     recipe.Process[j].Time,
				Parallel: recipe.Process[j].Parallel,
			}
		}
		shortRecipes[i] = ShortRecipe{
			Uid:    recipe.Uid,
			Name:   recipe.Name,
			LastSatate: recipe.LastState,
			Divide: tasks,
		}
	}
	


	// JSON形式で出力
	result, err := json.MarshalIndent(shortRecipes, "", "  ")
	if err != nil {
		return []ShortRecipe{}, err
	}

	fmt.Println(string(result))

	return shortRecipes, nil
}
