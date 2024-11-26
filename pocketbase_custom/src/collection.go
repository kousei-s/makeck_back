package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func CreateCollection(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(evt *core.ServeEvent) error {
		// ラベルを作成
		new_collection := &models.Collection{}

		// ユーザーのコレクションを取得
		collection, err := app.Dao().FindCollectionByNameOrId("_pb_users_auth_")

		// エラー処理
		if err != nil {
			log.Println(err)
			// return err
		}

		// nil のとき
		if collection.Options["IsInit"] == nil || !collection.Options["IsInit"].(bool) {
			// 初期化時
			// ラベルテーブルを作成する
			form := forms.NewCollectionUpsert(app, new_collection)
			form.Name = "label"
			form.Type = models.CollectionTypeBase
			form.ListRule = nil //管理者限定
			form.ViewRule = types.Pointer("@request.auth.id != ''")
			form.CreateRule = nil //管理者限定
			form.UpdateRule = nil //管理者限定
			form.DeleteRule = nil //管理者限定
			form.Schema.AddField(&schema.SchemaField{
				Name:     "name",
				Type:     schema.FieldTypeText,
				Required: true,
				Options:  &schema.TextOptions{},
			})

			// unique 追加
			form.Indexes = append(form.Indexes, "CREATE UNIQUE INDEX `unique_name` ON `label` (`name`)")
			

			// validate and submit (internally it calls app.Dao().SaveCollection(collection) in a transaction)
			if err := form.Submit(); err != nil {
				log.Println(err)
				return err
			}


			// プリセット追加
			err = InitLabel(app,new_collection)

			// エラー処理
			if err != nil {
				log.Println(err)
				return err
			}

			// Oauth 以外無効化
			collection.Options["allowEmailAuth"] = false
			collection.Options["allowUsernameAuth"] = false

			// 初期化済みにする
			collection.Options["IsInit"] = true

			// ラベルを追加
			collection.Schema.AddField(&schema.SchemaField{
				Name:     "labels",
				Type:     schema.FieldTypeRelation,
				Required: false,
				Options: &schema.RelationOptions{
					CollectionId:  new_collection.Id,
					CascadeDelete: true,
				},
			})

			// データを更新
			err = app.Dao().SaveCollection(collection)

			// エラー処理
			if err != nil {
				log.Println(err)
				return err
			}
		}

		return nil
	})
}

func CreateLabel(app *pocketbase.PocketBase, collection *models.Collection,name string) error {
	// admin ラベルを作成
	label_record := models.NewRecord(collection)
	// ラベル初期化
	new_label := forms.NewRecordUpsert(app, label_record)

	// データ挿入
	new_label.LoadData(map[string]any{
		"name": name,
	})

	// 保存
	// validate and submit (internally it calls app.Dao().SaveRecord(record) in a transaction)
	if err := new_label.Submit(); err != nil {
		return err
	}

	return nil
}

func InitLabel(app *pocketbase.PocketBase, collection *models.Collection) error {
	// admin ラベル作成
	err := CreateLabel(app,collection,"admin")

	// エラー処理
	if err != nil {
		return err
	}

	// owner ラベル作成
	err = CreateLabel(app,collection,"owner")

	// エラー処理
	if err != nil {
		return err
	}

	// subscriber ラベル作成
	err = CreateLabel(app,collection,"subscriber")

	// エラー処理
	if err != nil {
		return err
	}

	return nil
}

//go run -x . serve --http=0.0.0.0:8080
