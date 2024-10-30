async function Init() {
    try {
        // 認証情報取得
        const authData = await GetSession();
    } catch (ex) {
        console.error(ex);
    }
}

// 初期化
Init();