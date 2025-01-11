package models

import (
	"errors"
	"recipe/utils"
)

type Tools struct {
	Uid       string  `gorm:"primaryKey"` //器具ID
	Name      string  //器具名
	Count     float32 //個数
	Unit      string  //単位
	Processid string  //手順と紐づけ
}

type ToolArgs struct {
	name  string
	count float32
	unit  string
}

// データベースに器具を登録する処理
func Tools_Register(args ToolArgs) (Tools, error) {
	// レシピIDを生成
	uid, err := utils.Genid()
	if err != nil {
		return Tools{}, errors.New("uuid_create_error")
	}

	tool := Tools{}

	//toolに同じものが既に存在してるかどうか
	tool_len := dbconn.Where(Material{Name: args.name}).Where(Material{Count: args.count}).Where(Material{Unit: args.unit}).Find(&tool).RowsAffected

	//既にある場合そのuidを返す
	if tool_len != 0 {
		return tool, nil
	}

	// 新しい器具を作成
	newTool := Tools{
		Uid:   uid,
		Name:  args.name,
		Count: args.count,
		Unit:  args.unit,
	}
	result := dbconn.Create(&newTool)

	if result.Error != nil {
		return Tools{}, result.Error
	}

	return newTool, err
}
