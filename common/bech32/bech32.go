package bech32

import (
	btcbech32 "github.com/btcsuite/btcutil/bech32"
	"github.com/pkg/errors"
)

//ConvertAndEncode converts from a base64 encoded byte string to base32 encoded byte string and then to btcbech32
func ConvertAndEncode(hrp string, data []byte) (string, error) {
	converted, err := btcbech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", errors.Wrap(err, "encoding btcbech32 failed")
	}
	return btcbech32.Encode(hrp, converted)

}

//DecodeAndConvert decodes a btcbech32 encoded string and converts to base64 encoded bytes
func DecodeAndConvert(bech string) (string, []byte, error) {
	hrp, data, err := btcbech32.Decode(bech)
	if err != nil {
		return "", nil, errors.Wrap(err, "decoding btcbech32 failed")
	}
	converted, err := btcbech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, errors.Wrap(err, "decoding btcbech32 failed")
	}
	return hrp, converted, nil
}
