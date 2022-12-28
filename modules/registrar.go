package modules

import (
	"fmt"

	"github.com/forbole/juno/v3/modules"
	"github.com/forbole/juno/v3/modules/registrar"
	"github.com/forbole/juno/v3/node/builder"
	"github.com/forbole/juno/v3/node/remote"

	internaldb "github.com/AutonomyNetwork/iam/database"
	"github.com/AutonomyNetwork/iam/modules/bot"
	"github.com/AutonomyNetwork/iam/modules/marketplace"
)

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

	// Calling Channels
	BotConnection.Channels()
	BotConnection.AddHandleMsgCreate()
	BotConnection.GetAddressAndValidate()

	return []modules.Module{
		onftModule,
		BotConnection,
	}
}
