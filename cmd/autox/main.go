package main

import (
	"log"

	"github.com/cosmos/cosmos-sdk/types/module"
	gaiaapp "github.com/cosmos/gaia/v7/app"

	migratecmd "github.com/forbole/bdjuno/v3/cmd/migrate"
	parsecmd "github.com/forbole/bdjuno/v3/cmd/parse"
	startcmd "github.com/forbole/juno/v3/cmd/start"

	// "github.com/forbole/bdjuno/v3/types/config"

	"github.com/AutonomyNetwork/iam/types/config"
	"github.com/forbole/juno/v3/cmd"
	initcmd "github.com/forbole/juno/v3/cmd/init"
	parsetypes "github.com/forbole/juno/v3/cmd/parse/types"
	"github.com/forbole/juno/v3/modules/messages"

	internaldb "github.com/AutonomyNetwork/iam/database"
	"github.com/AutonomyNetwork/iam/modules"
	"github.com/OmniFlix/omniflixhub/app"
)

func main() {
	initCfg := initcmd.NewConfig().
		WithConfigCreator(config.Creator)
	parseCfg := parsetypes.NewConfig().
		WithEncodingConfigBuilder(config.MakeEncodingConfig(getBasicMangers())). // TODO: check
		WithRegistrar(modules.NewModuleRegistrar()).                             // modules.NewModuleRegistrar()
		WithDBBuilder(internaldb.Builder)

	cfg := cmd.NewConfig("autox").
		WithInitConfig(initCfg).
		WithParseConfig(parseCfg)

	// Run the commands
	rootCmd := cmd.RootCmd(cfg.GetName())

	rootCmd.AddCommand(
		cmd.VersionCmd(),
		initcmd.NewInitCmd(cfg.GetInitConfig()),
		parsecmd.NewParseCmd(cfg.GetParseConfig()),
		migratecmd.NewMigrateCmd(cfg.GetName(), cfg.GetParseConfig()),
		startcmd.NewStartCmd(cfg.GetParseConfig()),
	)

	executor := cmd.PrepareRootCmd(cfg.GetName(), rootCmd)
	err := executor.Execute()
	if err != nil {
		log.Fatal(err)
	}

}

func getBasicMangers() []module.BasicManager {
	return []module.BasicManager{
		gaiaapp.ModuleBasics,
		app.ModuleBasics,
	}
}

func getAddressesParser() messages.MessageAddressesParser {
	return messages.JoinMessageParsers(
		omniFlixMessageAddressesParser,
		messages.CosmosMessageAddressesParser,
	)
}
