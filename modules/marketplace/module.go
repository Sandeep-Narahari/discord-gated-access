package marketplace

import (
	"github.com/OmniFlix/onft/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v3/modules"
	"github.com/forbole/juno/v3/node"
	"google.golang.org/grpc"
	
	"github.com/AutonomyNetwork/iam/database"
)

var (
	_ modules.Module        = &Module{}
	_ modules.MessageModule = &Module{}
)

type Module struct {
	db     *database.Db
	cdc    codec.Codec
	node   node.Node
	client types.QueryClient
}

func NewModule(node node.Node, grpcConnection *grpc.ClientConn, cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		db:     db,
		cdc:    cdc,
		node:   node,
		client: types.NewQueryClient(grpcConnection),
	}
}

func (m *Module) Name() string {
	return "onft"
}
