import datetime
import yaml
import sys
from shutil import copytree,copy2
import os,subprocess

def copy2_verbose(src, dst):
    print('Copying {0}'.format(src))
    copy2(src,dst)

# 設定
configs = {
    "nginx_path" : "./nginx/default.conf",
    "docker-compose_path" : "./docker-compose.yaml",
    "template" : "./template",
}

# バリデーションinput
def check_input(text):
    # 入力取得
    idata = input(text)

    # 文字列がない時
    if idata == "":
        # 強制終了
        sys.exit(1)

    return idata

class Dumper(yaml.Dumper):
    def increase_indent(self, flow=False, *args, **kwargs):
        return super().increase_indent(flow=flow, indentless=False)

def main():
    # バックアップディレクトリ
    backup_dir = "./backup_yml"

    # フォルダを作成
    try:
        os.makedirs(backup_dir)
    except:
        pass

    # プロジェクト名　
    # project_name = "test"
    project_name = check_input("プロジェクト名>")

    # コピー先ディレクトリ
    copy_dir = check_input("フォルダ名>")
    # copy_dir = "test3"

    # ホスト名
    # hostname = "test3"
    hostname = check_input("ホスト名入れてね>")

    # yaml 読み込み
    with open(configs["docker-compose_path"]) as file:
        yml = yaml.safe_load(file)

    # サービス一覧
    services = yml["services"]

    # すでに存在するかの確認
    for pname in services.keys():
        if pname == project_name:
            print("同じ名前のプロジェクトが存在します")
            return
        
        # ホスト名が存在するか
        if services[pname]["hostname"] == hostname:
            print("同じホスト名が存在します")
            return
    
    # コピー先ディレクトリが存在するか
    if os.path.exists(copy_dir):
        print("コピー先が存在します")
        return
    
    print("compose yaml をバックアップしています")
    # バックアップ作成
    dt_now = datetime.datetime.now()
    now_time = dt_now.strftime('%Y-%m-%d_%H-%M-%S')
    with open(os.path.join(backup_dir,now_time + ".back.yml"),"w",encoding="utf-8") as backyml:
        yaml.dump(yml,backyml,allow_unicode=True,indent=4,Dumper=Dumper)

    # yaml に追加
    yml["services"][project_name] = {
        "restart" : "always",
        "container_name" : project_name,
        "build" : f"./{copy_dir}",
        "volumes" : [
            f"./{copy_dir}/src:/root/{project_name}"
        ],
        "tty" : True,
        "restart" : "always",
        "hostname" : hostname
    }

    # フォルダコピー
    copytree(configs["template"],copy_dir, copy_function=copy2_verbose)

    print("セットアップ実行")
    # setup.py 実行
    subprocess.run([sys.executable,f"{copy_dir}/setup.py",project_name,"0.0.0.0:8090",f"/root/{project_name}"])

    # yaml に追加
    with open(configs["docker-compose_path"],"w",encoding="utf-8") as writeyml:
        yaml.dump(yml,writeyml,allow_unicode=True,indent=4,Dumper=Dumper)

    print("Docker 実行")
    # docker 実行
    subprocess.run(["docker","compose","up","--build","-d",project_name])

    print("go mod tidy 実行")
    # go mod tidy 実行
    subprocess.run(["docker","compose","exec",project_name,"go","mod","tidy"])

    # default.conf 読み込み
    with open("./nginx/default.conf","r",encoding="utf-8") as rconf:
        nginx_conf = rconf.read()
    
    print("default.conf をバックアップしています")
    # バックアップ 書き込み
    with open(os.path.join(backup_dir,now_time + ".back.conf"),"w",encoding="utf-8") as backconf:
        backconf.write(nginx_conf)
    
    # 追加するデータ
    adddata = """
    location /REPLACEPENDPOINT/ {
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        proxy_pass http://REPLACEHOSTNAME:REPLACEPORT/;   
    }

    # ADDHERE
""".replace("REPLACEPENDPOINT",project_name).replace("REPLACEHOSTNAME",hostname).replace("REPLACEPORT","8090")
    
    # 置き換え
    replaced = nginx_conf.replace("# ADDHERE",adddata)

    print("設定を書き込んでいます")
    # default.conf に追記
    with open("./nginx/default.conf","w",encoding="utf-8") as wconf:
        wconf.write(replaced)
    
    # nginx を再起動
    subprocess.run(["docker","compose","restart","nginx"])
    
    print("セッツアップ完了")
if __name__ == "__main__":
    main()