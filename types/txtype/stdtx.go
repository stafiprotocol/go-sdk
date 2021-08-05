package txtype

import (
	"github.com/stafiprotocol/go-sdk/types/msgtype"
)

const Source int64 = 0

type Tx interface {

	// Gets the Msg.
	GetMsgs() []msgtype.Msg
}

// StdTx def
type StdTx struct {
	Msgs       []msgtype.Msg  `json:"msg"`
	Signatures []StdSignature `json:"signatures"`
	Memo       string         `json:"memo"`
	Source     int64          `json:"source"`
	Data       []byte         `json:"data"`
}

// NewStdTx to instantiate an instance
func NewStdTx(msgs []msgtype.Msg, sigs []StdSignature, memo string, source int64, data []byte) StdTx {
	return StdTx{
		Msgs:       msgs,
		Signatures: sigs,
		Memo:       memo,
		Source:     source,
		Data:       data,
	}
}

// GetMsgs def
func (tx StdTx) GetMsgs() []msgtype.Msg { return tx.Msgs }
