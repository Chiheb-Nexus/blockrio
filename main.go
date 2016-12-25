package main

import (
	"fmt"
	"blockrio/api"
)

func main() {

	/*
		Get multiple addresses balance.
		Method:
			- Using: GetAddressBalance(coin string, add []string) -> return []ResponseAddressBalance
			- Extracted data are:
				* Status -> string
				* Data -> ResponseAddressBalance.[]_DataBalance{
												Address -> string
												Balance -> float64
										BalanceMultiSig -> float64
												}
					* Code -> float64
					* Message -> string
			So, in order to get the addresses's balance we should range at the retern values of GetAddressBalance
				Example: v := GetAddressBalance(coin, add)
			Then range second time at the value of v.Data using "val" as key.
			Finally we can retrieve Address and Balance with:
				Address: val.Address
				Balance: val.Balance
	*/
	add := []string{
		"198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi",
		"1L8meqhMTRpxasdGt8DHSJfscxgHHzvPgk",
	}
	coin := "btc"
	addressBalance := blockrio.GetAddressBalance(coin, add)
	for _, value := range addressBalance {
		for _, val := range value.Data {
			fmt.Printf("Address: %s  |  Balance: %v\n", val.Address, val.Balance)
		}
	}
}
