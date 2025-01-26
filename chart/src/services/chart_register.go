package services

import (

)

// Recipe はレシピ情報を表します
type Recipes struct {
	Uid     string    `json:"uid"`     // レシピのユニークID
	Name    string    `json:"name"`    // レシピ名
	Process []Process `json:"process"` // プロセスのリスト
}

// Process はレシピのプロセス情報を表します
type Processes struct {
	Uid         string      `json:"uid"`         // プロセスのユニークID
	Name        string      `json:"name"`        // プロセス名
	Displayname string      `json:"displayname"` // 表示名
	Description string      `json:"description"` // 説明
	Parallel    bool        `json:"parallel"`    // 並行処理できるかどうか
	Time        int         `json:"time"`        // 所要時間
	Type        ProcessType `json:"type"`        // 手順の種類
}

// Task はタスク情報を表します
type Task struct {
	Tejuns    map[string]Tejun `json:"tejuns"`    // 手順情報
	StartTime int              `json:"startTime"` // 開始時間
}

// Tejun は手順情報を表します
type Tejun struct {
	Id       string `json:"id,omitempty"`   // 手順ID (オプション)
	Name     string `json:"name,omitempty"` // 手順名 (オプション)
	Time     int    `json:"time,omitempty"` // 所要時間 (オプション)
	Parallel bool   `json:"parallel"`       // 並行処理できるかどうか
}

// RecipeCollection はレシピとタスクのコレクションを表します
type RecipeCollection struct {
	Recipes []ShortRecipe `json:"recipies"` // レシピのリスト
	Tasks   []Task        `json:"tasks"`    // タスクのリスト
}

// ShortRecipe は簡略化されたレシピ情報を表します
type ShortRecipe struct {
	Uid        string // レシピのユニークID
	Name       string // レシピ名
	LastSatate LastSatate
	Divide     []TaskDivide // タスクの分割情報
	Times      int          // 累計時間
}

// TaskDivide はタスクの分割情報を表します
type TaskDivide struct {
	Uid         string      // タスクのユニークID
	Time        int         // 所要時間
	Parallel    bool        // 並行処理かどうか
	Type        ProcessType // 手順の種類
	Displayname string      // 表示名
}

// chart_Register はレシピからタスクを生成するメソッド
func chart_Register(recipes []Recipe) (RecipeCollection, error) {
	// レシピを簡略化した形式に変換
	simple_recipe, err := chart_Extraction(recipes)
	if err != nil {
		return RecipeCollection{}, err
	}

	// 優先度を決定するメソッド
	prioritys_recepi, err := chart_Priority(simple_recipe)
	if err != nil {
		return RecipeCollection{}, err
	}

	tasks := []Task{}

	recipe_nums := []int{
		len(simple_recipe[0].Divide),
		len(simple_recipe[1].Divide),
		len(simple_recipe[2].Divide),
		len(simple_recipe[3].Divide),
	}

	//現在のタスクの位置(次何番目のレシピを入れるか)
	location := 0

	//最大時間
	max_times := 0

	//最初のタスクのだけtrue
	frist_task := true
	for {
		// タスクの横列を一つ生成
		task_frame := []Task{
			{
				Tejuns: make(map[string]Tejun), // 手順情報のマップを初期化
			},
		}
		// 優先度に基づいてタスクを生成
		task, locations, max_time, err := chart_CreateTask(simple_recipe, prioritys_recepi, recipe_nums, task_frame, location, max_times, frist_task) // Task型を取得
		if err != nil {
			return RecipeCollection{}, err
		}
		max_times += max_time
		tasks = append(task, tasks...) // タスクを追加
		location = locations
		//全部使ったら無限ループを抜ける
		if recipe_nums[0] == 0 && recipe_nums[1] == 0 && recipe_nums[2] == 0 && recipe_nums[3] == 0 {
			break
		}
	}

	new_tasks, err := chart_createtime(tasks)

	if err != nil {
		return RecipeCollection{}, err
	}

	// タスクコレクションを生成
	tasks_collection := RecipeCollection{
		Recipes: simple_recipe,
		Tasks:   new_tasks,
	}

	return tasks_collection, nil // 生成したタスクを返す
}

// chart_CreateTask は優先度に基づいて横一列のタスクを作成するメソッド
func chart_CreateTask(recipes []ShortRecipe, priorities []string, recipe_nums []int, task_frame []Task, location int, max_time int, frist_task bool) ([]Task, int, int, error) {
	task_bool := true
	//順番を前回のタスクにする
	new_list := priorities[location:]
	new_list = append(new_list, priorities[:location]...)
	// sample3のインデックスを探す
	index := -1
	for i, v := range priorities {
		if v == new_list[location] {
			index = i + 1
			index %= len(recipes)
			break
		}
	}

	temp_time := max_time
	for i, priority := range new_list {
		i += location
		i %= len(recipes)
		if task_bool && recipe_nums[i] > 0 && (recipes[i].Divide[recipe_nums[i]-1].Time <= max_time || frist_task) {
			task_frame[0].Tejuns[priority] = Tejun{
				Id:       recipes[i].Divide[recipe_nums[i]-1].Uid,
				Name:     string(recipes[i].Divide[recipe_nums[i]-1].Displayname),
				Time:     recipes[i].Divide[recipe_nums[i]-1].Time,
				Parallel: recipes[i].Divide[recipe_nums[i]-1].Parallel,
			}

			recipe_nums[i]--

			// 並行処理可の場合は、次の手順を追加
			if !(recipes[i].Divide[recipe_nums[i]].Parallel) {
				task_bool = false
				location = i + 1
				location %= len(recipes)
			}

			//前述までのタスクより時間を越さないため
			if recipes[i].Divide[recipe_nums[i]].Time > temp_time {
				temp_time = recipes[i].Divide[recipe_nums[i]].Time
			}
		} else {
			task_frame[0].Tejuns[priority] = Tejun{
				Parallel: true,
			}
		}

	}
	frist_task = false
	max_time = temp_time
	return task_frame, location, max_time, nil
}

// chart_Extraction はレシピから必要な情報を抽出するメソッド
func chart_Extraction(recipes []Recipe) ([]ShortRecipe, error) {
	shortRecipes := make([]ShortRecipe, len(recipes)) // 簡略化レシピのスライスを初期化
	for i, recipe := range recipes {
		times := 0
		tasks := make([]TaskDivide, len(recipe.Process)) // タスクのスライスを初期化
		for j := range recipe.Process {
			tasks[j] = TaskDivide{
				Uid:         recipe.Process[j].Uid,
				Time:        recipe.Process[j].Time,
				Parallel:    recipe.Process[j].Parallel,
				Type:        recipe.Process[j].Type,
				Displayname: recipe.Process[j].Displayname,
			}
			times += recipe.Process[j].Time // 所要時間を合計
		}
		shortRecipes[i] = ShortRecipe{
			Uid:        recipe.Uid,
			Name:       recipe.Name,
			LastSatate: recipe.LastState,
			Divide:     tasks,
			Times:      times,
		}
	}

	return shortRecipes, nil // 簡略化レシピを返す
}

// chart_Priority はレシピの優先度を決定するメソッド
func chart_Priority(recipes []ShortRecipe) ([]string, error) {
	// 各状態に対応するレシピのUIDを格納するマップ
	stateMap := map[string][]string{
		"hot":    {},
		"reheat": {},
		"normal": {},
		"cool":   {},
	}

	// 各レシピの状態を確認し、マップに追加
	for _, recipe := range recipes {
		switch recipe.LastSatate {
		case "hot":
			stateMap["hot"] = append(stateMap["hot"], recipe.Uid)
		case "reheat":
			stateMap["reheat"] = append(stateMap["reheat"], recipe.Uid)
		case "normal":
			stateMap["normal"] = append(stateMap["normal"], recipe.Uid)
		case "cool":
			stateMap["cool"] = append(stateMap["cool"], recipe.Uid)
		}
	}

	// 優先度の順序でUIDを連結
	priority := append(append(append(stateMap["hot"], stateMap["reheat"]...), stateMap["normal"]...), stateMap["cool"]...)

	return priority, nil // 優先度リストを返す
}

// chart_createtime はタスクの時間を計算し、各タスクの開始時間を設定するメソッド
func chart_createtime(tasks []Task) ([]Task, error) {
	startTime := 0 // 現在の開始時間を保持

	status := false // タスクが追加されたかどうかを示すフラグ

	new_tasks := []Task{}


	// 各タスクに対してループ
	for i := 0; i < len(tasks); i++ {
		task := tasks[i]
		// 各タスクの手順に対してループ
		for _, tejun := range task.Tejuns {
			// 手順が並行処理可能であり、次のタスクが存在する場合
			if !(tejun.Parallel) && i < len(tasks)-1 {
				startTime += tejun.Time          // 開始時間を更新
				tasks[i + 1].StartTime = startTime // 次のタスクの開始時間を設定
				status = true                    // タスクが追加されたことを示すフラグ
			}
		}

		// もしタスクが追加された場合
		if status {
			new_tasks = append(new_tasks, task) // 現在のタスクを新しいタスクリストに追加
			continue                            // 次のタスクへ
		}

		// タスクの手順の中で最大の時間を見つける
		temp := 0 // 一時的な時間を保持する変数
		for _, tejun := range task.Tejuns {
			if temp < tejun.Time { // 最大の手順時間を見つける
				temp = tejun.Time
				startTime += temp // 開始時間を加算
				tasks[i + 1].StartTime = startTime
			}
		}

		// 新しいタスクを作成し、手順を設定
		new_tasks = append(new_tasks, Task{
			Tejuns:    tasks[i].Tejuns, // 次のタスクの手順を設定
		})
	}
	return new_tasks, nil // 新しいタスクリストを返す
}
