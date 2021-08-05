package transaction

import (
	"fmt"
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type BurnTokenResult struct {
	txtype.TxCommitResult
}

func (c *client) BurnToken(symbol string, amount int64, sync bool, options ...Option) (*BurnTokenResult, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Burn token symbol can't be empty ")
	}
	fromAddr := c.keyManager.GetAddr()

	burnMsg := msgtype.NewTokenBurnMsg(
		fromAddr,
		symbol,
		amount,
	)
	commit, err := c.broadcastMsg(burnMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &BurnTokenResult{*commit}, nil

}
