package controllers

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/microsoftonline"
)

var (
	// 認証に使うプロバイダ
	Provider string = ""
)

func Init() {
	// 認証バックエンドの設定
	key := os.Getenv("SECRETKEY")             // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30  // 30 days
	isProd := true       // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true   // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	// 認証プロバイダ初期化
	goth.UseProviders(
		microsoftonline.New(os.Getenv("MICROSOFT_KEY"), os.Getenv("MICROSOFT_SECRET"), os.Getenv("MICROSOFT_CALLBACK")),
	)

	// プロ台場設定
	Provider = os.Getenv("dprovider")
}