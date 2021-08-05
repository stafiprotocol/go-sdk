// +build cgo,ledgercosmos

package ledger

import (
	ledgercosmos "github.com/binance-chain/ledger-cosmos-go"
)

// If ledgercosmos support (build tag) has been enabled, which implies a CGO dependency,
// set the discoverLedger function which is responsible for loading the Ledger
// device at runtime or returning an error.
func init() {
	DiscoverLedger = func() (LedgerSecp256k1, error) {
		device, err := ledgercosmos.FindLedgerCosmosUserApp()
		if err != nil {
			return nil, err
		}

		return device, nil
	}
}
