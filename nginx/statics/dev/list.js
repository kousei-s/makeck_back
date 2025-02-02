// サンプルデータを作成
let recipes = [
    { id: 1, name: "カレー", type: "主食", image: "https://example.com/curry.jpg", status: "完成" },
    { id: 2, name: "サラダ", type: "副菜", image: "https://example.com/salad.jpg", status: "完成" }
];

// 料理リストを表示する関数
async function displayRecipes() {
    const authData = await GetSession();

    const req = await fetch('/recipe/debugRecipes', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': authData['token']
        },
    });

    const res = await req.json();
    console.log(res);

    recipes = res["recipes"];

    const recipeList = document.getElementById('recipe-list');
    recipeList.innerHTML = ''; // リストをクリア

    recipes.forEach(recipe => {
        const li = document.createElement('li');

        // IDを表示
        const idElement = document.createElement('span');
        idElement.textContent = `ID: ${recipe.id}`;
        idElement.style.marginRight = '10px';
        idElement.style.display = "none";

        const title = document.createElement('h3');
        title.textContent = `${recipe.name} (${recipe.type})`;

        const img = document.createElement('img');
        img.src = recipe.image;
        img.alt = recipe.name;
        img.style.maxWidth = '100px';

        const status = document.createElement('p');
        status.textContent = `最終状態: ${recipe.status}`;

        const deleteButton = document.createElement('button');
        deleteButton.className = 'delete';
        deleteButton.textContent = '削除';

        // 削除ボタンのイベントリスナーを追加
        deleteButton.addEventListener('click', async function () {
            const index = recipes.indexOf(recipe);
            if (index > -1) {
                const removedId = recipe.id; // 削除するレシピのIDを取得
                recipes.splice(index, 1); // レシピを削除

                // レシピを削除
                const authData = await GetSession();
                const req = await fetch('/recipe/debugDeleteRecipe', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': authData['token'],
                        "recipeId": removedId
                    },
                    // body: JSON.stringify({ id: removedId }) // 削除するレシピのIDを送信
                });

                const res = await req.json();
                console.log(res);

                await displayRecipes(); // リストを再表示
            }
        });

        li.appendChild(idElement);
        li.appendChild(title);
        li.appendChild(img);
        li.appendChild(status);
        li.appendChild(deleteButton);

        recipeList.appendChild(li);
    });
}

// 初期表示
displayRecipes();
