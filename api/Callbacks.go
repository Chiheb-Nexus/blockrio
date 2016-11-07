package Blockrio

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// Fetch the URL and return []byte
func FetchUrlByte(url string, user_agent string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal("Error while fetching url\n", err)
	}

	request.Header.Set("User-Agent", user_agent)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Error while trying to get response\n", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatal("Error status not OK!\n", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading body\n", err)
	}

	return body
}

// Use the best user-agent for the user
func GetUserAgent() string {
	var user_agent string

	switch runtime.GOOS {
	case "linux":
		user_agent = `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/53.0.2785.143 Chrome/53.0.2785.143 Safari/537.36`
	case "windows":
		user_agent = `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36`
	case "mac":
		user_agent = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36`
	}

	return user_agent
}

// Return coin Info
func GetCoinInfo(coin string) ResponseInfo {
	resp, i := ResponseInfo{}, ResponseInfo{}
	var get BlockrioGetCoinInfo
	get = i

	switch coin {
	case "btc":
		url := "http://btc.blockr.io/api/v1/coin/info"
		resp = get.LoadInfo(url)
	case "ltc":
		url := "http://ltc.blockr.io/api/v1/coin/info"
		resp = get.LoadInfo(url)
	case "dgc":
		url := "http://dgc.blockr.io/api/v1/coin/info"
		resp = get.LoadInfo(url)
	}

	return resp
}

// Return Exchange coin prices
func GetExchangeCurrent(coin string) ResponseCurrent {
	resp, i := ResponseCurrent{}, ResponseCurrent{}
	var get BlockrioGetExchangeCurrent
	get = i

	switch coin {
	case "btc":
		url := "http://btc.blockr.io/api/v1/exchangerate/current"
		resp = get.LoadCurrent(url)
	case "ltc":
		url := "http://ltc.blockr.io/api/v1/exchangerate/current"
		resp = get.LoadCurrent(url)
	case "dgc":
		url := "http://gdc.blockr.io/api/v1/exchangerate/current"
		resp = get.LoadCurrent(url)
	}

	return resp
}

// Return coin Block info
func GetBlockInfo(coin string, block_number interface{}) []ResponseBlock {
	resp := []ResponseBlock{}
	i := ResponseBlock{}
	var get BlockrioGetBlockInfo
	get = i

	value, err := block_number.([]string)
	if err != true {
		log.Fatal("Block number must be a []string\n", err)
	}

	val := strings.Join(value, ",")
	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/block/info/%v", val)
		resp = append(resp, get.LoadBlock(url))
	case "ltc":
		url := fmt.Sprintf("http://ltc.blockr.io/api/v1/block/info/%v", val)
		resp = append(resp, get.LoadBlock(url))
	case "dgc":
		url := fmt.Sprintf("http://dgc.blockr.io/api/v1/block/info/%v", val)
		resp = append(resp, get.LoadBlock(url))
	}
	return resp
}

// Return Block Transactions
func GetBlockTxs(coin string, block_number interface{}) []ResponseBlockTransactions {
	resp := []ResponseBlockTransactions{}
	i := ResponseBlockTransactions{}
	var get BlockrioGetBlockTxs
	get = i

	value, err := block_number.([]string)
	if err != true {
		log.Fatal("Block number must be a []string\n", err)
	}

	val := strings.Join(value, ",")
	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/block/txs/%v", val)
		resp = append(resp, get.LoadBlockTxs(url))
	case "ltc":
		url := fmt.Sprintf("http://ltc.blockr.io/api/v1/block/txs/%v", val)
		resp = append(resp, get.LoadBlockTxs(url))
	case "dgc":
		url := fmt.Sprintf("http://gdc.blockr.io/api/v1/block/txs/%v", val)
		resp = append(resp, get.LoadBlockTxs(url))
	}

	return resp
}

// Return Transactions Informations
func GetTxsInfo(coin string, block_number interface{}) []ResponseTransactionsInfo {
	resp := []ResponseTransactionsInfo{}
	i := ResponseTransactionsInfo{}
	var get BlockrioGetTxsInfo
	get = i

	value, err := block_number.([]string)
	if err != true {
		log.Fatal("Transactions number must be a []string\n", err)
	}

	val := strings.Join(value, ",")
	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/tx/info/%v", val)
		resp = append(resp, get.LoadTxsInfo(url))
	case "ltc":
		url := fmt.Sprintf("http://ltc.blockr.io/api/v1/tx/info/%v", val)
		resp = append(resp, get.LoadTxsInfo(url))
	case "dgc":
		url := fmt.Sprintf("http://gdc.blockr.io/api/v1/tx/info/%v", val)
		resp = append(resp, get.LoadTxsInfo(url))
	}

	return resp
}

// Load Transactions Block Raw
func GetBlockRaw(coin string, block_number interface{}) []ResponseBlockRaw {
	resp := []ResponseBlockRaw{}
	i := ResponseBlockRaw{}
	var get BlockrioGetBlockRaw
	get = i

	value, err := block_number.([]string)
	if err != true {
		log.Fatal("Block number must be a []string\n", err)
	}
	val := strings.Join(value, ",")

	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/block/raw/%v", val)
		resp = append(resp, get.LoadBlockRaw(url))
	case "ltc":
		url := fmt.Sprintf("http://ltc.blockr.io/api/v1/block/raw/%v", val)
		resp = append(resp, get.LoadBlockRaw(url))
	case "dgc":
		url := fmt.Sprintf("http://gdc.blockr.io/api/v1/block/raw/%v", val)
		resp = append(resp, get.LoadBlockRaw(url))
	}

	return resp
}

// Get Multiple Address Info
func GetAddressInfo(coin, address interface{}) []ResponseAddress {
	resp := []ResponseAddress{}
	i := ResponseAddress{}
	var get BlockrioGetAddressInfo
	get = i

	value, err := address.([]string)
	if !err {
		log.Fatal("Address must be a []string\n", err)
	}
	val := strings.Join(value, ",")

	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadAddressInfo(url))
	case "ltc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadAddressInfo(url))
	case "dgc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadAddressInfo(url))
	}

	return resp

}

// Get Multiple Address Balance
func GetAddressBalance(coin, address interface{}) []ResponseAddressBalance {
	resp := []ResponseAddressBalance{}
	i := ResponseAddressBalance{}
	var get BlockrioGetAddressBalance
	get = i

	value, err := address.([]string)
	if !err {
		log.Fatal("Address must be a []string\n", err)
	}
	val := strings.Join(value, ",")

	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadAddressBalance(url))
	case "ltc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadAddressBalance(url))
	case "dgc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadAddressBalance(url))
	}

	return resp

}

// Get Multiple Address Transactions
func GetAddressTransactions(coin, address interface{}) []ResponseAddressTransactionss {
	resp := []ResponseAddressTransactionss{}
	i := ResponseAddressTransactionss{}
	var get BlockrioGetAddressTransactions
	get = i

	value, err := address.([]string)
	if !err {
		log.Fatal("Address must be a []string\n", err)
	}
	val := strings.Join(value, ",")

	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadAddressTransactions(url))
	case "ltc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadAddressTransactions(url))
	case "dgc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadAddressTransactions(url))
	}

	return resp

}

// Get Address Unspent balance
func GetAddressUnspent(coin, address interface{}) []ResponseAddressUnspent {
	resp := []ResponseAddressUnspent{}
	i := ResponseAddressUnspent{}
	var get BlockrioGetAddressUnspent
	get = i

	value, err := address.([]string)
	if !err {
		log.Fatal("Address must be a []string\n", err)
	}
	val := strings.Join(value, ",")

	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadAddressUnspent(url))
	case "ltc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadAddressUnspent(url))
	case "dgc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadAddressUnspent(url))
	}

	return resp

}

// Get Inconfirmed Balance
func GetUnconfirmedBalance(coin, address interface{}) []ResponseUnconfirmedTx {
	resp := []ResponseUnconfirmedTx{}
	i := ResponseUnconfirmedTx{}
	var get BlockrioGetUnconfirmedBalance
	get = i

	value, err := address.([]string)
	if !err {
		log.Fatal("Address must be a []string\n", err)
	}
	val := strings.Join(value, ",")

	switch coin {
	case "btc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadUnconfirmedBalance(url))
	case "ltc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadUnconfirmedBalance(url))
	case "dgc":
		url := fmt.Sprintf("http://btc.blockr.io/api/v1/address/info/%v", val)
		resp = append(resp, get.LoadUnconfirmedBalance(url))
	}

	return resp

}
