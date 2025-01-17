package rpc

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	cmn "github.com/stafiprotocol/tendermint/libs/common"
	"github.com/stafiprotocol/tendermint/rpc/client"
	ctypes "github.com/stafiprotocol/tendermint/rpc/core/types"
	"github.com/stafiprotocol/tendermint/rpc/lib/client"
	"github.com/stafiprotocol/tendermint/types"

	ntypes "github.com/stafiprotocol/go-sdk/common/types"
	"github.com/stafiprotocol/go-sdk/keys"
	"github.com/stafiprotocol/go-sdk/types/txtype"
)

var DefaultTimeout = 5 * time.Second

type ABCIClient interface {
	// Reading from abci app
	ABCIInfo() (*ctypes.ResultABCIInfo, error)
	ABCIQuery(path string, data cmn.HexBytes) (*ctypes.ResultABCIQuery, error)
	ABCIQueryWithOptions(path string, data cmn.HexBytes,
		opts client.ABCIQueryOptions) (*ctypes.ResultABCIQuery, error)

	// Writing to abci app
	BroadcastTxCommit(tx types.Tx) (*ResultBroadcastTxCommit, error)
	BroadcastTxAsync(tx types.Tx) (*ctypes.ResultBroadcastTx, error)
	BroadcastTxSync(tx types.Tx) (*ctypes.ResultBroadcastTx, error)
}

type SignClient interface {
	Block(height *int64) (*ctypes.ResultBlock, error)
	BlockResults(height *int64) (*ResultBlockResults, error)
	Commit(height *int64) (*ctypes.ResultCommit, error)
	Validators(height *int64) (*ctypes.ResultValidators, error)
	Tx(hash []byte, prove bool) (*ResultTx, error)
	TxSearch(query string, prove bool, page, perPage int) (*ResultTxSearch, error)
}

type Client interface {
	cmn.Service
	ABCIClient
	SignClient
	client.HistoryClient
	client.StatusClient
	EventsClient
	DexClient
	OpsClient
	StakingClient
}

type EventsClient interface {
	Subscribe(query string, outCapacity ...int) (out chan ctypes.ResultEvent, err error)
	Unsubscribe(query string) error
	UnsubscribeAll() error
}

func NewRPCClient(nodeURI string, network ntypes.ChainNetwork) *HTTP {
	ntypes.Network = network
	return NewHTTP(nodeURI, "/websocket")
}

type HTTP struct {
	*WSEvents

	key keys.KeyManager
}

// NewHTTP takes a remote endpoint in the form tcp://<host>:<port>
// and the websocket path (which always seems to be "/websocket")
func NewHTTP(remote, wsEndpoint string) *HTTP {
	rc := rpcclient.NewJSONRPCClient(remote)
	cdc := rc.Codec()
	ctypes.RegisterAmino(cdc)
	ntypes.RegisterWire(cdc)
	txtype.RegisterCodec(cdc)

	rc.SetCodec(cdc)
	wsEvent := newWSEvents(cdc, remote, wsEndpoint)
	client := &HTTP{
		WSEvents: wsEvent,
	}
	client.Start()
	return client
}

func (c *HTTP) Status() (*ctypes.ResultStatus, error) {
	return c.WSEvents.Status()
}

func (c *HTTP) ABCIInfo() (*ctypes.ResultABCIInfo, error) {
	return c.WSEvents.ABCIInfo()
}

func (c *HTTP) ABCIQuery(path string, data cmn.HexBytes) (*ctypes.ResultABCIQuery, error) {
	return c.ABCIQueryWithOptions(path, data, client.DefaultABCIQueryOptions)
}

func (c *HTTP) ABCIQueryWithOptions(path string, data cmn.HexBytes, opts client.ABCIQueryOptions) (*ctypes.ResultABCIQuery, error) {
	if err := ValidateABCIPath(path); err != nil {
		return nil, err
	}
	if err := ValidateABCIData(data); err != nil {
		return nil, err
	}
	return c.WSEvents.ABCIQueryWithOptions(path, data, opts)
}

func (c *HTTP) BroadcastTxCommit(tx types.Tx) (*ResultBroadcastTxCommit, error) {
	if err := ValidateTx(tx); err != nil {
		return nil, err
	}
	return c.WSEvents.BroadcastTxCommit(tx)
}

func (c *HTTP) BroadcastTxAsync(tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	if err := ValidateTx(tx); err != nil {
		return nil, err
	}
	return c.WSEvents.BroadcastTx("broadcast_tx_async", tx)
}

func (c *HTTP) BroadcastTxSync(tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	if err := ValidateTx(tx); err != nil {
		return nil, err
	}
	return c.WSEvents.BroadcastTx("broadcast_tx_sync", tx)
}

func (c *HTTP) UnconfirmedTxs(limit int) (*ctypes.ResultUnconfirmedTxs, error) {
	if err := ValidateUnConfirmedTxsLimit(limit); err != nil {
		return nil, err
	}
	return c.WSEvents.UnconfirmedTxs(limit)
}

func (c *HTTP) NumUnconfirmedTxs() (*ctypes.ResultUnconfirmedTxs, error) {
	return c.WSEvents.NumUnconfirmedTxs()
}

func (c *HTTP) NetInfo() (*ctypes.ResultNetInfo, error) {
	return c.WSEvents.NetInfo()
}

func (c *HTTP) DumpConsensusState() (*ctypes.ResultDumpConsensusState, error) {
	return c.WSEvents.DumpConsensusState()
}

func (c *HTTP) ConsensusState() (*ctypes.ResultConsensusState, error) {
	return c.WSEvents.ConsensusState()
}

func (c *HTTP) Health() (*ctypes.ResultHealth, error) {
	return c.WSEvents.Health()
}

func (c *HTTP) BlockchainInfo(minHeight, maxHeight int64) (*ctypes.ResultBlockchainInfo, error) {
	if err := ValidateHeightRange(minHeight, maxHeight); err != nil {
		return nil, err
	}
	return c.WSEvents.BlockchainInfo(minHeight, maxHeight)
}

func (c *HTTP) Genesis() (*ctypes.ResultGenesis, error) {
	return c.WSEvents.Genesis()
}

func (c *HTTP) Block(height *int64) (*ctypes.ResultBlock, error) {
	if err := ValidateHeight(height); err != nil {
		return nil, err
	}
	return c.WSEvents.Block(height)
}

func (c *HTTP) BlockResults(height *int64) (*ResultBlockResults, error) {
	if err := ValidateHeight(height); err != nil {
		return nil, err
	}
	return c.WSEvents.BlockResults(height)
}

func (c *HTTP) Commit(height *int64) (*ctypes.ResultCommit, error) {
	if err := ValidateHeight(height); err != nil {
		return nil, err
	}
	return c.WSEvents.Commit(height)
}

func (c *HTTP) Tx(hash []byte, prove bool) (*ResultTx, error) {
	if err := ValidateHash(hash); err != nil {
		return nil, err
	}
	return c.WSEvents.Tx(hash, prove)
}

func (c *HTTP) TxSearch(query string, prove bool, page, perPage int) (*ResultTxSearch, error) {
	if err := ValidateABCIQueryStr(query); err != nil {
		return nil, err
	}
	return c.WSEvents.TxSearch(query, prove, page, perPage)
}

func (c *HTTP) Validators(height *int64) (*ctypes.ResultValidators, error) {
	if err := ValidateHeight(height); err != nil {
		return nil, err
	}
	return c.WSEvents.Validators(height)
}

func (c *HTTP) QueryWithData(path string, data cmn.HexBytes) ([]byte, error) {
	result, err := c.ABCIQuery(path, data)

	if err != nil {
		return nil, err
	}

	resp := result.Response
	if !resp.IsOK() {
		return nil, errors.Errorf(resp.Log)
	}

	return resp.Value, nil
}

func (c *HTTP) QueryStore(key cmn.HexBytes, storeName string) ([]byte, error) {
	path := fmt.Sprintf("/store/%s/%s", storeName, "key")
	result, err := c.ABCIQuery(path, key)
	if err != nil {
		return nil, err
	}
	resp := result.Response
	if !resp.IsOK() {
		return nil, errors.Errorf(resp.Log)
	}
	return resp.Value, nil
}

func (c *HTTP) QueryStoreSubspace(key cmn.HexBytes, storeName string) (res []cmn.KVPair, err error) {
	path := fmt.Sprintf("/store/%s/subspace", storeName)
	result, err := c.ABCIQuery(path, key)
	if err != nil {
		return res, err
	}

	resp := result.Response
	if !resp.IsOK() {
		return nil, errors.Errorf(resp.Log)
	}

	if len(resp.Value) == 0 {
		return nil, EmptyResultError
	}

	err = c.cdc.UnmarshalBinaryLengthPrefixed(resp.Value, &res)

	if err != nil {
		return nil, err
	}

	return
}

func (c *HTTP) SetKeyManager(k keys.KeyManager) {
	c.key = k
}
