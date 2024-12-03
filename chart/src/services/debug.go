package services

import (
	"log"
)

type LastSatate string

const (
	Hot    = LastSatate("hot")
	Cool   = LastSatate("cool")
	Normal = LastSatate("normal")
)

type Material struct {
	Uid       string `gorm:"primaryKey"` //材料ID
	Name      string //材料名
	Count     int    //個数
	Unit      string //単位
	Processid string //手順と紐づけ
}

// レシピテーブルの構造体宣言
type Recipe struct {
	Uid       string      `gorm:"primaryKey"` //レシピID
	Name      string      //料理名
	Image     string      //画像パス
	Process   []Process   `gorm:"foreignKey:recipeid"`                      //手順
	LastState LastSatate  //最終状態
}


type Tools struct {
	Uid       string `gorm:"primaryKey"` //器具ID
	Name      string //器具名
	Count     int    //個数
	Unit      string //単位
	Processid string //手順と紐づけ
}

type Process struct {
	Uid         string     `gorm:"primaryKey"` //手順ID
	Name        string     //名前
	Description string     //説明
	Parallel    bool       //平行可、不可
	Time        int        //所要時間
	Tools       []Tools    `gorm:"foreignKey:processid"` //必要器具
	Material    []Material `gorm:"foreignKey:processid"` //材料
	Recipeid    string     //レシピと紐づけ
}

func Debug()() {
	recipe1 := Recipe{
		Uid:       "recipe1",
		Name:      "スパゲティ・ボロネーゼ",
		Image:     "/images/spaghetti_bolognese.jpg",
		LastState: Hot,
		Process: []Process{
			{
				Uid:         "process1",
				Name:        "材料の準備",
				Description: "野菜と肉を刻む。",
				Parallel:    false,
				Time:        15,
				Tools:       []Tools{{Uid: "tool1", Name: "包丁", Count: 1, Unit: "本"}},
				Material: []Material{
					{Uid: "material1", Name: "ひき肉", Count: 200, Unit: "g", Processid: "process1"},
					{Uid: "material2", Name: "玉ねぎ", Count: 1, Unit: "個", Processid: "process1"},
					{Uid: "material3", Name: "にんじん", Count: 1, Unit: "本", Processid: "process1"},
				},
				Recipeid: "recipe1",
			},
			{
				Uid:         "process2",
				Name:        "煮込む",
				Description: "材料を炒め、トマトソースを加えて煮込む。",
				Parallel:    true,
				Time:        30,
				Tools:       []Tools{{Uid: "tool2", Name: "鍋", Count: 1, Unit: "個"}},
				Material:    []Material{{Uid: "material4", Name: "トマトソース", Count: 400, Unit: "g", Processid: "process2"}},
				Recipeid:    "recipe1",
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
				Name:        "材料の準備",
				Description: "鶏肉と野菜を切る。",
				Parallel:    false,
				Time:        10,
				Tools:       []Tools{{Uid: "tool1", Name: "包丁", Count: 1, Unit: "本"}},
				Material: []Material{
					{Uid: "material1", Name: "鶏肉", Count: 300, Unit: "g", Processid: "process1"},
					{Uid: "material2", Name: "玉ねぎ", Count: 1, Unit: "個", Processid: "process1"},
					{Uid: "material3", Name: "じゃがいも", Count: 2, Unit: "個", Processid: "process1"},
				},
				Recipeid: "recipe2",
			},
			{
				Uid:         "process2",
				Name:        "煮込む",
				Description: "材料を炒めて、スパイスと水を加えて煮込む。",
				Parallel:    true,
				Time:        40,
				Tools:       []Tools{{Uid: "tool2", Name: "鍋", Count: 1, Unit: "個"}},
				Material:    []Material{{Uid: "material4", Name: "カレースパイス", Count: 1, Unit: "袋", Processid: "process2"}},
				Recipeid:    "recipe2",
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
				Name:        "材料の準備",
				Description: "野菜を切る。",
				Parallel:    true,
				Time:        5,
				Tools:       []Tools{{Uid: "tool1", Name: "包丁", Count: 1, Unit: "本"}},
				Material: []Material{
					{Uid: "material1", Name: "レタス", Count: 1, Unit: "個", Processid: "process1"},
					{Uid: "material2", Name: "トマト", Count: 2, Unit: "個", Processid: "process1"},
				},
				Recipeid: "recipe3",
			},
			{
				Uid:         "process2",
				Name:        "ドレッシングを作る",
				Description: "ドレッシングを混ぜる。",
				Parallel:    false,
				Time:        5,
				Tools:       []Tools{{Uid: "tool3", Name: "ボウル", Count: 1, Unit: "個"}},
				Material: []Material{
					{Uid: "material3", Name: "オリーブオイル", Count: 50, Unit: "ml", Processid: "process2"},
					{Uid: "material4", Name: "酢", Count: 20, Unit: "ml", Processid: "process2"},
				},
				Recipeid: "recipe3",
			},
			{
				Uid:         "process3",
				Name:        "サラダを組み立てる",
				Description: "切った野菜とドレッシングを混ぜる。",
				Parallel:    false,
				Time:        10,
				Tools:       []Tools{{Uid: "tool4", Name: "サラダボウル", Count: 1, Unit: "個"}},
				Material: []Material{
					{Uid: "material1", Name: "レタス", Count: 1, Unit: "個", Processid: "process3"},
					{Uid: "material2", Name: "トマト", Count: 2, Unit: "個", Processid: "process3"},
					{Uid: "material5", Name: "ドレッシング", Count: 1, Unit: "杯", Processid: "process3"},
				},
				Recipeid: "recipe3",
			},
			{
				Uid:         "process4",
				Name:        "盛り付け",
				Description: "サラダを皿に盛り付ける。",
				Parallel:    false,
				Time:        5,
				Tools:       []Tools{{Uid: "tool5", Name: "皿", Count: 2, Unit: "枚"}},
				Material: []Material{
					{Uid: "material6", Name: "完成したサラダ", Count: 1, Unit: "皿", Processid: "process4"},
				},
				Recipeid: "recipe3",
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
				Name:        "材料の準備",
				Description: "フルーツを切る。",
				Parallel:    true,
				Time:        10,
				Tools:       []Tools{{Uid: "tool1", Name: "包丁", Count: 1, Unit: "本"}},
				Material: []Material{
					{Uid: "material1", Name: "イチゴ", Count: 100, Unit: "g", Processid: "process1"},
					{Uid: "material2", Name: "キウイ", Count: 2, Unit: "個", Processid: "process1"},
					{Uid: "material3", Name: "オレンジ", Count: 1, Unit: "個", Processid: "process1"},
				},
				Recipeid: "recipe4",
			},
			{
				Uid:         "process2",
				Name:        "混ぜる",
				Description: "フルーツを混ぜて、ジュースを加える。",
				Parallel:    false,
				Time:        5,
				Tools:       []Tools{{Uid: "tool3", Name: "ボウル", Count: 1, Unit: "個"}},
				Material:    []Material{{Uid: "material4", Name: "ジュース", Count: 200, Unit: "ml", Processid: "process2"}},
				Recipeid:    "recipe4",
			},
		},
	}

	recipes := []Recipe{
		recipe1,
		recipe2,
		recipe3,
		recipe4,
	}

	json,err := chart_Register(recipes)
	if err != nil {
		log.Println(err)
	}
	log.Print(json)

	// err := chart_Register(recipes)
	// log.Print(err)
}


