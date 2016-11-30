package blockrio

import (
	"fmt"
	"strings"
)

// Return coin Info
func GetCoinInfo(coin string) ResponseInfo {
	resp := ResponseInfo{}

	switch coin {
	case "btc":
		url := "http://btc.blockr.io/api/v1/coin/info"
		resp = LoadInfo(url)
	case "ltc":
		url := "http://ltc.blockr.io/api/v1/coin/info"
		resp = LoadInfo(url)
	case "dgc":
		url := "http://dgc.blockr.io/api/v1/coin/info"
		resp = LoadInfo(url)
	}

	return resp
}

// Return Exchange coin prices
func GetExchangeCurrent(coin string) ResponseCurrent {
	resp := ResponseCurrent{}

	switch coin {
	case "btc":
		url := "http://btc.blockr.io/api/v1/exchangerate/current"
		resp = LoadCurrent(url)
	case "ltc":
		url := "http://ltc.blockr.io/api/v1/exchangerate/current"
		resp = LoadCurrent(url)
	case "dgc":
		url := "http://gdc.blockr.io/api/v1/exchangerate/current"
		resp = LoadCurrent(url)
	}

	return resp
}

// Return coin Block info
func GetBlockInfo(coin string, block_number []string) []ResponseBlock {
	resp := []ResponseBlock{}

	val := strings.Join(block_number, ",")
	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/block/info/%v", val)
		resp = append(resp, LoadBlock(url))
	case "ltc":
		url := fmt.Sprintf("http://ltc.blockr.io/api/v1/block/info/%v", val)
		resp = append(resp, LoadBlock(url))
	case "dgc":
		url := fmt.Sprintf("http://dgc.blockr.io/api/v1/block/info/%v", val)
		resp = append(resp, LoadBlock(url))
	}
	return resp
}

// Return Block Transactions
func GetBlockTxs(coin string, block_number []string) []ResponseBlockTransactions {
	resp := []ResponseBlockTransactions{}

	val := strings.Join(block_number, ",")
	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/block/txs/%v", val)
		resp = append(resp, LoadBlockTxs(url))
	case "ltc":
		url := fmt.Sprintf("http://ltc.blockr.io/api/v1/block/txs/%v", val)
		resp = append(resp, LoadBlockTxs(url))
	case "dgc":
		url := fmt.Sprintf("http://gdc.blockr.io/api/v1/block/txs/%v", val)
		resp = append(resp, LoadBlockTxs(url))
	}

	return resp
}

// Return Transactions Informations
func GetTxsInfo(coin string, block_number []string) []ResponseTransactionsInfo {
	resp := []ResponseTransactionsInfo{}

	val := strings.Join(block_number, ",")
	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/tx/info/%v", val)
		resp = append(resp, LoadTxsInfo(url))
	case "ltc":
		url := fmt.Sprintf("http://ltc.blockr.io/api/v1/tx/info/%v", val)
		resp = append(resp, LoadTxsInfo(url))
	case "dgc":
		url := fmt.Sprintf("http://gdc.blockr.io/api/v1/tx/info/%v", val)
		resp = append(resp, LoadTxsInfo(url))
	}

	return resp
}

// Load Transactions Block Raw
func GetBlockRaw(coin string, block_number []string) []ResponseBlockRaw {
	resp := []ResponseBlockRaw{}

	val := strings.Join(block_number, ",")

	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/block/raw/%v", val)
		resp = append(resp, LoadBlockRaw(url))
	case "ltc":
		url := fmt.Sprintf("http://ltc.blockr.io/api/v1/block/raw/%v", val)
		resp = append(resp, LoadBlockRaw(url))
	case "dgc":
		url := fmt.Sprintf("http://gdc.blockr.io/api/v1/block/raw/%v", val)
		resp = append(resp, LoadBlockRaw(url))
	}

	return resp
}

// Get Multiple Address Info
func GetAddressInfo(coin string, address []string) []ResponseAddress {
	resp := []ResponseAddress{}
	val := strings.Join(address, ",")

	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, LoadAddressInfo(url))
	case "ltc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, LoadAddressInfo(url))
	case "dgc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, LoadAddressInfo(url))
	}

	return resp
}

// Get Multiple Address Balance
func GetAddressBalance(coin string, address []string) []ResponseAddressBalance {
	resp := []ResponseAddressBalance{}
	val := strings.Join(address, ",")

	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/balance/%v", val)
		resp = append(resp, LoadAddressBalance(url))
	case "ltc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/balance/%v", val)
		resp = append(resp, LoadAddressBalance(url))
	case "dgc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/balance/%v", val)
		resp = append(resp, LoadAddressBalance(url))
	}

	return resp

}

// Get Multiple Address Transactions
func GetAddressTransactions(coin string, address []string) []ResponseAddressTransactions {
	resp := []ResponseAddressTransactions{}
	val := strings.Join(address, ",")

	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/txs/%v", val)
		resp = append(resp, LoadAddressTransactions(url))
	case "ltc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/txs/%v", val)
		resp = append(resp, LoadAddressTransactions(url))
	case "dgc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/txs/%v", val)
		resp = append(resp, LoadAddressTransactions(url))
	}

	return resp

}

// Get Address Unspent balance
func GetAddressUnspent(coin string, address []string) []ResponseAddressUnspent {
	resp := []ResponseAddressUnspent{}
	val := strings.Join(address, ",")

	switch coin {
	case "btc":
		url := fmt.Sprintf("http://blockr.io/api/v1/address/unspent/%v", val)
		resp = append(resp, LoadAddressUnspent(url))
	case "ltc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/unspent/%v", val)
		resp = append(resp, LoadAddressUnspent(url))
	case "dgc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/unspent/%v", val)
		resp = append(resp, LoadAddressUnspent(url))
	}

	return resp

}

// Get Inconfirmed Balance
func GetUnconfirmedTxBalance(coin string, address []string) []ResponseUnconfirmedTx {
	resp := []ResponseUnconfirmedTx{}
	val := strings.Join(address, ",")

	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/unconfirmed/%v", val)
		resp = append(resp, LoadUnconfirmedBalance(url))
	case "ltc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/unconfirmed/%v", val)
		resp = append(resp, LoadUnconfirmedBalance(url))
	case "dgc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/unconfirmed/%v", val)
		resp = append(resp, LoadUnconfirmedBalance(url))
	}

	return resp

}
