package main

import (
	"log"
	"os"
	"recipe/models"
)

func main() {
	// 初期化
	Init()

	// モデルのテスト実行
	DebugModel()

	// サーバー起動
	// mainServer()
}

func DebugModel() {
	// モデル初期化
	models.Init()

	//でバッグ実行
	models.RunDebug()
}

func mainServer() {
	log.Println("サーバーを起動しています")

	// サーバー初期化
	server := InitServer()

	// サーバー起動
	server.Logger.Fatal(server.Start(os.Getenv("BindAddr")))
}