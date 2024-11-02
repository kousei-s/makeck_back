# sisukai3
シスカイ３のバックエンドリポジトリ

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