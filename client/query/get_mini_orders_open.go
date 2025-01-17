package query

import (
	"encoding/json"

	"github.com/stafiprotocol/go-sdk/common"
	"github.com/stafiprotocol/go-sdk/common/types"
)

// GetMiniOpenOrders returns array of mini open orders
func (c *client) GetMiniOpenOrders(openOrderQuery *types.OpenOrdersQuery) (*types.OpenOrders, error) {
	err := openOrderQuery.Check()
	if err != nil {
		return nil, err
	}
	qp, err := common.QueryParamToMap(*openOrderQuery)
	if err != nil {
		return nil, err
	}

	resp, _, err := c.baseClient.Get("/mini/orders/open", qp)
	if err != nil {
		return nil, err
	}

	var openOrders types.OpenOrders
	if err := json.Unmarshal(resp, &openOrders); err != nil {
		return nil, err
	}

	return &openOrders, nil
}
