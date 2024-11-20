package main

import (
	"log"
	"os"
)


func main() {
	// 初期化
	Init()

	log.Println("サーバーを起動しています")

	// サーバー初期化
	server := InitServer()

	// サーバー起動
	server.Logger.Fatal(server.Start(os.Getenv("BindAddr")))
}