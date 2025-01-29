package recipe_rpc

import (
	"chart/utils"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	client RecipeServiceClient
)

func InitClient() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("recipe:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	// GRPC クライアント
	client = NewRecipeServiceClient(conn)

	utils.Println("recipe rpc client initialized")
}

func GetRecipe(uid string) (*Recipe, error) {
	response, err := client.GetRecipe(context.Background(), &RecipeRequest{
		Uid: uid,
	})

	// エラー処理
	if err != nil {
		utils.Println(err)
		return nil, err
	}

	return response, nil
}
