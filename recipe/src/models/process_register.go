package models

import (
	"errors"
	"recipe/utils"

	"gorm.io/gorm/clause"
)

type ProcessType int

const (
	CookProcess    ProcessType = 0
	PrepareProcess ProcessType = 1
	FinishProcess  ProcessType = 2
)

const ()

type Process struct {
	Uid         string     `gorm:"primaryKey"` //手順ID
	Name        string     //名前
	Description string     //説明
	Parallel    bool       //平行可、不可
	Time        int        //所要時間
	Tools       []Tools    `gorm:"foreignKey:processid"` //必要器具
	Material    []Material `gorm:"foreignKey:processid"` //材料
	Recipeid    string     //レシピと紐づけ
	Type        int        //手順の種類
}

type ProcessArgs struct {
	name        string     //名前
	description string     //説明
	parallel    bool       //平行可、不可
	time        int        //所要時間
	tools       []Tools    //必要器具
	materials   []Material //材料
	Type        int        //手順の種類
}

// データベースに手順を登録する関数
func Process_Register(args ProcessArgs) (Process, error) {
	// 手順IDを生成
	uid, err := utils.Genid()
	if err != nil {
		return Process{}, errors.New("uuid_create_error")
	}
	// 新しい手順を作成
	newProcess := Process{
		Uid:      uid,
		Name:     args.name,
		Parallel: args.parallel,
		Time:     args.time,
		Tools:    args.tools,
		Material: args.materials,
		Type:     args.Type,
	}
	result := dbconn.Create(&newProcess)

	if result.Error != nil {
		return Process{}, result.Error
	}
	return newProcess, err
}

func GetProcessById(processid string) (Process, error) {
	var process Process
	result := dbconn.Where(&Process{
		Uid: processid,
	}).Preload(clause.Associations).First(&process)
	return process, result.Error
}

func (proc *Process) GetTools() ([]Tools, error) {
	var tools []Tools
	result := dbconn.Where(&Tools{
		Processid: proc.Uid,
	}).Preload(clause.Associations).Find(&tools)
	return tools, result.Error
}

func (proc *Process) GetMaterials() ([]Material, error) {
	var materials []Material
	result := dbconn.Where(&Material{
		Processid: proc.Uid,
	}).Preload(clause.Associations).Find(&materials)
	return materials, result.Error
}