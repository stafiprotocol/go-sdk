package transaction

import (
	"github.com/stafiprotocol/go-sdk/common/types"
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type HTLTResult struct {
	txtype.TxCommitResult
}

func (c *client) HTLT(recipient types.AccAddress, recipientOtherChain, senderOtherChain string, randomNumberHash []byte, timestamp int64,
	amount types.Coins, expectedIncome string, heightSpan int64, crossChain bool, sync bool, options ...Option) (*HTLTResult, error) {
	fromAddr := c.keyManager.GetAddr()
	htltMsg := msgtype.NewHTLTMsg(
		fromAddr,
		recipient,
		recipientOtherChain,
		senderOtherChain,
		randomNumberHash,
		timestamp,
		amount,
		expectedIncome,
		heightSpan,
		crossChain,
	)
	commit, err := c.broadcastMsg(htltMsg, sync, options...)
	if err != nil {
		return nil, err
	}
	return &HTLTResult{*commit}, nil
}

type DepositHTLTResult struct {
	txtype.TxCommitResult
}

func (c *client) DepositHTLT(swapID []byte, amount types.Coins,
	sync bool, options ...Option) (*DepositHTLTResult, error) {
	fromAddr := c.keyManager.GetAddr()
	depositHTLTMsg := msgtype.NewDepositHTLTMsg(
		fromAddr,
		swapID,
		amount,
	)
	commit, err := c.broadcastMsg(depositHTLTMsg, sync, options...)
	if err != nil {
		return nil, err
	}
	return &DepositHTLTResult{*commit}, nil
}

type ClaimHTLTResult struct {
	txtype.TxCommitResult
}

func (c *client) ClaimHTLT(swapID []byte, randomNumber []byte, sync bool, options ...Option) (*ClaimHTLTResult, error) {
	fromAddr := c.keyManager.GetAddr()
	claimHTLTMsg := msgtype.NewClaimHTLTMsg(
		fromAddr,
		swapID,
		randomNumber,
	)
	commit, err := c.broadcastMsg(claimHTLTMsg, sync, options...)
	if err != nil {
		return nil, err
	}
	return &ClaimHTLTResult{*commit}, nil
}

type RefundHTLTResult struct {
	txtype.TxCommitResult
}

func (c *client) RefundHTLT(swapID []byte, sync bool, options ...Option) (*RefundHTLTResult, error) {
	fromAddr := c.keyManager.GetAddr()
	refundHTLTMsg := msgtype.NewRefundHTLTMsg(
		fromAddr,
		swapID,
	)
	commit, err := c.broadcastMsg(refundHTLTMsg, sync, options...)
	if err != nil {
		return nil, err
	}
	return &RefundHTLTResult{*commit}, nil
}
