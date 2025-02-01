package models

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbconn *gorm.DB = nil
)

func Init() {
	// データベースを開く
	db, err := gorm.Open(sqlite.Open(os.Getenv("DBPATH")), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// グローバル変数に格納
	dbconn = db

	// マイグレーション
	db.AutoMigrate(&Recipe{})

	// マイグレーション
	db.AutoMigrate(&Category{})
	//カテゴリーテーブル作成
	initcategory()

	// マイグレーション
	db.AutoMigrate(&Material{})

	// マイグレーション
	db.AutoMigrate(&Process{})

	// マイグレーション
	db.AutoMigrate(&Tools{})

}