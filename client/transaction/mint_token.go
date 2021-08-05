package transaction

import (
	"fmt"
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type MintTokenResult struct {
	txtype.TxCommitResult
}

func (c *client) MintToken(symbol string, amount int64, sync bool, options ...Option) (*MintTokenResult, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Mint token symbol can't be empty ")
	}
	fromAddr := c.keyManager.GetAddr()

	mintMsg := msgtype.NewMintMsg(
		fromAddr,
		symbol,
		amount,
	)
	commit, err := c.broadcastMsg(mintMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &MintTokenResult{*commit}, nil

}
