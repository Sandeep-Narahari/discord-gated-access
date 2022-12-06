package main

import (
	"log"
	
	"github.com/cosmos/cosmos-sdk/types/module"
	gaiaapp "github.com/cosmos/gaia/v7/app"
	
	"github.com/forbole/bdjuno/v3/types/config"
	"github.com/forbole/juno/v3/cmd"
	parsetypes "github.com/forbole/juno/v3/cmd/parse/types"
	"github.com/forbole/juno/v3/modules/messages"
	
	"github.com/OmniFlix/omniflixhub/app"
	
	internaldb "github.com/PrithviDevs/ojuno/database"
)

func main() {
	parseCfg := parsetypes.NewConfig().
		WithEncodingConfigBuilder(config.MakeEncodingConfig(getBasicMangers())). // TODO: check
		WithRegistrar(nil). // modules.NewModuleRegistrar()
		WithDBBuilder(internaldb.Builder)
	
	cfg := cmd.NewConfig("ojuno").
		WithParseConfig(parseCfg)
	
	// Run the commands
	executor := cmd.BuildDefaultExecutor(cfg)
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
