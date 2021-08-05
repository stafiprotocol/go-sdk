package transaction

import (
	"fmt"
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type UnfreezeTokenResult struct {
	txtype.TxCommitResult
}

func (c *client) UnfreezeToken(symbol string, amount int64, sync bool, options ...Option) (*UnfreezeTokenResult, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Unfreeze token symbol can't be empty ")
	}
	fromAddr := c.keyManager.GetAddr()

	unfreezeMsg := msgtype.NewUnfreezeMsg(
		fromAddr,
		symbol,
		amount,
	)
	commit, err := c.broadcastMsg(unfreezeMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &UnfreezeTokenResult{*commit}, nil

}
