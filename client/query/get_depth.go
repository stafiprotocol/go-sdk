package query

import (
	"encoding/json"

	"github.com/stafiprotocol/go-sdk/common"
	"github.com/stafiprotocol/go-sdk/common/types"
)

// GetDepth returns market depth records
func (c *client) GetDepth(depthQuery *types.DepthQuery) (*types.MarketDepth, error) {
	err := depthQuery.Check()
	if err != nil {
		return nil, err
	}
	qp, err := common.QueryParamToMap(*depthQuery)
	if err != nil {
		return nil, err
	}
	resp, _, err := c.baseClient.Get("/depth", qp)
	if err != nil {
		return nil, err
	}

	var MarketDepth types.MarketDepth
	if err := json.Unmarshal(resp, &MarketDepth); err != nil {
		return nil, err
	}

	return &MarketDepth, nil
}
