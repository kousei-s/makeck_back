import os
import sys

# 作業ディレクトリを変える
os.chdir(os.path.dirname(os.path.dirname(sys.argv[0])))

import subprocess

# 各種設定
compose_yaml = "./composes/docker-compose_release.yaml"
project_name = "atkit_release"

def runCommand(command):
    print("--------------------------------------------------")
    print(f"{command} を実行します")

    # コマンドを実行
    subprocess.run(command,shell=True,cwd=os.getcwd())

# コンテナを落とす
runCommand(f"docker compose -f {compose_yaml} -p {project_name} down")

# コンテナを起動する
runCommand(f"docker compose -f {compose_yaml} -p {project_name} up -d --build")