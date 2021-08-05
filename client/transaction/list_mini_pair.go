package transaction

import (
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type ListMiniPairResult struct {
	txtype.TxCommitResult
}

func (c *client) ListMiniPair(baseAssetSymbol string, quoteAssetSymbol string, initPrice int64, sync bool, options ...Option) (*ListMiniPairResult, error) {
	fromAddr := c.keyManager.GetAddr()

	listMsg := msgtype.NewListMiniMsg(fromAddr, baseAssetSymbol, quoteAssetSymbol, initPrice)
	commit, err := c.broadcastMsg(listMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &ListMiniPairResult{*commit}, nil

}
