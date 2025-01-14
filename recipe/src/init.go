package main

import (
	"log"
	"recipe/controllers"
	"recipe/models"
	"recipe/recipe_rpc"

	"github.com/joho/godotenv"
)

// .envを呼び出します。
func loadEnv() {
	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		log.Fatalf("読み込み出来ませんでした: %v", err)
	}

}

func Init() {
	//ENV を読み込み
	loadEnv()

	// コントローラー初期化
	controllers.Init()

	// モデル初期化
	models.Init()


	// GRPCサーバーを起動
	go recipe_rpc.RunServer()
}
