import subprocess
from shutil import copytree, copy2, rmtree
import os
import sys
import json

# 作業フォルダ変更
os.chdir(os.path.dirname(sys.argv[0]))

config = {
    "containers_dir": "../composes/containers",
    "nginx_conf_dir": "../nginx/conf.d/include",
    "template_dir": "../template",
    "projects": [],
}


def copy2_verbose(src, dst):
    print('Copying {0}'.format(src))
    copy2(src, dst)


def CheckCreateCompose(project_name):
    global config

    container_dir = config['containers_dir']
    # ファイルが存在するか
    if os.path.exists(f"{container_dir}/{project_name}.yaml"):
        print("すでにファイルが存在します")
        # 存在する時
        sys.exit(1)


def CreateCompose(project_name, hostname):
    global config

    # チェックする
    CheckCreateCompose(project_name)

    container_dir = config['containers_dir']
    # ファイルを生成する
    with open(f"{container_dir}/{project_name}.yaml", "w", encoding="utf-8") as writeCompose:
        writeCompose.write(f"""
services:
    {project_name}:
        build: ../{project_name}
        hostname: {hostname}
        restart: always
        tty: true
        volumes:
            - ../{project_name}/src:/root/{project_name}
""")


def CheckNginxSetting(hostname):
    global config

    # ファイルが存在するか
    nginx_conf_dir = config['nginx_conf_dir']
    if os.path.exists(f"{nginx_conf_dir}/{hostname}.conf"):
        print("すでにファイルが存在します")
        # 存在する時
        sys.exit(1)


def CreateNginxSetting(hostname):
    global config

    # チェックする
    CheckNginxSetting(hostname)

    nginx_conf_dir = config['nginx_conf_dir']
    # ファイルに書き込む
    with open(f"{nginx_conf_dir}/{hostname}.conf", "w", encoding="utf-8") as writeNginx:
        writeNginx.write("""
location /REPLACEPENDPOINT/ {
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;

    proxy_pass http://REPLACEHOSTNAME:REPLACEPORT/;   
}
""".replace("REPLACEPENDPOINT", hostname).replace("REPLACEHOSTNAME", hostname).replace("REPLACEPORT", "8090"))

# プロジェクトが存在するか


def checkExistsProject(project_name, hostname):
    global config

    # 既存のプロジェクトと被らないかを判定
    for pdata in config["projects"]:
        # プロジェクトが存在する場合
        if (pdata["project_name"] == project_name):
            # すでにプロジェクトが存在する場合
            print("プロジェクトがすでに存在します")
            return True

        # ホスト名が存在する場合
        if (pdata["hostname"] == hostname):
            # すでにプロジェクトが存在する場合
            print("プロジェクトがすでに存在します")
            return True

    return False


def CreateInput():
    while True:
        print("プロジェクトの設定 (ctrl+c でキャンセル)")
        project_name = input("プロジェクト名を入力してください>")
        hostname = input("ホスト名を入力してください>")

        # すでに存在するか判定
        if (checkExistsProject(project_name, hostname)):
            continue

        break

    return project_name, hostname


# 設定ファイルのパス
setting_path = "./utils.json"

# 設定ファイルを読みこむ関数


def LoadConfig():
    global config

    # 設定ファイルが存在しないとき作成する
    if os.path.exists(setting_path):
        # 存在する時読みこむ
        with open(setting_path, "r", encoding="utf-8") as readConfig:
            config = json.load(readConfig)

    else:
        # 存在しない時
        # 作成する
        with open(setting_path, "w", encoding="utf-8") as writeConfig:
            json.dump(config, writeConfig)


def WriteConfig():
    global config

    # 作成する
    with open(setting_path, "w", encoding="utf-8") as writeConfig:
        json.dump(config, writeConfig)


# 設定ファイルを読みこむ
LoadConfig()

# 書き込む
WriteConfig()

# コンテナフォルダを作成する
try:
    os.makedirs(config["containers_dir"])
except:
    pass


def RefreshProject():
    subprocess.run([sys.executable, "./develop.py"], cwd=os.getcwd())


def CreateProject(hostname, project_name):
    # nginx の設定を生成
    CreateNginxSetting(hostname)
    CreateCompose(project_name, hostname)

    # 出力先
    dst_dir = f"../{project_name}"

    # テンプレートからコピーする
    copytree(config["template_dir"], dst_dir, copy_function=copy2_verbose)

    # セットアップファイルを実行する
    subprocess.run([sys.executable, "setup.py", project_name,"0.0.0.0:8090", f"/root/{project_name}"], cwd=dst_dir)

    container_dir = config['containers_dir']
    nginx_conf_dir = config['nginx_conf_dir']
    # リストに追加する
    config["projects"].append({
        "project_name": project_name,
        "hostname": hostname,
        "compose_path": f"{container_dir}/{project_name}.yaml",
        "conf_path": f"{nginx_conf_dir}/{hostname}.conf",
        "setup_script" : f"{dst_dir}/setup.py",
        "project_dir": dst_dir
    })

    # 設定を書き込む
    WriteConfig()

    # プロジェクトをリフレッシュ
    RefreshProject()


def DeleteProject():
    global config

    while True:
        print("削除するプロジェクト番号を入力してください (ctrl+c でキャンセル)")
        for index, pdata in enumerate(config["projects"]):
            project_name = pdata["project_name"]
            hostname = pdata["hostname"]
            print(f"プロジェクト番号: {index}  名前: {project_name}   ホスト名: {hostname}")

        # 入力を受け取る
        input_data = input(">")

        try:
            # 削除する
            pdata = config["projects"].pop(int(input_data))
            
            # プロジェクトを落とす
            subprocess.run([sys.executable, "./down_develop.py"])

            # docker のファイルを削除する
            print("docker ファイルを削除しています")
            os.remove(pdata["compose_path"])

            # nginx の設定ファイルを削除する
            print("nginx の設定ファイルを削除しています")
            os.remove(pdata["conf_path"])

            # フォルダを削除する
            print("プロジェクトフォルダを削除しています")
            rmtree(pdata["project_dir"])

            break
        except:
            print("削除に失敗しました")

    # 設定ファイルに書き込む
    WriteConfig()

    # プロジェクトをリフレッシュ
    RefreshProject()

def RegenProject():
    global config

    print("再生成します")

    for index, pdata in enumerate(config["projects"]):
        project_name = pdata["project_name"]
        print(f"{project_name} を再生成しています")
        # プロジェクトを再生成しています
        try:
            subprocess.run([sys.executable,pdata["setup_script"],"regen"])
        except:
            import traceback
            traceback.print_exc()

    print("再生成完了")
def ListInput(inputdict: dict):
    while True:
        # 表示するリスト
        show_list = []

        print("項目を選んでください")
        # 辞書を回す
        for ikey in inputdict.keys():
            print(f"{ikey}:{inputdict[ikey]}")
            show_list.append(str(inputdict[ikey]).lower())

        # 入力を受け取る
        show_str = ",".join(show_list)
        input_tag = input(f"項目を入力してください [{show_str}]>")

        if (input_tag.lower() in show_list):
            # 項目のバリデーションが成功した場合
            return input_tag.lower()


while True:
    # 入力を受け取る
    input_mode = ListInput({
        "プロジェクトを作成する": "c",
        "プロジェクトを削除する": "d",
        "再生成": "r",
        "キャンセル": "q"
    })

    # 項目ごとに変更
    if input_mode == "c":
        # 設定
        project_name, hostname = CreateInput()

        # # プロジェクトを生成
        CreateProject(hostname, project_name)

    elif input_mode == "d":
        # プロジェクトを削除する
        DeleteProject()
    elif input_mode == "r":
        # env等を再生成する
        RegenProject()
    elif input_mode == "q":
        break
