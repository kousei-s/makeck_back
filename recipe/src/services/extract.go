package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"recipe/utils"
)

func Extract(URL string) (string, error) {
	// Jsonを送信する
	body := map[string]string{
		"Url": URL,
	}

	data, _ := json.Marshal(body)

	// POST リクエストを送信する (body を JSON として送信)
	req, err := http.NewRequest("POST", "http://ai-api:8000/recipe", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	// リクエストヘッダーを設定する
	req.Header.Set("Content-Type", "application/json")

	// リクエストを送信する
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// レスポンスを読み込む
	if resp.StatusCode != 200 {
		return "", err
	}

	// レスポンスを読み込む
	resBody, err := io.ReadAll(resp.Body)

	// エラー処理
	if err != nil {
		return "", err
	}

	// レスポンスを読み込む
	response := string(resBody)

	utils.Println(response)

	return response, nil
}