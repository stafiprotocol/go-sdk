package transaction

import (
	"fmt"
	"github.com/stafiprotocol/go-sdk/common/types"
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type SetAccountFlagsResult struct {
	txtype.TxCommitResult
}

func (c *client) AddAccountFlags(flagOptions []types.FlagOption, sync bool, options ...Option) (*SetAccountFlagsResult, error) {
	fromAddr := c.keyManager.GetAddr()
	acc, err := c.queryClient.GetAccount(fromAddr.String())
	if err != nil {
		return nil, err
	}
	if len(flagOptions) == 0 {
		return nil, fmt.Errorf("missing flagOptions")
	}
	flags := acc.Flags
	for _, f := range flagOptions {
		flags = flags | uint64(f)
	}
	setAccMsg := msgtype.NewSetAccountFlagsMsg(
		fromAddr,
		flags,
	)
	commit, err := c.broadcastMsg(setAccMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &SetAccountFlagsResult{*commit}, nil
}

func (c *client) SetAccountFlags(flags uint64, sync bool, options ...Option) (*SetAccountFlagsResult, error) {
	fromAddr := c.keyManager.GetAddr()
	setAccMsg := msgtype.NewSetAccountFlagsMsg(
		fromAddr,
		flags,
	)
	commit, err := c.broadcastMsg(setAccMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &SetAccountFlagsResult{*commit}, nil
}
