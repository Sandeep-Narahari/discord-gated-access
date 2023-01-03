package bot

import (
	// "github.com/forbole/juno/v3/database/config"

	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/AutonomyNetwork/iam/database"
	"google.golang.org/grpc"

	botpb "github.com/AutonomyNetwork/iam/types/bot/v1/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/forbole/juno/v3/modules"

	"github.com/forbole/juno/v3/types/config"
)

const (
	ModuleName = "bot"
)

var (
	_ modules.Module = &Module{}
	// _ modules.AdditionalOperationsModule = &Module{}
)

type Module struct {
	db      *database.Db
	cfg     *Config
	session *discordgo.Session
	grpc    *grpc.Server
}

func NewModule(c config.Config, db *database.Db) *Module {
	s := make(chan *discordgo.Session)
	grpc_in := make(chan *grpc.Server)

	bz, err := c.GetBytes()
	if err != nil {
		panic(err)
	}
	botCfg, err := ParseConfig(bz)
	if err != nil {
		panic(err.Error())
	}
	go connectDiscord(botCfg.TokenId, s)
	go connectServer(botCfg.Tcp_port, grpc_in)

	session := <-s
	grpc_session := <-grpc_in

	return &Module{
		cfg:     botCfg,
		db:      db,
		session: session,
		grpc:    grpc_session,
	}

}

func (m *Module) Name() string {
	// fmt.Println("TESTING BOT ID", m.cfg.TokenId)
	return ModuleName
}

func connectDiscord(tokenId string, ch chan *discordgo.Session) {

	dg, err := discordgo.New("Bot " + tokenId)
	if err != nil {
		fmt.Println("error creating Discord session,", err)

	}

	fmt.Println("Bot is now running.")

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)

	}
	ch <- dg
	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	fmt.Println("hello")
	// Cleanly close down the Discord session.
	dg.Close()

}

func connectServer(tcp_port string, grpc_ch chan *grpc.Server) {
	fmt.Println(tcp_port)

	// create new gRPC server
	server := grpc.NewServer()

	// register the CommunityServicesServerImpl and User Service on the gRPC server
	botpb.RegisterCommunityServicesServer(server, &Server{})
	botpb.RegisterUserServiceServer(server, &NewServer().UnimplementedUserServiceServer)

	// start listening on port :8080 for a tcp connection
	l, err := net.Listen("tcp", tcp_port)
	if err != nil {
		log.Fatal("error in listening on port :", tcp_port, err)
	}
	// the gRPC server
	if err := server.Serve(l); err != nil {
		log.Fatal("unable to start server", err)
	}
	fmt.Println("CONNECTED SERVER")
	fmt.Println("SERVER CONNECTED", server)
	fmt.Println("LISTNER", l)

	grpc_ch <- server
	defer l.Close()
}
