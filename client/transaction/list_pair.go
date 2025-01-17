package transaction

import (
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type ListPairResult struct {
	txtype.TxCommitResult
}

func (c *client) ListPair(proposalId int64, baseAssetSymbol string, quoteAssetSymbol string, initPrice int64, sync bool, options ...Option) (*ListPairResult, error) {
	fromAddr := c.keyManager.GetAddr()

	burnMsg := msgtype.NewDexListMsg(fromAddr, proposalId, baseAssetSymbol, quoteAssetSymbol, initPrice)
	commit, err := c.broadcastMsg(burnMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &ListPairResult{*commit}, nil

}
