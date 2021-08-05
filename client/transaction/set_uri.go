package transaction

import (
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type SetUriResult struct {
	txtype.TxCommitResult
}

func (c *client) SetURI(symbol, tokenURI string, sync bool, options ...Option) (*SetUriResult, error) {
	fromAddr := c.keyManager.GetAddr()

	setURIMsg := msgtype.NewSetUriMsg(fromAddr, symbol, tokenURI)
	commit, err := c.broadcastMsg(setURIMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &SetUriResult{*commit}, nil

}
