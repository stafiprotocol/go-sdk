package transaction

import (
	"github.com/stafiprotocol/go-sdk/common/types"
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type SendTokenResult struct {
	txtype.TxCommitResult
}

func (c *client) SendToken(transfers []msgtype.Transfer, sync bool, options ...Option) (*SendTokenResult, error) {
	fromAddr := c.keyManager.GetAddr()
	fromCoins := types.Coins{}
	for _, t := range transfers {
		t.Coins = t.Coins.Sort()
		fromCoins = fromCoins.Plus(t.Coins)
	}
	sendMsg := msgtype.CreateSendMsg(fromAddr, fromCoins, transfers)
	commit, err := c.broadcastMsg(sendMsg, sync, options...)
	if err != nil {
		return nil, err
	}
	return &SendTokenResult{*commit}, err

}
