package query

import (
	"encoding/json"

	"github.com/stafiprotocol/go-sdk/common"
	"github.com/stafiprotocol/go-sdk/common/types"
)

// GetTrades returns transaction details
func (c *client) GetTrades(tradesQuery *types.TradesQuery) (*types.Trades, error) {
	err := tradesQuery.Check()
	if err != nil {
		return nil, err
	}
	qp, err := common.QueryParamToMap(*tradesQuery)
	if err != nil {
		return nil, err
	}

	resp, _, err := c.baseClient.Get("/trades", qp)
	if err != nil {
		return nil, err
	}

	var trades types.Trades
	if err := json.Unmarshal(resp, &trades); err != nil {
		return nil, err
	}

	return &trades, nil
}
