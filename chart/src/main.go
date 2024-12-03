package main

import (
	"chart/services"
	// "log"
	// "os"
)


func main() {
	// 初期化
	Init()

	services.Debug()

	// log.Println("サーバーを起動しています")

	// // サーバー初期化
	// server := InitServer()

	// // サーバー起動
	// server.Logger.Fatal(server.Start(os.Getenv("BindAddr")))
}