package transaction

import (
	"fmt"
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type FreezeTokenResult struct {
	txtype.TxCommitResult
}

func (c *client) FreezeToken(symbol string, amount int64, sync bool, options ...Option) (*FreezeTokenResult, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Freeze token symbol can't be empty ")
	}
	fromAddr := c.keyManager.GetAddr()

	freezeMsg := msgtype.NewFreezeMsg(
		fromAddr,
		symbol,
		amount,
	)
	commit, err := c.broadcastMsg(freezeMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &FreezeTokenResult{*commit}, nil

}
