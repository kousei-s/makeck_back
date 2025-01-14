package recipe_server

import (
	"context"
	"log"
	"net"
	"recipe/models"

	"google.golang.org/grpc"
)

func main() {
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
	}

	return &Recipe{
		Uid:       recipe.Uid,
		Name:      recipe.Name,
		Image:     recipe.Image,
		Process:   processes,
		LastState: recipe.LastState,
	}, nil
}

// model の process を gRPC の process に変換
func recipeProcessToProcess(process models.Process) Process {
	return Process{
		Uid:         process.Uid,
		Name:        process.Name,
		Description: process.Description,
		Parallel:    process.Parallel,
		Time:        int32(process.Time),
		Tools:       []*Tools{},
		Material:    []*Material{},
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
