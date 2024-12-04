document.addEventListener('DOMContentLoaded', function() {
    loadDraft(); // ページが読み込まれたときに下書きを読み込む

    // フォームの送信時に下書きを削除
    document.getElementById('recipeForm').addEventListener('submit', function(evt) {
        evt.preventDefault();

        //TODO 送信処理を書く

        // 下書きを削除
        localStorage.removeItem('recipeDraft');
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
            <input type="number" name="ingredientQuantity" required>
            
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
            <input type="number" name="utensilQuantity" required>
            
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
    const recipeName = document.getElementById('recipeName').value;
    const recipeImage = document.getElementById('recipeImage').files[0] ? document.getElementById('recipeImage').files[0].name : "";
    const steps = Array.from(document.querySelectorAll('.step')).map(step => {
        return {
            name: step.querySelector('input[name="stepName"]').value,
            time: step.querySelector('input[name="stepTime"]').value,
            concurrent: step.querySelector('select[name="stepConcurrent"]').value,
            ingredients: Array.from(step.querySelectorAll('.ingredient')).map(ingredient => {
                return {
                    name: ingredient.querySelector('input[name="ingredientName"]').value,
                    quantity: ingredient.querySelector('input[name="ingredientQuantity"]').value,
                    unit: ingredient.querySelector('input[name="ingredientUnit"]').value
                };
            }),
            utensils: Array.from(step.querySelectorAll('.utensil')).map(utensil => {
                return {
                    name: utensil.querySelector('input[name="utensilName"]').value,
                    quantity: utensil.querySelector('input[name="utensilQuantity"]').value,
                    unit: utensil.querySelector('input[name="utensilUnit"]').value
                };
            }),
            description: step.querySelector('textarea[name="stopDescription"]').value
        };
    });

    const draft = { recipeName, recipeImage, steps };
    localStorage.setItem('recipeDraft', JSON.stringify(draft)); // 下書きを保存
}

function loadDraft() {
    const draft = JSON.parse(localStorage.getItem('recipeDraft'));
    if (draft) {
        document.getElementById('recipeName').value = draft.recipeName || "";
        
        if (draft.recipeImage) {
            document.getElementById('recipeImage').value = draft.recipeImage;
        }

        draft.steps.forEach(step => {
            addStep();
            const lastStep = document.querySelector('.step:last-child');
            lastStep.querySelector('input[name="stepName"]').value = step.name;
            lastStep.querySelector('input[name="stepTime"]').value = step.time;
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
