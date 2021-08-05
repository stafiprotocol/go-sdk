package query

import (
	"encoding/json"

	"github.com/stafiprotocol/go-sdk/common"
	"github.com/stafiprotocol/go-sdk/common/types"
)

// GetMiniMarkets returns list of trading pairs
func (c *client) GetMiniMarkets(marketsQuery *types.MarketsQuery) ([]types.TradingPair, error) {
	err := marketsQuery.Check()
	if err != nil {
		return nil, err
	}
	qp, err := common.QueryParamToMap(*marketsQuery)
	if err != nil {
		return nil, err
	}
	resp, _, err := c.baseClient.Get("/mini/markets", qp)
	if err != nil {
		return nil, err
	}
	var listOfPairs []types.TradingPair
	if err := json.Unmarshal(resp, &listOfPairs); err != nil {
		return nil, err
	}

	return listOfPairs, nil
}
