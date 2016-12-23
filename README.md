# Wrapper for Blockr.io

[![Build status](https://ci.appveyor.com/api/projects/status/e5b3n9w516had2fd?svg=true)](https://ci.appveyor.com/project/Chiheb-Nexus/blockrio)
[![Go Report Card](https://goreportcard.com/badge/github.com/Chiheb-Nexus/blockrio)](https://goreportcard.com/report/github.com/Chiheb-Nexus/blockrio)

This wrapper is written in Golang and uses the API documentation of [blockr.io](http://blockr.io/documentation/api)

# Understand this wrapper
In order to understand this wrapper, you should understand how it works.

> main.go (main function)

> API_Callbacks.go (All the callbacks for this wrapper)

> API_FetchUrl.go (Load data from API URL's)

> API_UnmarshalJSONFromURL.go (Load structs and unmarshal returned JSON data)

# Usage

First of all setup your **$GOPATH** variable 
```bash
$GOPATH="/home/MyGoFolderProjects"
```
Then:

```bash
$ cd /home/MyGoFolderProject
$ git clone https://github.com/Chiheb-Nexus/blockr.io
$ cd Blockr.io/main/
$ touch main.go
```
For better use of this wrapper, you have to know two things:

1. Functions accept up to two arguments

2. Data returned from the functions are []TypeOfStructUsed

# List of features

1. [blockrio.GetCoinInfo(coin string)](http://btc.blockr.io/api/v1/coin/info)

2. [blockrio.GetExchangeCurrent(coin string)](http://btc.blockr.io/api/v1/exchangerate/current)

3. [blockrio.GetBlockInfo(coin string, block_number []string)](http://btc.blockr.io/api/v1/block/info/12345)

4. [blockrio.GetBlockTxs(coin string, block_number []string)](http://btc.blockr.io/api/v1/block/txs/last)

5. [blockrio.GetTxsInfo(coin string, address []string)](http://btc.blockr.io/api/v1/tx/info/60c1f1a3160042152114e2bba45600a5045711c3a8a458016248acec59653471)

6. [blockrio.GetBlockRaw(coin string, block_number []string)](http://btc.blockr.io/api/v1/block/raw/last)

7. [blockrio.GetAddressInfo(coin string, address []string)](http://btc.blockr.io/api/v1/address/info/198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi)

8. [blockrio.GetAddressBalance(coin string, address []string)](http://btc.blockr.io/api/v1/address/balance/198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi)

9. [blockrio.GetAddressTransactions(coin string, address []string)](http://btc.blockr.io/api/v1/address/txs/198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi,1L8meqhMTRpxasdGt8DHSJfscxgHHzvPgk)

10. [blockrio.GetAddressUnspent(coin string, address []string)](http://btc.blockr.io/api/v1/address/unspent/198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi)

11. [blockrio.GetUnconfirmedTxBalance(coin string, address []string)](http://btc.blockr.io/api/v1/address/unconfirmed/198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi) 

# Example:

This is an example of the file main.go

```go
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
	address_balance := blockrio.GetAddressBalance(coin, add)
	for _, value := range address_balance {
		for _, val := range value.Data {
			fmt.Printf("Address: %s  |  Balance: %v\n", val.Address, val.Balance)
		}
	}
}
		
		
```







