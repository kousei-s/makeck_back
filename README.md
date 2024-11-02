# sisukai3
シスカイ３のバックエンドリポジトリ

# 必要なもの
- Docker
- Python3
- OpenSSL

# 環境構築
## 鍵の生成
- MAC Linux の場合
    - ./Genkey.sh を実行する 
- Windows の場合
    - ./Genkey.bat を実行する
## Docker コンテナの生成
※ポート 8449 を開けてください
```
docker compose up -d --build
```
を実行する

## Pocketbase の管理画面
URL https://localhost:8449/auth/_

## Pocketbase のアプリケーションURL を変更する
設定画面  https://localhost:8449/auth/_/#/settings

Applocation URL を 
```
https://localhost:8449/auth/
```
(サンプル) に変更する

## 機能を追加する場合
docker-compose.yaml があるディレクトリで
```
python3 utils.py 
```
を実行して質問に答えてください

### もしpyyaml が存在しない場合
```
pip install pyyaml
```
を実行してください