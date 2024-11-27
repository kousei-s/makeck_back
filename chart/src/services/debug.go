package services

import "time"

type LastSatate string

const (
	Hot    = LastSatate("hot")
	Cool   = LastSatate("cool")
	Normal = LastSatate("normal")
)

type Material struct {
	Uid         string `gorm:"primaryKey"`					//材料ID
	Name        string 										//材料名
	Count		int											//個数
	Unit  		string 										//単位
	Processid   string                                      //手順と紐づけ
}

// レシピテーブルの構造体宣言
type Recipe struct {
	Uid       string      `gorm:"primaryKey"` //レシピID
	Name      string      //料理名
	Image     string      //画像パス
	Category  []*Category `gorm:"many2many:recipe_category;foreignKey:uid"` //カテゴリー
	Process   []Process   `gorm:"foreignKey:recipeid"`                      //手順
	LastState LastSatate //最終状態
}

type Category struct {
	Id      int       `gorm:"primaryKey"` //カテゴリーID
	Name    string    //カテゴリー名
	Recipes []*Recipe `gorm:"many2many:recipe_category"`
}

type Tools struct {
	Uid         string `gorm:"primaryKey"`	//器具ID
	Name        string 						//器具名
	Count		int 						//個数
	Unit  		string 						//単位
	Processid   string                      //手順と紐づけ
}

type RecipeCategory struct {
	CategoryID int    `gorm:"primaryKey"`
	RecipeID   string `gorm:"primaryKey"`
	CreatedAt  time.Time
}

type Process struct {
	Uid      string     `gorm:"primaryKey"` //手順ID
	Name     string     //名前
	Description string  //説明
	Parallel bool       //平行可、不可
	Time     int        //所要時間
	Tools    []Tools    `gorm:"foreignKey:processid"` //必要器具
	Material []Material `gorm:"foreignKey:processid"` //材料
	Recipeid string     //レシピと紐づけ
}

func Debug() {
	
}