package transaction

import (
	"github.com/stafiprotocol/go-sdk/common/types"
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type TransferTokenOwnershipResult struct {
	txtype.TxCommitResult
}

func (c *client) TransferTokenOwnership(symbol string, newOwner types.AccAddress, sync bool, options ...Option) (*TransferTokenOwnershipResult, error) {
	fromAddr := c.keyManager.GetAddr()
	transferOwnershipMsg := msgtype.NewTransferOwnershipMsg(fromAddr, symbol, newOwner)
	commit, err := c.broadcastMsg(transferOwnershipMsg, sync, options...)
	if err != nil {
		return nil, err
	}
	return &TransferTokenOwnershipResult{*commit}, nil
}
