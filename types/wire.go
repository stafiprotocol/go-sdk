package types

import (
	ntypes "github.com/stafiprotocol/go-sdk/common/types"
	"github.com/stafiprotocol/go-sdk/types/txtype"
	rpctypes "github.com/stafiprotocol/tendermint/rpc/core/types"
	"github.com/tendermint/go-amino"
)

func NewCodec() *amino.Codec {
	cdc := amino.NewCodec()
	rpctypes.RegisterAmino(cdc)
	ntypes.RegisterWire(cdc)
	txtype.RegisterCodec(cdc)
	return cdc
}
