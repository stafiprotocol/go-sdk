package query

import (
	"encoding/json"

	"github.com/stafiprotocol/go-sdk/common"
	"github.com/stafiprotocol/go-sdk/common/types"
)

// GetClosedOrders returns array of open orders
func (c *client) GetClosedOrders(closeOrderQuery *types.ClosedOrdersQuery) (*types.CloseOrders, error) {
	err := closeOrderQuery.Check()
	if err != nil {
		return nil, err
	}
	qp, err := common.QueryParamToMap(*closeOrderQuery)
	if err != nil {
		return nil, err
	}
	resp, _, err := c.baseClient.Get("/orders/closed", qp)
	if err != nil {
		return nil, err
	}

	var orders types.CloseOrders
	if err := json.Unmarshal(resp, &orders); err != nil {
		return nil, err
	}

	return &orders, nil
}
