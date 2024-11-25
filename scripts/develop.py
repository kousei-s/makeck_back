import os
import sys
import json

# 作業ディレクトリを変える
os.chdir(os.path.dirname(sys.argv[0]))

import subprocess

# 各種設定
compose_yaml = "../composes/docker-compose_develop.yaml"
project_name = "atkit_develop"

def runCommand(command):
    print("--------------------------------------------------")
    print(f"{command} を実行します")

    # コマンドを実行
    subprocess.run(command,shell=True,cwd=os.getcwd())

# json を読みこむ
with open("./utils.json","r",encoding="utf-8") as readUtils:
    readData = json.load(readUtils)

# コマンド文字列
command_str = ""

# プロジェクトを回す
for pdata in readData["projects"]:
    command_str += f"-f {pdata["compose_path"]} "


# コンテナを落とす
runCommand(f"docker compose -f {compose_yaml} {command_str} -p {project_name} down")


# コンテナを起動する
runCommand(f"docker compose -f {compose_yaml} {command_str} -p {project_name} up -d --build")