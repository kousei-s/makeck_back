package services

type LastSatate string

const (
	Hot    = LastSatate("hot")
	Reheat = LastSatate("reheat")
	Cool   = LastSatate("cool")
	Normal = LastSatate("normal")
)

type ProcessType string

const (
	// 調理
	Cook = ProcessType("cook")
	// 下準備
	Prepare = ProcessType("prepare")
	// 仕上げ
	Finish = ProcessType("finish")	
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
	Uid       string     `gorm:"primaryKey"` //レシピID
	Name      string     //料理名
	Image     string     //画像パス
	Process   []Process  `gorm:"foreignKey:recipeid"` //手順
	LastState LastSatate //最終状態
}

type Tools struct {
	Uid       string `gorm:"primaryKey"` //器具ID
	Name      string //器具名
	Count     int    //個数
	Unit      string //単位
	Processid string //手順と紐づけ
}

type Process struct {
	Uid         string      `gorm:"primaryKey"` //手順ID
	Name        string      //名前
	Displayname string      //表示名
	Description string      //説明
	Parallel    bool        //平行可、不可
	Time        int         //所要時間
	Tools       []Tools     `gorm:"foreignKey:processid"` //必要器具
	Material    []Material  `gorm:"foreignKey:processid"` //材料
	Recipeid    string      //レシピと紐づけ
	Type        ProcessType `json:"type"` // 手順の種類
}