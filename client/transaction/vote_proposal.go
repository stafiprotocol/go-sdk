package transaction

import (
	"github.com/stafiprotocol/go-sdk/types/msgtype"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

type VoteProposalResult struct {
	txtype.TxCommitResult
}

func (c *client) VoteProposal(proposalID int64, option msgtype.VoteOption, sync bool, options ...Option) (*VoteProposalResult, error) {
	fromAddr := c.keyManager.GetAddr()
	voteMsg := msgtype.NewMsgVote(fromAddr, proposalID, option)
	commit, err := c.broadcastMsg(voteMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &VoteProposalResult{*commit}, err

}
