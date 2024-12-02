function addStep() {
    const stepContainer = document.createElement('div');
    stepContainer.className = 'step';

    // 手順内容を作成
    stepContainer.innerHTML = `
        <label for="stepName">手順名:</label>
        <input type="text" name="stepName" required>
        
        <label for="stepTime">時間 (分):</label>
        <input type="number" name="stepTime" required>
        
        <label for="stepConcurrent">並行不可:</label>
        <input type="checkbox" name="stepConcurrent">
        
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

    // 手順を表示する場所に手順を追加
    const stepsList = document.getElementById('stepsList');
    stepsList.appendChild(stepContainer);
}

function addIngredient(button) {
    const ingredientContainer = document.createElement('div');
    ingredientContainer.className = 'ingredient';

    ingredientContainer.innerHTML = `
        <label for="ingredientName">材料名:</label>
        <input type="text" name="ingredientName" required>
        
        <label for="ingredientQuantity">個数:</label>
        <input type="number" name="ingredientQuantity" required>
        
        <label for="ingredientUnit">単位:</label>
        <input type="text" name="ingredientUnit" required>
        
        <button type="button" class="remove-button" onclick="removeIngredient(this)">材料を削除</button>
    `;

    button.parentElement.appendChild(ingredientContainer);
}

function addUtensil(button) {
    const utensilContainer = document.createElement('div');
    utensilContainer.className = 'utensil';

    utensilContainer.innerHTML = `
        <label for="utensilName">器具名:</label>
        <input type="text" name="utensilName" required>
        
        <label for="utensilQuantity">個数:</label>
        <input type="number" name="utensilQuantity" required>
        
        <label for="utensilUnit">単位:</label>
        <input type="text" name="utensilUnit" required>
        
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
