package models

import (
	"errors"
	"recipe/utils"
)

type Material struct {
	Uid         string `gorm:"primaryKey"`					//材料ID
	Name        string 										//材料名
	Count		string 										//個数
	Unit  		string 										//単位
	Processid   string                                      //手順と紐づけ
}

// データベースに材料を登録する関数
func Material_Register(name string,count string,unit string,process []Process) (string,error) {
	// レシピIDを生成
	uid, err := utils.Genid()
	if err != nil {
		return "",errors.New("uuid_create_error")
	}

	// 新しいレシピを作成
	newMaterial := Material{
		Uid:   uid,
		Name:  name,
		Count: count,
		Unit:  unit,
	}
	result := dbconn.Create(&newMaterial)
	
	if result.Error != nil {
		return "",result.Error
	}

	return uid,err
}
