package main

import (
	"fmt"
	"strconv"

	"github.com/stafiprotocol/go-sdk/client"
	"github.com/stafiprotocol/go-sdk/common/types"
	"github.com/stafiprotocol/go-sdk/keys"
	"github.com/stafiprotocol/go-sdk/types/msgtype"
)

// To run this example, please make sure your key address have more than 1:BNB on testnet
func main() {
	types.Network = types.TestNetwork

	keyManager, err := keys.NewPrivateKeyManager("64967ded205b00b1f872f59242031d4cc02a1bcca47017361d7f3854e86c545e")
	keyManager.GetAddr()

	receiverAddr, err := types.AccAddressFromBech32("tbnb1tt84yhkvh6q23kksttfq36dujnyfh2cldrzux5")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dexClient, err := client.NewDexClient("testnet-dex.binance.org:443", types.TestNetwork, keyManager)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	account, err := dexClient.GetAccount(keyManager.GetAddr().String())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	floatAmount := 0.0
	for _, coin := range account.Balances {
		if coin.Symbol == "BNB" {
			fmt.Println(fmt.Sprintf("Your account has %s:BNB", coin.Free))
			floatAmount, err = strconv.ParseFloat(coin.Free.String(), 64)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			break
		}
	}
	if floatAmount <= 1.0 {
		fmt.Println("Your account doesn't have enough bnb")
	}

	fmt.Println(fmt.Sprintf("Please verify sign key address (%s) and transaction data", types.AccAddress(keyManager.GetAddr()).String()))
	sendResult, err := dexClient.SendToken([]msgtype.Transfer{{receiverAddr, types.Coins{types.Coin{Denom: "BNB", Amount: 10000000}}}}, true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("Send result: %t", sendResult.Ok))
}
