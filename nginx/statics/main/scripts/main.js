async function Init() {
    try {
        // 認証情報取得
        const authData = await GetSession();

        console.log(authData["record"]);

        // ユーザーデータ取得
        const UserData = authData["record"];

        // 画像設定
        document.getElementById("usericon").src = GetIcon(UserData["id"]);
    } catch (ex) {
        console.error(ex);

        // 認証していない場合ログインに飛ばす
        window.location.href = LoginURL;
    }
}

// 初期化
Init();

async function DoLogout() {
    try {
        await Logout(true);
    } catch (ex) {
        console.error(ex);

        alert("ログアウトに失敗しました");
    }
}

async function TestSearch() {
    const req = await fetch("/recipe/search",{
        method: "POST",
        headers: {
            "Content-Type" : "application/json"
        },
        body: JSON.stringify({
            "name" : "",
            "category" : "主菜"
        })
    })

    console.log(await req.json());
}