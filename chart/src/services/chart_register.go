package services

import "log"

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
	TotalTime int         `json:"totaltime"` // 累計時間
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

	log.Println("優先順",prioritys_recepi)

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
	first_task := true
	for {
		// タスクの横列を一つ生成
		task_frame := []Task{
			{
				Tejuns: make(map[string]Tejun), // 手順情報のマップを初期化
			},
		}

		// 優先度に基づいてタスクを生成
		task, locations, max_time,first, err := chart_CreateTask(simple_recipe, prioritys_recepi, recipe_nums, task_frame, location, max_times, first_task) // Task型を取得
		if err != nil {
			return RecipeCollection{}, err
		}
		first_task = first
		max_times += max_time
		tasks = append(task, tasks...) // タスクを追加
		location = locations
		//全部使ったら無限ループを抜ける
		if recipe_nums[0] == 0 && recipe_nums[1] == 0 && recipe_nums[2] == 0 && recipe_nums[3] == 0 {
			break
		}
	}

	new_tasks,totalTime, err := chart_createtime(tasks)

	if err != nil {
		return RecipeCollection{}, err
	}



	// タスクコレクションを生成
	tasks_collection := RecipeCollection{
		Recipes:   simple_recipe,
		Tasks:     new_tasks,
		TotalTime: totalTime,
	}

	return tasks_collection, nil // 生成したタスクを返す
}

// chart_CreateTask は優先度に基づいて横一列のタスクを作成するメソッド
func chart_CreateTask(recipes []ShortRecipe, priorities []string, recipe_nums []int, task_frame []Task, location int, max_time int, first_task bool) ([]Task, int, int,bool, error) {
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
	first_time := 0

	for i, priority := range new_list {
		// current index iをlocationで調整し、recipesの長さで割った余りを取る
		i += location
		i %= len(recipes)
		
		// タスクが有効で、レシピが存在し、時間条件を満たす場合
		if task_bool && recipe_nums[i] > 0 && (recipes[i].Divide[recipe_nums[i]-1].Time <= max_time || first_task) {			
			// 初回タスクの場合の処理
			if first_task {
				// 初めての時間設定
				if first_time == 0 {
					first_time = recipes[i].Divide[recipe_nums[i]-1].Time
				// 既に初回の時間が設定されている場合
				} else if first_time < recipes[i].Divide[recipe_nums[i]-1].Time {
					// 時間が初回より長い場合は空を設定
					task_frame[0].Tejuns[priority] = Tejun{
						Parallel: true,
					}
					continue
				}
			}
			
			// タスクフレームにタスクを追加
			task_frame[0].Tejuns[priority] = Tejun{
				Id:       recipes[i].Divide[recipe_nums[i]-1].Uid,
				Name:     string(recipes[i].Divide[recipe_nums[i]-1].Displayname),
				Time:     recipes[i].Divide[recipe_nums[i]-1].Time,
				Parallel: recipes[i].Divide[recipe_nums[i]-1].Parallel,
			}
	
			// 使用したレシピのカウントを減少
			recipe_nums[i]--
	
			// 並行処理ができない場合の処理
			if !(recipes[i].Divide[recipe_nums[i]].Parallel) {
				task_bool = false
				location = i + 1
				location %= len(recipes)
			}
	
			if first_task {
				temp_time += first_time
			}

			// 時間制限を超えないように調整
			if recipes[i].Divide[recipe_nums[i]].Time > temp_time {
				temp_time = recipes[i].Divide[recipe_nums[i]].Time
			}


			
		} else {
			// 条件を満たさない場合は並行処理として設定
			task_frame[0].Tejuns[priority] = Tejun{
				Parallel: true,
			}
		}
	}
	

	first_task = false
	max_time = temp_time

	return task_frame, location, max_time,first_task,nil
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
		log.Print(recipe.Name, recipe.LastSatate)
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
func chart_createtime(tasks []Task) ([]Task,int, error) {
	startTime := 0 // 現在の開始時間を保持

	status := false // タスクが追加されたかどうかを示すフラグ

	new_tasks := []Task{}

	totaltime := 0


	// 各タスクに対してループ
	for i, task := range tasks  { 
		// 各タスクの手順に対してループ
		for _, tejun := range task.Tejuns {
			// 手順が並行処理不可能であり、次のタスクが存在する場合
			if !(tejun.Parallel) && i < len(tasks)-1 {
				startTime += tejun.Time          // 開始時間を更新
				tasks[i + 1].StartTime = startTime // 次のタスクの開始時間を設定
				status = true                    // タスクが追加されたことを示すフラグ
			}

			if i == len(tasks)-1 && totaltime < tejun.Time {
				totaltime = tejun.Time
			}
			
		}

		// もしタスクが追加された場合
		if status {
			new_tasks = append(new_tasks, task) // 現在のタスクを新しいタスクリストに追加
			continue                            // 次のタスクへ
		}

		// タスクの手順の中で最大の時間を見つける
		temptime := 0 // 一時的な時間を保持する変数
		for _, tejun := range task.Tejuns {
			if temptime < tejun.Time { // 最大の手順時間を見つける
				temptime = tejun.Time
				startTime += temptime // 開始時間を加算
				tasks[i + 1].StartTime = startTime
			}
		}

		// 新しいタスクを作成し、手順を設定
		new_tasks = append(new_tasks, Task{
			Tejuns:    tasks[i].Tejuns, // 次のタスクの手順を設定
		})
	}

	totaltime += startTime

	for _,task := range tasks {
		//手順ごとの最大の時間
		maxtime := 0
		for _ , tejun := range task.Tejuns {
			if maxtime < tejun.Time {
				maxtime = tejun.Time
			}
		}

		if totaltime < task.StartTime + maxtime {
			totaltime = task.StartTime + maxtime
		}
	}
	
	return new_tasks,totaltime, nil // 新しいタスクリストを返す
}