import os
import sys
import secrets
import string

def get_random_password_string(length):
    pass_chars = string.ascii_letters + string.digits
    password = ''.join(secrets.choice(pass_chars) for x in range(length))
    return password

def replaceFile(src,dst,replaces):
    with open(src,"r",encoding="utf-8") as rdata:
        readdata = rdata.read()
    
    # 辞書を回す
    for rdata in replaces.keys():
        # 置き換え
        readdata = readdata.replace(rdata,replaces[rdata])

    # 書き込み
    with open(dst,"w",encoding="utf-8") as wdata:
        wdata.write(readdata)

def Setup(projectName,bindAddr,workdir):
    router_file = "./src/router.go"
    template_router_file = "./template/router.go"

    init_file = "./src/init.go"
    template_init_file = "./template/init.go"

    mod_file = "./src/go.mod"
    template_mod_file = "./template/go.mod"

    env_path = "./src/.env"

    env_data = f"""AUTH_URL = "http://pocketbasec:8080/jwt"
BindAddr = "{bindAddr}"

DBPATH = "./{projectName}.db"

SECRETKEY = "{get_random_password_string(64)}"
"""
    # env がある時
    if os.path.exists(env_path):
        print("すでに設定済みです")
        return

    print("env 書き込み")
    # env を書き込む
    with open(env_path,"w",encoding="utf-8") as wenv:
        wenv.write(env_data)
    
    print("router.go 作成")
    # Router.go 作成
    replaceFile(template_router_file,router_file,{
        "template/middlewares" : f"{projectName}/middlewares",
        "template/controllers" : f"{projectName}/controllers"
    })

    print("init.go 作成")
    # init.go 作成
    replaceFile(template_init_file,init_file,{
        "template/controllers" : f"{projectName}/controllers",
    })

    print("go.mod 作成")
    # go.mod 作成
    replaceFile(template_mod_file,mod_file,{
        "module template" : f"module {projectName}"
    })

    # Dockerfile 読み込み
    with open("./template/dockerfile","r",encoding="utf-8") as readDocker:
        rdocker = readDocker.read()
    
    # 書き込み
    with open("./dockerfile","w",encoding="utf-8") as writeDocker:
        writeDocker.write(rdocker)
        writeDocker.write(f"WORKDIR {workdir}")

if __name__ == "__main__":
    # フォルダを移動
    os.chdir(os.path.dirname(sys.argv[0]))

    Setup(sys.argv[1],sys.argv[2],sys.argv[3])