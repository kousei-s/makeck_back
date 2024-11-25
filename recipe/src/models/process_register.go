package models

import (
	"errors"
	"recipe/utils"
)

type Process struct {
	Uid         string `gorm:"primaryKey"`				//手順ID
	Name        string 									//説明
	Parallel    bool									//平行可、不可
	Time        int 									//所要時間
	Tools	  []Tools   `gorm:"foreignKey:processid"`   //必要器具
	Material  []Material`gorm:"foreignKey:processid"`	//材料
	Recipeid  	string                                    //レシピと紐づけ
}

// データベースに手順を登録する関数
func Process_Register(name string,parallel bool,time int,tools []Tools,materials []Material) (string,error) {
	// 手順IDを生成
	uid, err := utils.Genid()
	if err != nil {
		return "",errors.New("uuid_create_error")
	}
	// 新しい手順を作成
	newProcess := Process{
		Uid:   uid,
		Name:  name,
		Parallel: parallel,
		Time: time,
		Tools: tools,
		Material: materials,
	}
	result := dbconn.Create(&newProcess)
	
	if result.Error != nil {
		return "",result.Error
	}
	return uid,err
}
