package main

import (
	"os"
	"recipe/models"

	// "recipe/services"
	"recipe/utils"
)

func main() {
	//TODO データベースを消す
	// os.Remove("recipe.db")

	// 初期化
	Init()

	// サーバー起動
	mainServer()
}

func DebugModel() {
	//データベース消去
	os.Remove("recipe.db")

	// モデル初期化
	models.Init()

	//デバッグ実行
	models.RunDebug()

}

func mainServer() {
	utils.Println("サーバーを起動しています")

	// サーバー初期化
	server := InitServer()

	// サーバー起動
	server.Logger.Fatal(server.Start(os.Getenv("BindAddr")))
}
