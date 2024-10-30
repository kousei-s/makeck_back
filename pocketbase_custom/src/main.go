package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	// "github.com/pocketbase/pocketbase/tools/types"
)

type UserData struct {
	UserID      string    // ユーザーID
	UserName    string    // ユーザー名
	Email       string    // メールアドレス
	Labels      []string  // ユーザーについたラベル
	ProviderUID string    // 認証プロバイダのユーザーID
	Provider    string    // 認証プロバイダ
	Created     time.Time // 作成日
	Updated     time.Time // 更新日
}

type UserInfo struct {
	UserName string   // ユーザー名
	Labels   []string // ユーザーについたラベル
}

func main() {
	app := pocketbase.New()

	// ユーザー認証
	app.OnBeforeServe().Add(func(evt *core.ServeEvent) error {
		// ユーザーを認証する関数
		evt.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/jwt",
			Handler: func(ctx echo.Context) error {
				// ユーザー取得
				user := ctx.Get(apis.ContextAuthRecordKey).(*models.Record)

				// 認証プロバイダの情報取得
				records, err := app.Dao().FindAllExternalAuthsByRecord(user)

				// エラー処理
				if err != nil {
					log.Println(err)
				}

				// 最初のレコード取得
				provider := records[0]

				// 返すデータ
				return_data := UserData{
					UserID:      user.Id,
					UserName:    user.Username(),
					Email:       user.Email(),
					Labels:      []string{},
					ProviderUID: provider.ProviderId,
					Provider:    provider.Provider,
					Created:     user.GetCreated().Time(),
					Updated:     user.GetUpdated().Time(),
				}

				return ctx.JSON(http.StatusOK, echo.Map{
					"result": return_data,
				})
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
				apis.RequireRecordAuth("users"),
			},
		})

		evt.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/info/:userid",
			Handler: func(ctx echo.Context) error {
				// 画像を取得するコレクション
				targetCollection := "_pb_users_auth_"

				// ユーザーID取得
				userid := ctx.PathParam("userid")

				// ユーザーを取得する
				user, err := app.Dao().FindRecordById(targetCollection, userid)

				// エラー処理
				if err != nil {
					log.Println(err)
					return ctx.JSON(http.StatusInternalServerError, echo.Map{
						"result": "Failed to get user",
					})
				}

				return ctx.JSON(http.StatusOK, echo.Map{
					"result": UserInfo{
						UserName: user.Username(),
						Labels:   []string{},
					},
				})
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		evt.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/icon/:userid",
			Handler: func(ctx echo.Context) error {
				// 画像を取得するコレクション
				targetCollection := "_pb_users_auth_"

				// ユーザーID取得
				userid := ctx.PathParam("userid")

				// ユーザーを取得する
				user, err := app.Dao().FindRecordById(targetCollection, userid)

				// エラー処理
				if err != nil {
					log.Println(err)
					return ctx.JSON(http.StatusInternalServerError, echo.Map{
						"result": "Failed to get user",
					})
				}

				// アバターのURL
				avatar := app.Settings().Meta.AppUrl + "/api/files/" + user.Collection().Id + "/" + user.Id + "/" + user.GetString("avatar")

				return ctx.Redirect(http.StatusTemporaryRedirect, avatar)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})

	// ラベルを生成
	CreateCollection(app)

	// 設定のフックをインストール
	InstallHook(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
