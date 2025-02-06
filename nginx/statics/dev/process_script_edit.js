document.addEventListener('DOMContentLoaded', function() {
    loadDraft(); // ページが読み込まれたときに下書きを読み込む

    // フォームの送信時に下書きを削除
    document.getElementById('recipeForm').addEventListener('submit',async function(evt) {
        evt.preventDefault();

        // データ保存
        saveDraft();

        //TODO 送信処理を書く
        console.log(localStorage.getItem("recipeDraft"));

        // トークン取得
        const authData = await GetSession();

        // URLからUIDを取得
        const urlParams = new URLSearchParams(window.location.search);
        const removedId = urlParams.get('uid');

        // レシピを削除
        const req = await fetch('/recipe/debugDeleteRecipe', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': authData['token'],
                "recipeId": removedId
            },
        });

        const res = await req.json();
        console.log(res);

        // 登録
        const RegisterReq = await fetch("/recipe/register_recipe",{
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization" : authData["token"],
            },
            body : localStorage.getItem("recipeDraft")
        });

        // 結果を取得
        const result = await RegisterReq.json();

        const payload = new FormData();
        payload.append("image",document.getElementById("recipeImage").files[0]);

        // 画像をアップロード
        const imgReq = await fetch("/recipe/upload_image",{
            method: "POST",
            headers: {
                "uid" : result["result"],
                "Authorization" : authData["token"],
            },
            body: payload
        })

        // 結果を取得
        const imgResult = await imgReq.json();

        console.log(imgResult);

        // リダイレクト
        window.location.href = "./process_edit.html?uid=" + result["result"];

        // 下書きを削除
        // localStorage.removeItem('recipeDraft');
    });
});

function addStep() {
    const stepContainer = document.createElement('div');
    stepContainer.className = 'step';

    stepContainer.innerHTML = `
        <label for="stepName">手順名:</label>
        <input type="text" name="stepName" required>
        
        <label for="stepTime">時間 (分):</label>
        <input type="number" name="stepTime" required>
        
        <label for="stepType">タスクの種類:</label>
        <select name="stepType" required>
            <option value="下準備">下準備</option>
            <option value="調理">調理</option>
            <option value="仕上げ">仕上げ</option>
        </select>

        <label for="stepConcurrent">並行可・不可:</label>
        <select name="stepConcurrent" required>
            <option value="可">並行可</option>
            <option value="不可">並行不可</option>
        </select>

        <label for="stopDescription">手順の説明</label>
        <textarea name="stopDescription" required></textarea>
        
        <div class="ingredientsContainer">
            <h3>手順で使用する材料</h3>
            <button type="button" class="add-button" onclick="addIngredient(this)">+ 材料を追加</button>
        </div>

        <div class="utensilsContainer">
            <h3>手順で使用する器具</h3>
            <button type="button" class="add-button" onclick="addUtensil(this)">+ 器具を追加</button>
        </div>

        <button type="button" class="remove-button" onclick="removeStep(this)">手順を削除</button>
    `;

    const stepsList = document.getElementById('stepsList');
    stepsList.appendChild(stepContainer);
}


function addIngredient(button) {
    const ingredientContainer = document.createElement('div');
    ingredientContainer.className = 'ingredient';

    ingredientContainer.innerHTML = `
        <div class="ingredient-row">
            <label for="ingredientName">材料名:</label>
            <input type="text" name="ingredientName" required>
            
            <label for="ingredientQuantity">個数:</label>
            <input type="number" step="0.1" name="ingredientQuantity" required>
            
            <label for="ingredientUnit">単位:</label>
            <input type="text" name="ingredientUnit" required>
        </div>
        
        <button type="button" class="remove-button" onclick="removeIngredient(this)">材料を削除</button>
    `;

    button.parentElement.appendChild(ingredientContainer);
}

function addUtensil(button) {
    const utensilContainer = document.createElement('div');
    utensilContainer.className = 'utensil';

    utensilContainer.innerHTML = `
        <div class="utensil-row">
            <label for="utensilName">器具名:</label>
            <input type="text" name="utensilName" required>
            
            <label for="utensilQuantity">個数:</label>
            <input type="number" step="0.1" name="utensilQuantity" required>
            
            <label for="utensilUnit">単位:</label>
            <input type="text" name="utensilUnit" required>
        </div>
        
        <button type="button" class="remove-button" onclick="removeUtensil(this)">器具を削除</button>
    `;

    button.parentElement.appendChild(utensilContainer);
}

function removeStep(button) {
    button.parentElement.remove();
}

function removeIngredient(button) {
    button.parentElement.remove();
}

function removeUtensil(button) {
    button.parentElement.remove();
}

function saveDraft() {
    const recipeCategory = document.getElementById('recipeCategory').value; // 料理の項目を取得
    const recipeName = document.getElementById('recipeName').value;
    const finalState = document.getElementById('finalState').value; // 最終状態を取得
    const steps = Array.from(document.querySelectorAll('.step')).map(step => {
        return {
            name: step.querySelector('input[name="stepName"]').value,
            time: Number(step.querySelector('input[name="stepTime"]').value), // 数値型に変換
            type: step.querySelector('select[name="stepType"]').value, // タスクの種類を取得
            concurrent: step.querySelector('select[name="stepConcurrent"]').value,
            ingredients: Array.from(step.querySelectorAll('.ingredient')).map(ingredient => {
                return {
                    name: ingredient.querySelector('input[name="ingredientName"]').value,
                    quantity: Number(ingredient.querySelector('input[name="ingredientQuantity"]').value), // 数値型に変換
                    unit: ingredient.querySelector('input[name="ingredientUnit"]').value
                };
            }),
            utensils: Array.from(step.querySelectorAll('.utensil')).map(utensil => {
                return {
                    name: utensil.querySelector('input[name="utensilName"]').value,
                    quantity: Number(utensil.querySelector('input[name="utensilQuantity"]').value), // 数値型に変換
                    unit: utensil.querySelector('input[name="utensilUnit"]').value
                };
            }),
            description: step.querySelector('textarea[name="stopDescription"]').value
        };
    });

    const draft = { recipeCategory, recipeName, finalState, steps }; // 最終状態を含める
    localStorage.setItem('recipeDraft', JSON.stringify(draft)); // 下書きを保存
}

function loadDraft() {
    const draft = JSON.parse(localStorage.getItem('recipeDraft'));
    if (draft) {
        document.getElementById('recipeCategory').value = draft.recipeCategory || ""; // 料理の項目
        document.getElementById('recipeName').value = draft.recipeName || "";

        // 最終状態の読み込み
        document.getElementById('finalState').value = draft.finalState || ""; // 最終状態

        draft.steps.forEach(step => {
            addStep();
            const lastStep = document.querySelector('.step:last-child');
            lastStep.querySelector('input[name="stepName"]').value = step.name;
            lastStep.querySelector('input[name="stepTime"]').value = step.time;
            lastStep.querySelector('select[name="stepType"]').value = step.type; // タスクの種類を設定
            lastStep.querySelector('select[name="stepConcurrent"]').value = step.concurrent;
            lastStep.querySelector('textarea[name="stopDescription"]').value = step.description;

            step.ingredients.forEach(ingredient => {
                addIngredient(lastStep.querySelector('.ingredientsContainer .add-button'));
                const lastIngredient = lastStep.querySelector('.ingredient:last-child');
                lastIngredient.querySelector('input[name="ingredientName"]').value = ingredient.name;
                lastIngredient.querySelector('input[name="ingredientQuantity"]').value = ingredient.quantity;
                lastIngredient.querySelector('input[name="ingredientUnit"]').value = ingredient.unit;
            });

            step.utensils.forEach(utensil => {
                addUtensil(lastStep.querySelector('.utensilsContainer .add-button'));
                const lastUtensil = lastStep.querySelector('.utensil:last-child');
                lastUtensil.querySelector('input[name="utensilName"]').value = utensil.name;
                lastUtensil.querySelector('input[name="utensilQuantity"]').value = utensil.quantity;
                lastUtensil.querySelector('input[name="utensilUnit"]').value = utensil.unit;
            });
        });
    }
}

async function Init() {
    try {
        // 認証情報取得
        const authData = await GetSession();

        console.log(authData["record"]);
    } catch (ex) {
        console.error(ex);

        // 認証していない場合ログインに飛ばす
        window.location.href = LoginURL;
    }
}

// 初期化
Init();

const extractButton = document.getElementById('extractButton');
extractButton.addEventListener('click', async () => {
    try {
        extractButton.textContent = "処理中";

        const authData = await GetSession();

        console.log(authData["token"]);

        const req = await fetch("/recipe/extract",{
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization" : authData["token"],
            },
            body : JSON.stringify({
                "url" : document.getElementById("extractURL").value
            })
        });

        const res = await req.json();
        localStorage.setItem('recipeDraft', res["result"]); // 下書きを保存 (res["result"]);
        // reload
        location.reload();
    } catch (ex) {
        console.error(ex);
        extractButton.textContent = "抽出";
        alert("失敗しました");
    }
});