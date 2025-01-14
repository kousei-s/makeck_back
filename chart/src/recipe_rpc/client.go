package recipe_rpc

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func InitClient() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("recipe:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	// GRPC クライアント
	client := NewRecipeServiceClient(conn)

	response, err := client.GetRecipe(context.Background(), &RecipeRequest{
		Uid: "098807d7-04ef-4f03-b1e9-e5792da2aeb1",
	})

	// エラー処理
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Print(response)
	for _,val  := range response.Process {
		log.Println(val)
		log.Println(val.Material)
	}

	defer conn.Close()
}
