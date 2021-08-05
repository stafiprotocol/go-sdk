package transaction

import (
	"fmt"

	"github.com/stafiprotocol/go-sdk/common"
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type CancelOrderResult struct {
	txtype.TxCommitResult
}

func (c *client) CancelOrder(baseAssetSymbol, quoteAssetSymbol, refId string, sync bool, options ...Option) (*CancelOrderResult, error) {
	if baseAssetSymbol == "" || quoteAssetSymbol == "" {
		return nil, fmt.Errorf("BaseAssetSymbol or QuoteAssetSymbol is missing. ")
	}
	if refId == "" {
		return nil, fmt.Errorf("OrderId or Order RefId is missing. ")
	}

	fromAddr := c.keyManager.GetAddr()

	cancelOrderMsg := msgtype.NewCancelOrderMsg(fromAddr, common.CombineSymbol(baseAssetSymbol, quoteAssetSymbol), refId)
	commit, err := c.broadcastMsg(cancelOrderMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &CancelOrderResult{*commit}, nil
}
