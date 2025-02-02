package recipe_rpc

import (
	"context"
	"log"
	"net"
	"recipe/models"
	"recipe/utils"

	"google.golang.org/grpc"
)

func RunServer() {
	log.Print("main start")

	// 9000番ポートでクライアントからのリクエストを受け付けるようにする
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// gRPCサーバーを起動
	RegisterRecipeServiceServer(grpcServer, &RecipeServer{})

	// 以下でリッスンし続ける
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Print("main end")
}

type RecipeServer struct{}

// GetRecipe implements RecipeServiceServer.
func (rserver *RecipeServer) GetRecipe(ctx context.Context, req *RecipeRequest) (*Recipe, error) {
	log.Print("GetRecipe")

	// レシピ取得
	recipe, err := models.GetRecipe(req.Uid)

	// エラー処理
	if err != nil {
		return nil, err
	}

	processes := []*Process{}
	for _, val := range recipe.Process {
		process := recipeProcessToProcess(val)
		processes = append(processes, &process)
		// utils.Println(process.Tools)
		utils.Println(process.Material)
	}

	return &Recipe{
		Uid:       recipe.Uid,
		Name:      recipe.Name,
		Image:     recipe.Image,
		Process:   processes,
		LastState: recipeLastStateToLastState(recipe.LastState),
	}, nil
}

// model の LastState を gRPC の LastState に変換
func recipeLastStateToLastState(lastState models.LastSatate) LastState {
	if lastState == models.Hot {
		return LastState_HOT
	} else if lastState == models.Cool {
		return LastState_COOL
	} else {
		return LastState_NORMAL
	}
}

// model の material を gRPC の material に変換
func recipeMaterialToMaterial(material models.Material) Material {
	return Material{
		Uid:   material.Uid,
		Name:  material.Name,
		Count: int32(material.Count),
		Unit:  material.Unit,
	}
}

// model の process を gRPC の process に変換
func recipeProcessToProcess(process models.Process) Process {
	// tool を変換
	tools := []*Tools{}

	// toolsを取得
	getTools, err := process.GetTools()
	if err != nil {
		utils.Println(err)
	}

	// 変換
	for _, val := range getTools {
		converted := toolToToools(val)
		tools = append(tools, &converted)
	}

	// material を変換
	materials := []*Material{}

	// materialを取得
	getMaterials, err := process.GetMaterials()
	if err != nil {
		utils.Println(err)
	}

	// 変換
	for _, val := range getMaterials {
		converted := recipeMaterialToMaterial(val)
		materials = append(materials, &converted)
	}

	return Process{
		Uid:         process.Uid,
		Name:        process.Name,
		Description: process.Description,
		Parallel:    process.Parallel,
		Time:        int32(process.Time),
		Tools:       tools,
		Material:    materials,
		Recipeid:    process.Recipeid,
		Type:        ProcessType(process.Type),
	}
}

// model の tools を gRPC の tools に変換
func toolToToools(tool models.Tools) Tools {
	return Tools{
		Uid:   tool.Uid,
		Name:  tool.Name,
		Count: int32(tool.Count),
		Unit:  tool.Unit,
	}
}
