package bot

import (
	// "github.com/forbole/juno/v3/database/config"

	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AutonomyNetwork/iam/database"

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
}

func NewModule(c config.Config, db *database.Db) *Module {
	s := make(chan *discordgo.Session)
	bz, err := c.GetBytes()
	if err != nil {
		panic(err)
	}
	botCfg, err := ParseConfig(bz)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(s)
	go connectDiscord(botCfg.TokenId, s)

	session := <-s

	return &Module{
		cfg:     botCfg,
		db:      db,
		session: session,
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
