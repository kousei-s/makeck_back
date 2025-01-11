package models

import (
	"errors"
	"recipe/utils"
)

type Material struct {
	Uid       string  `gorm:"primaryKey"` //材料ID
	Name      string  //材料名
	Count     float32 //個数
	Unit      string  //単位
	Processid string  //手順と紐づけ
}

type MaterialArgs struct {
	name  string
	count float32
	unit  string
}

// データベースに材料を登録する関数
func Material_Register(args MaterialArgs) (Material, error) {

	// レシピIDを生成
	uid, err := utils.Genid()
	if err != nil {
		return Material{}, errors.New("uuid_create_error")
	}

	material := Material{}

	//材料に同じものが既に存在してるかどうか
	material_len := dbconn.Where(Material{Name: args.name}).Where(Material{Count: args.count}).Where(Material{Unit: args.unit}).Find(&material).RowsAffected

	//既にある場合そのuidを返す
	if material_len != 0 {
		return material, nil
	}

	// 新しいレシピを作成
	newMaterial := Material{
		Uid:   uid,
		Name:  args.name,
		Count: args.count,
		Unit:  args.unit,
	}
	result := dbconn.Create(&newMaterial)

	if result.Error != nil {
		return Material{}, result.Error
	}

	return newMaterial, nil
}
