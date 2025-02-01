package services

import (

)

type RecipeArgs struct {
	Recipes []string
}

func GenChart(args RecipeArgs) (RecipeCollection, error) {
	
	convertedRecipes,err :=converter(args)
	if err != nil {
		return RecipeCollection{}, err
	}
	
	return chart_Register(convertedRecipes)
}

