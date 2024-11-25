package models

import (
	"errors"
	"recipe/utils"
)

type Tools struct {
	Uid         string `gorm:"primaryKey"`	//器具ID
	Name        string 						//器具名
	Count		int 						//個数
	Unit  		string 						//単位
	Processid   string                      //手順と紐づけ
}

// データベースに器具を登録する処理
func Tools_Register(name string,count int,unit string) (string,error) {
	// レシピIDを生成
	uid, err := utils.Genid()
	if err != nil {
		return "",errors.New("uuid_create_error")
	}

	// 新しい器具を作成
	newTool := Tools{
		Uid:   uid,
		Name:  name,
		Count: count,
		Unit:  unit,
	}
	result := dbconn.Create(&newTool)
	
	if result.Error != nil {
		return "",result.Error
	}

	return uid,err
}
