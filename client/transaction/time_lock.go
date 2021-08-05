package transaction

import (
	"strconv"

	"github.com/stafiprotocol/go-sdk/common/types"
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type TimeLockResult struct {
	txtype.TxCommitResult
	LockId int64 `json:"lock_id"`
}

func (c *client) TimeLock(description string, amount types.Coins, lockTime int64, sync bool, options ...Option) (*TimeLockResult, error) {
	fromAddr := c.keyManager.GetAddr()

	lockMsg := msgtype.NewTimeLockMsg(fromAddr, description, amount, lockTime)
	commit, err := c.broadcastMsg(lockMsg, sync, options...)
	if err != nil {
		return nil, err
	}
	var lockId int64
	if commit.Ok && sync {
		lockId, err = strconv.ParseInt(string(commit.Data), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return &TimeLockResult{*commit, lockId}, err
}

type TimeUnLockResult struct {
	txtype.TxCommitResult
	LockId int64 `json:"lock_id"`
}

func (c *client) TimeUnLock(id int64, sync bool, options ...Option) (*TimeUnLockResult, error) {
	fromAddr := c.keyManager.GetAddr()

	unlockMsg := msgtype.NewTimeUnlockMsg(fromAddr, id)
	err := unlockMsg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	commit, err := c.broadcastMsg(unlockMsg, sync, options...)
	if err != nil {
		return nil, err
	}
	var lockId int64
	if commit.Ok && sync {
		lockId, err = strconv.ParseInt(string(commit.Data), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return &TimeUnLockResult{*commit, lockId}, err
}

type TimeReLockResult struct {
	txtype.TxCommitResult
	LockId int64 `json:"lock_id"`
}

func (c *client) TimeReLock(id int64, description string, amount types.Coins, lockTime int64, sync bool, options ...Option) (*TimeReLockResult, error) {
	fromAddr := c.keyManager.GetAddr()

	relockMsg := msgtype.NewTimeRelockMsg(fromAddr, id, description, amount, lockTime)
	err := relockMsg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	commit, err := c.broadcastMsg(relockMsg, sync, options...)
	if err != nil {
		return nil, err
	}
	var lockId int64
	if commit.Ok && sync {
		lockId, err = strconv.ParseInt(string(commit.Data), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return &TimeReLockResult{*commit, lockId}, err
}
