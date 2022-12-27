package marketplace

import (
	"fmt"

	"github.com/OmniFlix/marketplace/x/marketplace/types"
	onft "github.com/OmniFlix/onft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
	"github.com/gogo/protobuf/proto"
	"github.com/rs/zerolog/log"
)

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	log.Debug().Str("module", "onfts").Str("message", proto.MessageName(msg)).
		Int64("height", tx.Height).Msg("handled message")

	switch omniflixMsg := msg.(type) {
	case *types.MsgListNFT:
		// return m.handleMsgListNFT(tx.Height, omniflixMsg) //TODO: add
	case *onft.MsgCreateDenom:
		return m.handleMsgCreateDenom(tx.Height, omniflixMsg)
	default:
		return fmt.Errorf("invalid type to write %w", omniflixMsg)
	}

	return nil
}

func (m *Module) handleMsgListNFT(height int64, msg *types.MsgListNFT) {

}

func (m *Module) handleMsgCreateDenom(height int64, msg *onft.MsgCreateDenom) error {
	// denom, err := m.client.Denom(remote.GetHeightRequestContext(context.Background(), height), &onft.QueryDenomRequest{DenomId: msg.Id})
	// if err != nil {
	// 	return err
	// }
	// m.convertDenom(height, denom.Denom)
	// return m.db.SaveDenomsData()

	return nil
}
