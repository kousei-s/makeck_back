package services

import (
	"encoding/json"
	"log"
)



func Debug() {
	recipe1 := Recipe{
		Uid:       "recipe1",
		Name:      "スパゲティ・ボロネーゼ",
		Image:     "/images/spaghetti_bolognese.jpg",
		LastState: Hot,
		Process: []Process{
			{
				Uid:         "process1",
				Name:        "肉と野菜を切る",
				Displayname: "下準備1",
				Description: "玉ねぎを1個、ニンジンを1本,ひき肉を200g切る",
				Parallel:    true,
				Time:        15,
				Tools:       []Tools{{Uid: "tool1", Name: "包丁", Count: 1, Unit: "本"}},
				Material: []Material{
					{Uid: "material1", Name: "ひき肉", Count: 200, Unit: "g", Processid: "process1"},
					{Uid: "material2", Name: "玉ねぎ", Count: 1, Unit: "個", Processid: "process1"},
					{Uid: "material3", Name: "にんじん", Count: 1, Unit: "本", Processid: "process1"},
				},
				Recipeid: "recipe1",
				Type:     Prepare,
			},
			{
				Uid:         "process2",
				Name:        "材料を炒め、トマトソースを加えて煮込む。",
				Displayname: "調理1",
				Description: "下準備で切った肉と野菜をにトマトソースを400gいれ、30分煮込む",
				Parallel:    true,
				Time:        30,
				Tools:       []Tools{{Uid: "tool2", Name: "鍋", Count: 1, Unit: "個"}},
				Material:    []Material{{Uid: "material4", Name: "トマトソース", Count: 400, Unit: "g", Processid: "process2"}},
				Recipeid:    "recipe1",
				Type:        Cook,
			},
		},
	}

	recipe2 := Recipe{
		Uid:       "recipe2",
		Name:      "チキンカレー",
		Image:     "/images/chicken_curry.jpg",
		LastState: Hot,
		Process: []Process{
			{
				Uid:         "process1",
				Name:        "鶏肉と野菜を切る",
				Displayname: "下準備1",
				Description: "玉ねぎを1個、じゃがいもを2個,鶏肉を300g切る",
				Parallel:    false,
				Time:        10,
				Tools:       []Tools{{Uid: "tool1", Name: "包丁", Count: 1, Unit: "本"}},
				Material: []Material{
					{Uid: "material1", Name: "鶏肉", Count: 300, Unit: "g", Processid: "process1"},
					{Uid: "material2", Name: "玉ねぎ", Count: 1, Unit: "個", Processid: "process1"},
					{Uid: "material3", Name: "じゃがいも", Count: 2, Unit: "個", Processid: "process1"},
				},
				Recipeid: "recipe2",
				Type:     Prepare,
			},
			{
				Uid:         "process2",
				Name:        "材料を炒めて、スパイスと水を加えて煮込む。",
				Displayname: "調理1",
				Description: "下準備で切った肉と野菜をにカレースパイスを1袋いれ、40分煮込む",
				Parallel:    true,
				Time:        40,
				Tools:       []Tools{{Uid: "tool2", Name: "鍋", Count: 1, Unit: "個"}},
				Material:    []Material{{Uid: "material4", Name: "カレースパイス", Count: 1, Unit: "袋", Processid: "process2"}},
				Recipeid:    "recipe2",
				Type:        Cook,
			},
		},
	}

	recipe3 := Recipe{
		Uid:       "recipe3",
		Name:      "シーザーサラダ",
		Image:     "/images/caesar_salad.jpg",
		LastState: Cool,
		Process: []Process{
			{
				Uid:         "process1",
				Name:        "野菜を切る",
				Displayname: "下準備1",
				Description: "レタスを1個、トマトを2個切る",
				Parallel:    true,
				Time:        5,
				Tools:       []Tools{{Uid: "tool1", Name: "包丁", Count: 1, Unit: "本"}},
				Material: []Material{
					{Uid: "material1", Name: "レタス", Count: 1, Unit: "個", Processid: "process1"},
					{Uid: "material2", Name: "トマト", Count: 2, Unit: "個", Processid: "process1"},
				},
				Recipeid: "recipe3",
				Type:     Prepare,
			},
			{
				Uid:         "process2",
				Name:        "ドレッシングを作る",
				Displayname: "下準備2",
				Description: "下準備で切った野菜をにオリーブオイルを50ml、酢を20mlいれ、混ぜる",
				Parallel:    false,
				Time:        5,
				Tools:       []Tools{{Uid: "tool3", Name: "ボウル", Count: 1, Unit: "個"}},
				Material: []Material{
					{Uid: "material3", Name: "オリーブオイル", Count: 50, Unit: "ml", Processid: "process2"},
					{Uid: "material4", Name: "酢", Count: 20, Unit: "ml", Processid: "process2"},
				},
				Recipeid: "recipe3",
				Type:     Prepare,
			},
			{
				Uid:         "process3",
				Name:        "野菜とドレッシングを混ぜる",
				Displayname: "調理1",
				Description: "下準備で切った野菜とドレッシングを混ぜる",
				Parallel:    false,
				Time:        10,
				Tools:       []Tools{{Uid: "tool4", Name: "サラダボウル", Count: 1, Unit: "個"}},
				Material: []Material{
					{Uid: "material1", Name: "レタス", Count: 1, Unit: "個", Processid: "process3"},
					{Uid: "material2", Name: "トマト", Count: 2, Unit: "個", Processid: "process3"},
					{Uid: "material5", Name: "ドレッシング", Count: 1, Unit: "杯", Processid: "process3"},
				},
				Recipeid: "recipe3",
				Type:     Cook,
			},
			{
				Uid:         "process4",
				Name:        "盛り付け",
				Displayname: "仕上げ",
				Description: "サラダを皿に盛り付ける。",
				Parallel:    false,
				Time:        5,
				Tools:       []Tools{{Uid: "tool5", Name: "皿", Count: 2, Unit: "枚"}},
				Material: []Material{
					{Uid: "material6", Name: "完成したサラダ", Count: 1, Unit: "皿", Processid: "process4"},
				},
				Recipeid: "recipe3",
				Type:     Finish,
			},
		},
	}

	recipe4 := Recipe{
		Uid:       "recipe4",
		Name:      "フルーツポンチ",
		Image:     "/images/fruits_punch.jpg",
		LastState: Cool,
		Process: []Process{
			{
				Uid:         "process1",
				Name:        "果物を切る",
				Displayname: "下準備1",
				Description: "イチゴを100g、キウイを2個、オレンジを1個切る。",
				Parallel:    true,
				Time:        10,
				Tools:       []Tools{{Uid: "tool1", Name: "包丁", Count: 1, Unit: "本"}},
				Material: []Material{
					{Uid: "material1", Name: "イチゴ", Count: 100, Unit: "g", Processid: "process1"},
					{Uid: "material2", Name: "キウイ", Count: 2, Unit: "個", Processid: "process1"},
					{Uid: "material3", Name: "オレンジ", Count: 1, Unit: "個", Processid: "process1"},
				},
				Recipeid: "recipe4",
				Type:     Prepare,
			},
			{
				Uid:         "process2",
				Name:        "混ぜる",
				Displayname: "調理1",
				Description: "フルーツを混ぜて、ジュースを加える。",
				Parallel:    false,
				Time:        5,
				Tools:       []Tools{{Uid: "tool3", Name: "ボウル", Count: 1, Unit: "個"}},
				Material:    []Material{{Uid: "material4", Name: "ジュース", Count: 200, Unit: "ml", Processid: "process2"}},
				Recipeid:    "recipe4",
				Type:        Cook,
			},
		},
	}

	recipes := []Recipe{
		recipe1,
		recipe2,
		recipe3,
		recipe4,
	}

	

	chart, err := chart_Register(recipes)
	if err != nil {
		log.Println(err)
	}
	// JSON形式で出力に変更して出力
	result, err := json.MarshalIndent(chart, "", "  ")
	if err != nil {
		log.Println(err)
	}

	log.Print(string(result))
}
