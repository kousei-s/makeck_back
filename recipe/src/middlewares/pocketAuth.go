package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"recipe/utils"
	"time"

	"github.com/labstack/echo/v4"
)

type ReturnResult struct {
	Result UserData `json:"result"`
}

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

func PocketAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// ヘッダ取得
			token := ctx.Request().Header.Get("Authorization")

			// トークンを検証する
			user, err := VerifyToken(token)

			// エラー処理
			if err != nil {
				utils.Println(err)
				return ctx.NoContent(http.StatusUnauthorized)
			}

			// ユーザーを設定
			ctx.Set("user", user)

			return next(ctx)
		}
	}
}

func VerifyToken(token string) (UserData, error) {
	// リクエスト送信
	req, _ := http.NewRequest("POST", os.Getenv("AUTH_URL"), nil)

	// トークンを追加する
	req.Header.Set("Authorization", token)

	// リクエストを送信する
	client := new(http.Client)
	resp, err := client.Do(req)

	// エラー処理
	if err != nil {
		return UserData{}, err
	}

	defer resp.Body.Close()

	// エラー処理
	if resp.StatusCode != 200 {
		return UserData{}, errors.New(fmt.Sprint("Error: status code", resp.StatusCode))
	}

	var bind_data ReturnResult
	// Struct にバインドする
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &bind_data); err != nil {
		return UserData{}, err
	}

	return bind_data.Result, nil
}
