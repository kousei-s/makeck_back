package main

import (
	"context"
	"embed"
	"log"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/tools/filesystem"
)

// デフォルトファイル
var defaultIcons embed.FS

func InstallHook(app *pocketbase.PocketBase) {
	app.OnSettingsBeforeUpdateRequest().Add(func(evt *core.SettingsUpdateEvent) error {
		// log.Println("updateSetting")
		// log.Println(evt.HttpContext)
		// log.Println(evt.OldSettings)
		// log.Println(evt.NewSettings)
		return nil
	})

	// Oauth のフック
	app.OnRecordAfterAuthWithOAuth2Request().Add(func(evt *core.RecordAuthWithOAuth2Event) error {
		log.Println(evt.OAuth2User)

		// 新規作成の時
		if !evt.IsNewRecord {
			// 既にログイン済みの時
			return nil
		}

		// コンテキスト生成
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

		// URLからファイル
		AvaterData, err := filesystem.NewFileFromUrl(ctx, evt.OAuth2User.AvatarUrl)

		// エラー処理
		if err != nil {
			// 読み込みに失敗したとき (デフォルトにする)
			defaultIcon, err := defaultIcons.ReadFile("default/usericon.png")

			// エラー処理
			if err != nil {
				return err
			}

			// バイナリからファイルにする
			AvaterData, err = filesystem.NewFileFromBytes(defaultIcon, "default.png")

			// エラー処理
			if err != nil {
				return err
			}
		}

		log.Println("User Record")

		// フォーム生成
		UserForm := forms.NewRecordUpsert(app, evt.Record)

		// ファイルを追加
		UserForm.AddFiles("avatar", AvaterData)

		// 現在のデータ取得
		nowData := UserForm.Data()

		// 名前を変更
		nowData["name"] = evt.OAuth2User.Name

		log.Println(nowData)

		// データを保存
		UserForm.LoadData(nowData)

		// レコード更新
		if err := UserForm.Submit(); err != nil {
			return err
		}

		return nil
	})
}
