package services

import (

)	


type MaterialArgs struct {
	Recipes []string
}


func SerMaterials(args RecipeArgs) ([]Recipes_list, error) {
	convertedRecipes,err := converter(args)
	if err != nil {
		return []Recipes_list{}, err
	}

	return SearchMaterials(convertedRecipes)
}

