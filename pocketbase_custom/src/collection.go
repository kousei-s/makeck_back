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

		form := forms.NewCollectionUpsert(app, new_collection)
		form.Name = "label"
		form.Type = models.CollectionTypeBase
		form.ListRule = nil
		form.ViewRule = types.Pointer("@request.auth.id != ''")
		form.CreateRule = nil //types.Pointer("")
		form.UpdateRule = nil //types.Pointer("@request.auth.id != ''")
		form.DeleteRule = nil
		form.Schema.AddField(&schema.SchemaField{
			Name:     "name",
			Type:     schema.FieldTypeText,
			Required: true,
			Options: &schema.TextOptions{},
		})

		// validate and submit (internally it calls app.Dao().SaveCollection(collection) in a transaction)
		if err := form.Submit(); err != nil {
			log.Println(err)
		}


		// ユーザーのコレクションを取得
		collection, err := app.Dao().FindCollectionByNameOrId("_pb_users_auth_")

		// エラー処理
		if err != nil {
			log.Println(err)
			// return err
		}

		// nil のとき
		if collection.Options["IsInit"] == nil || !collection.Options["IsInit"].(bool) {
			// Oauth 以外無効化
			collection.Options["allowEmailAuth"] = false
			collection.Options["allowUsernameAuth"] = false
			collection.Options["IsInit"] = true

			// ラベルを追加
			collection.Schema.AddField(&schema.SchemaField{
				Name:     "labels",
				Type:     schema.FieldTypeRelation,
				Required: false,
				Options:  &schema.RelationOptions{
					CollectionId: new_collection.Id,
					CascadeDelete: true,
				},
			})

			// データを更新
			err = app.Dao().SaveCollection(collection)

			// エラー処理
			if err != nil {
				log.Println(err)
				return nil
			}
		}

		return nil
	})
}

//go run -x . serve --http=0.0.0.0:8080