package Blockrio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
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

// Determine the best user-agent for the user
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

// Load coin Info
func LoadInfo(url string) ResponseInfo {
	body := FetchUrlByte(url, GetUserAgent())
	res := ResponseInfo{}
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\n", err)
	}

	return res
}

// Return coin Info for respective coin
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

// Load Exchange Current prices
func LoadCurrent(url string) ResponseCurrent {
	body := FetchUrlByte(url, GetUserAgent())
	res := ResponseCurrent{}
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\n", err)
	}

	return res
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

// Load coin Block Info
func LoadBlock(url string) ResponseBlock {
	body := FetchUrlByte(url, GetUserAgent())
	res := ResponseBlock{}
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\n", err)
	}

	return res
}

// Return coin Block info
func GetBlockInfo(coin string, block_number interface{}) []ResponseBlock {
	resp := []ResponseBlock{}

	switch block_number.(type) {
	default:
		log.Fatalf("Expected block_number as int64, got %T\n", block_number)
	case int:
		switch coin {
		case "btc":
			url := fmt.Sprintf("http://btc.blockr.io/api/v1/block/info/%d", block_number)
			resp = append(resp, LoadBlock(url))
		case "ltc":
			url := fmt.Sprintf("http://ltc.blockr.io/api/v1/block/info/%d", block_number)
			resp = append(resp, LoadBlock(url))
		case "dgc":
			url := fmt.Sprintf("http://dgc.blockr.io/api/v1/block/info/%d", block_number)
			resp = append(resp, LoadBlock(url))
		}
	case string:
		if block_number == "last" {
			switch coin {
			case "btc":
				url := fmt.Sprintf("http://btc.blockr.io/api/v1/block/info/%v", block_number)
				resp = append(resp, LoadBlock(url))
			case "ltc":
				url := fmt.Sprintf("http://ltc.blockr.io/api/v1/block/info/%v", block_number)
				resp = append(resp, LoadBlock(url))
			case "dgc":
				url := fmt.Sprintf("http://dgc.blockr.io/api/v1/block/info/%v", block_number)
				resp = append(resp, LoadBlock(url))
			}
		} else {
			log.Fatal("API accept int or 'last' as parameters\n")
		}
	case []string:
		value := block_number.([]string)
		if len(value) >= 1 {

			switch coin {

			case "btc":
				var url string
				for _, val := range value {
					url = fmt.Sprintf("http://btc.blockr.io/api/v1/block/info/%v", val)
					resp = append(resp, LoadBlock(url))
				}

			case "ltc":
				var url string
				for _, val := range value {
					url = fmt.Sprintf("http://ltc.blockr.io/api/v1/block/info/%v", val)
					resp = append(resp, LoadBlock(url))
				}
			case "dgc":
				var url string
				for _, val := range value {
					url = fmt.Sprintf("http://dgc.blockr.io/api/v1/block/info/%v", val)
					resp = append(resp, LoadBlock(url))
				}

			}
		}

	}

	return resp
}

// Load Block Transactions
func LoadBlockTxs(url string) ResponseBlockTransactions {
	body := FetchUrlByte(url, GetUserAgent())
	res := ResponseBlockTransactions{}
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\n", err)
	}

	return res
}

// Return Block Transactions
func GetBlockTxs(coin string, block_number interface{}) []ResponseBlockTransactions {
	resp := []ResponseBlockTransactions{}
	value, err := block_number.([]string)
	if err != true {
		log.Fatal("Block number must be a []string\n", err)
	}

	for _, val := range value {
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
	}

	return resp
}

// Load Transactions Informations
func LoadTxsInfo(url string) ResponseTransactionsInfo {
	body := FetchUrlByte(url, GetUserAgent())
	res := ResponseTransactionsInfo{}
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\n", err)
	}

	return res
}

// Return Transactions Informations
func GetTxsInfo(coin string, block_number interface{}) []ResponseTransactionsInfo {
	resp := []ResponseTransactionsInfo{}
	value, err := block_number.([]string)
	if err != true {
		log.Fatal("Transactions number must be a []string\n", err)
	}

	for _, val := range value {
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
	}

	return resp
}
