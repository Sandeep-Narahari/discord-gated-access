package modules

import (
	"context"
	"fmt"
	"log"
	"time"

	botpb "github.com/AutonomyNetwork/iam/types/bot/v1/bot"
	"github.com/forbole/juno/v3/modules"
	"github.com/forbole/juno/v3/modules/registrar"
	"github.com/forbole/juno/v3/node/builder"
	"github.com/forbole/juno/v3/node/remote"
	"google.golang.org/grpc"

	internaldb "github.com/AutonomyNetwork/iam/database"
	"github.com/AutonomyNetwork/iam/modules/bot"
	"github.com/AutonomyNetwork/iam/modules/marketplace"
)

type Server struct {
	botpb.UnimplementedCommunityServicesServer
	botpb.UnimplementedUserServiceServer
}

func NewServer() *Server {
	return &Server{}
}

type ModuleRegistar struct {
}

// Context represents the context of the modules registrar
// type Context struct {
// 	BotConfig bot.Config
// }

func NewModuleRegistrar() *ModuleRegistar {
	return &ModuleRegistar{}
}

func (r *ModuleRegistar) BuildModules(ctx registrar.Context) modules.Modules {
	cdc := ctx.EncodingConfig.Marshaler
	ojunoDb := internaldb.Cast(ctx.Database)

	remoteCfg, ok := ctx.JunoConfig.Node.Details.(*remote.Details)
	if !ok {
		panic(fmt.Errorf("cannot run OJuno on local node"))
	}

	node, err := builder.BuildNode(ctx.JunoConfig.Node, ctx.EncodingConfig)
	if err != nil {
		panic(fmt.Errorf("cannot build node:%s", err))
	}

	BotConnection := bot.NewModule(ctx.JunoConfig, ojunoDb)

	grpcConnection := remote.MustCreateGrpcConnection(remoteCfg.GRPC)

	// fmt.Println("BOT CONECTINON====", BotConnection, grpcConnection)
	onftModule := marketplace.NewModule(node, grpcConnection, cdc, ojunoDb)

	// Handle Discord Session
	// BotConnection.ConnectDiscord()
	// go connectGrpc()

	// Calling Channels
	BotConnection.Channels()
	BotConnection.AddHandleMsgCreate()
	BotConnection.GetAddressAndValidate()
	fmt.Println("connected")

	go client()
	fmt.Println("connected")

	return []modules.Module{
		onftModule,
		BotConnection,
	}
}

// func connectGrpc() {
// 	// create new gRPC server
// 	server := grpc.NewServer()
// 	// register the GreeterServerImpl on the gRPC server
// 	botpb.RegisterCommunityServicesServer(server, &Server{})
// 	botpb.RegisterUserServiceServer(server, &NewServer().UnimplementedUserServiceServer)
// 	// start listening on port :8080 for a tcp connection
// 	l, err := net.Listen("tcp", ":5080")
// 	if err != nil {
// 		log.Fatal("error in listening on port :5080", err)
// 	}
// 	// the gRPC server
// 	if err := server.Serve(l); err != nil {
// 		log.Fatal("unable to start server", err)
// 	}
// 	fmt.Println("CONNECTED SERVER")
// 	fmt.Println("SERVER CONNECTED", server)
// }

func client() {
	conn, err := grpc.Dial("localhost:7080", grpc.WithInsecure())
	fmt.Println(conn)
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
	fmt.Println(r.GatedCollectionId)
	fmt.Println("CONNECTED CLIENT", conn)
}
func (s *Server) CreateCommunity(ctx context.Context, r *botpb.CreateCommunityRequest) (*botpb.CreateCommunityResponse, error) {
	fmt.Println("CREATE COMMUNITY")
	fmt.Println(&botpb.CreateCommunityResponse{GatedCollectionId: r.GatedCollectionId})
	return &botpb.CreateCommunityResponse{GatedCollectionId: r.GatedCollectionId}, nil
}
