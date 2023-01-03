package bot

import (
	"context"
	"log"
	"time"

	botpb "github.com/AutonomyNetwork/iam/types/bot/v1/bot"
	"google.golang.org/grpc"
)

func client() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	client := botpb.NewCommunityServicesClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.CreateCommunity(ctx, &botpb.CreateCommunityRequest{UserDiscordId: "111", DiscordCategoryName: "aaa", AccessRoleName: "abc", AccountAddress: "10101010", GatedCollectionId: "1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf(r.GatedCollectionId)
}
