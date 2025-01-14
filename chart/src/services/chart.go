package services

import (
	"chart/recipe_rpc"
	"chart/utils"
)

type RecipeArgs struct {
	Recipes []string
}

func GenChart(args RecipeArgs) (RecipeCollection, error) {
	recipe1, err := recipe_rpc.GetRecipe(args.Recipes[0])

	// エラー処理
	if err != nil {
		return RecipeCollection{}, err
	}

	recipe2, err := recipe_rpc.GetRecipe(args.Recipes[1])

	// エラー処理
	if err != nil {
		return RecipeCollection{}, err
	}

	recipe3, err := recipe_rpc.GetRecipe(args.Recipes[2])

	// エラー処理
	if err != nil {
		return RecipeCollection{}, err
	}

	recipe4, err := recipe_rpc.GetRecipe(args.Recipes[3])

	// エラー処理
	if err != nil {
		return RecipeCollection{}, err
	}

	convertedRecipes := []Recipe{}
	// 1 ~ 4 を変換
	for _, recipe := range []recipe_rpc.Recipe{*recipe1, *recipe2, *recipe3, *recipe4} {
		recipe, err := RpcRecipeToRecipe(recipe)
		if err != nil {
			utils.Println(err)
			return RecipeCollection{}, err
		}
		convertedRecipes = append(convertedRecipes, recipe)
	}

	return chart_Register(convertedRecipes)
}

func RpcRecipeToRecipe(recipe recipe_rpc.Recipe) (Recipe, error) {
	// ツールを変換する
	processess := []Process{}

	for _, process := range recipe.Process {
		process, err := RpcProcessToProcess(*process)
		if err != nil {
			utils.Println(err)
			return Recipe{}, err
		}
		processess = append(processess, process)
	}


	return Recipe{
		Uid:       recipe.Uid,
		Name:      recipe.Name,
		Image:     "",
		Process:   processess,
		LastState: RpcLastSatateToLastSatate(recipe.LastState),
	}, nil
}

func RpcProcessToProcess(process recipe_rpc.Process) (Process, error) {
	return Process{
		Uid:         process.Uid,
		Name:        process.Name,
		Displayname: process.Name,
		Description: process.Description,
		Parallel:    process.Parallel,
		Time:        int(process.Time),
		// Tools:       process.Tools,
		// Material:    process.Material,
		Recipeid: process.Recipeid,
		Type:     RpcTypeToProcessType(process.Type),
	}, nil
}

func RpcTypeToProcessType(processType recipe_rpc.ProcessType) ProcessType {
	if processType == recipe_rpc.ProcessType_COOK {
		return Cook
	}
	if processType == recipe_rpc.ProcessType_PREPARE {
		return Prepare
	}
	if processType == recipe_rpc.ProcessType_FINISH {
		return Finish
	}

	return ProcessType(processType)
}

func RpcLastSatateToLastSatate(lastSatate recipe_rpc.LastState) LastSatate {
	if lastSatate == recipe_rpc.LastState_HOT {
		return Hot
	}
	if lastSatate == recipe_rpc.LastState_COOL {
		return Cool
	}
	if lastSatate == recipe_rpc.LastState_REHEAT {
		return Reheat
	}
	if lastSatate == recipe_rpc.LastState_NORMAL {
		return Normal
	}

	return LastSatate(lastSatate)
}