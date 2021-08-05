package txtype

import (
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	cryptoAmino "github.com/stafiprotocol/tendermint/crypto/encoding/amino"
	goamino "github.com/tendermint/go-amino"
)

// cdc global variable
var Cdc = goamino.NewCodec()

func RegisterCodec(cdc *goamino.Codec) {
	cdc.RegisterInterface((*Tx)(nil), nil)
	cdc.RegisterConcrete(StdTx{}, "auth/StdTx", nil)
	msgtype.RegisterCodec(cdc)
}

func init() {
	cryptoAmino.RegisterAmino(Cdc)
	RegisterCodec(Cdc)
}
