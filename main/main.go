package main

import (
	"Blockr.io/api"
	"fmt"
)

func main() {
	//coin_info := Blockrio.GetCoinInfo("btc")
	//fmt.Println("coin_info: ", coin_info)

	//fmt.Println("-------------------------------------")
	//coin_current := Blockrio.GetExchangeCurrent("btc")
	//fmt.Println("coin_current: ", coin_current)

	//fmt.Println("-------------------------------------")
	//add := []string{"1","2"}
	//block := Blockrio.GetBlockInfo("btc", add)
	//fmt.Println("block: ", block)

	//fmt.Println("-------------------------------------")
	//add := []string{"12345","12346"}
	//block_tx := Blockrio.GetBlockTxs("btc", add)
	//fmt.Println("block_tx: ", block_tx)

	//fmt.Println("-------------------------------------")
	//add := []string{
	//	"60c1f1a3160042152114e2bba45600a5045711c3a8a458016248acec59653471",
	//	"4addbc5ec75e087a44a34e8c1c3bd05fd495771072879a89a8c9aaa356255cb2",
	//	}
	//tx_info := Blockrio.GetTxsInfo("btc", add)
	//fmt.Println("tx_info: ", tx_info)

	//fmt.Println("-------------------------------------")
	//add := []string{"12345","12346"}
	//block_raw := Blockrio.GetBlockRaw("btc", add)
	//fmt.Println("block_raw: ", block_raw)

	//fmt.Println("-------------------------------------")
	//add := []string{
	//	"60c1f1a3160042152114e2bba45600a5045711c3a8a458016248acec59653471",
	//	"4addbc5ec75e087a44a34e8c1c3bd05fd495771072879a89a8c9aaa356255cb2",
	//	}
	//tx_info := Blockrio.GetAddressInfo("btc", add)
	//fmt.Println("tx_info: ", tx_info)

	//fmt.Println("-------------------------------------")
	//add := []string{
	//	"198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi",
	//	"1L8meqhMTRpxasdGt8DHSJfscxgHHzvPgk",
	//	}
	//address_balance := Blockrio.GetAddressBalance("btc", add)
	//fmt.Println("address_balance: ", address_balance)

	//fmt.Println("-------------------------------------")
	//add := []string{
	//	"198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi",
	//	"1L8meqhMTRpxasdGt8DHSJfscxgHHzvPgk",
	//	}
	//address_transactions := Blockrio.GetAddressTransactions("btc", add)
	//fmt.Println("address_transactions: ", address_transactions)

	//fmt.Println("-------------------------------------")
	//add := []string{
	//	"198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi",
	//	"1L8meqhMTRpxasdGt8DHSJfscxgHHzvPgk",
	//	}
	//address_unspent := Blockrio.GetAddressUnspent("btc", add)
	//fmt.Println("address_unspent: ", address_unspent)

	fmt.Println("-------------------------------------")
	add := []string{
		"3PszH9JUX9XTuwSMhEJyk2j7NKpVStT9Ho",
		"39zNw2y4xtrFyn1qkndBFhRzPSth8V89ev",
	}
	address_unconfirmed_tx := Blockrio.GetUnconfirmedTxBalance("btc", add)
	fmt.Println("address_unconfirmed_tx: ", address_unconfirmed_tx)

}
