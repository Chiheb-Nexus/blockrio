package Blockrio

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

// Load coin Info
func LoadInfo(url string) ResponseInfo {
	res := ResponseInfo{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadInfo", err)
	}

	return res
}

// Load Exchange Current prices
func LoadCurrent(url string) ResponseCurrent {
	res := ResponseCurrent{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadCurrent", err)
	}

	return res
}

// Load coin Block Info
func LoadBlock(url string) ResponseBlock {
	res := ResponseBlock{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadBlock", err)
	}

	return res
}

// Load Block Transactions
func LoadBlockTxs(url string) ResponseBlockTransactions {
	res := ResponseBlockTransactions{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed in !\nDebug: LoadBlockTxs\n", err)
	}

	return res
}

// Load Transactions Informations
func LoadTxsInfo(url string) ResponseTransactionsInfo {
	res := ResponseTransactionsInfo{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadTxsInfo\n", err)
	}

	return res
}

// Load Raws from a Block
func LoadBlockRaw(url string) ResponseBlockRaw {
	res := ResponseBlockRaw{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadBlockRaw", err)
	}

	return res
}

// Load Multiple Address Info
func LoadAddressInfo(url string) ResponseAddress {
	res := ResponseAddress{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressInfo\n", err)
	}

	return res
}

// Load Multiple Address Balance
func LoadAddressBalance(url string) ResponseAddressBalance {
	res := ResponseAddressBalance{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressBalance\n", err)
	}

	return res
}

// Load Multiple Address Balance
func LoadAddressTransactions(url string) ResponseAddressTransactions {
	res := ResponseAddressTransactions{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressTransactions\n", err)
	}

	return res
}

// Load Multiple Address Balance
func LoadAddressUnspent(url string) ResponseAddressUnspent {
	res := ResponseAddressUnspent{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressTransactions\n", err)
	}

	return res
}

// Load Multiple Address Unconfirmed Balance
func LoadUnconfirmedBalance(url string) ResponseUnconfirmedTx {
	res := ResponseUnconfirmedTx{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressTransactions\n", err)
	}

	return res
}

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
