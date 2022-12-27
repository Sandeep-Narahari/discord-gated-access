package main

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	junomessages "github.com/forbole/juno/v3/modules/messages"

	"github.com/OmniFlix/marketplace/x/marketplace/types"
	onft "github.com/OmniFlix/onft/types"
)

var omniFlixMessageAddressesParser = junomessages.JoinMessageParsers(
	MarketPlaceAddressesParser,
	OnftAddressesParser,
)

func MarketPlaceAddressesParser(cdc codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	fmt.Println("Codec", cdc)
	switch msg := cosmosMsg.(type) {
	case *types.MsgListNFT:
		return []string{msg.Owner}, nil
	case *types.MsgEditListing:
		return []string{msg.Owner}, nil
	case *types.MsgDeListNFT:
		return []string{msg.Owner}, nil
	case *types.MsgBuyNFT:
		return []string{msg.Buyer}, nil
	case *types.MsgCreateAuction:
		return []string{msg.Owner}, nil
	case *types.MsgCancelAuction:
		return []string{msg.Owner}, nil
	case *types.MsgPlaceBid:
		return []string{msg.Bidder}, nil
	}
	return nil, junomessages.MessageNotSupported(cosmosMsg)
}

func OnftAddressesParser(cdc codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {
	case *onft.MsgCreateDenom:
		return []string{msg.Sender}, nil
	case *onft.MsgUpdateDenom:
		return []string{msg.Sender}, nil
	case *onft.MsgTransferDenom:
		return []string{msg.Sender}, nil
	case *onft.MsgMintONFT:
		return []string{msg.Sender}, nil
	case *onft.MsgTransferONFT:
		return []string{msg.Sender}, nil
	case *onft.MsgBurnONFT:
		return []string{msg.Sender}, nil
	}
	return nil, junomessages.MessageNotSupported(cosmosMsg)
}
