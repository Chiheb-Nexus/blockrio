package blockrio

import (
	"encoding/json"
	"log"
)

// LoadInfo: Load coin Info
func LoadInfo(url string) ResponseInfo {
	res := ResponseInfo{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadInfo", err)
	}

	return res
}

// LoadCurrent: Load Exchange Current prices
func LoadCurrent(url string) ResponseCurrent {
	res := ResponseCurrent{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadCurrent", err)
	}

	return res
}

// LoadBlock: Load coin Block Info
func LoadBlock(url string) ResponseBlock {
	res := ResponseBlock{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadBlock", err)
	}

	return res
}

// LoadBlockTxs: Load Block Transactions
func LoadBlockTxs(url string) ResponseBlockTransactions {
	res := ResponseBlockTransactions{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed in !\nDebug: LoadBlockTxs\n", err)
	}

	return res
}

// LoadTxsInfo: Load Transactions Informations
func LoadTxsInfo(url string) ResponseTransactionsInfo {
	res := ResponseTransactionsInfo{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadTxsInfo\n", err)
	}

	return res
}

// LoadBlockRaw: Load Raws from a Block
func LoadBlockRaw(url string) ResponseBlockRaw {
	res := ResponseBlockRaw{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadBlockRaw", err)
	}

	return res
}

// LoadAddressInfo: Load Multiple Address Info
func LoadAddressInfo(url string) ResponseAddress {
	res := ResponseAddress{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressInfo\n", err)
	}

	return res
}

// LoadAddressBalance: Load Multiple Address Balance
func LoadAddressBalance(url string) ResponseAddressBalance {
	res := ResponseAddressBalance{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressBalance\n", err)
	}

	return res
}

// LoadAddressTransactions: Load Multiple Address Balance
func LoadAddressTransactions(url string) ResponseAddressTransactions {
	res := ResponseAddressTransactions{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressTransactions\n", err)
	}

	return res
}

// LoadAddressUnspent: Load Multiple Address Balance
func LoadAddressUnspent(url string) ResponseAddressUnspent {
	res := ResponseAddressUnspent{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressTransactions\n", err)
	}

	return res
}

// LoadUnconfirmedBalance: Load Multiple Address Unconfirmed Balance
func LoadUnconfirmedBalance(url string) ResponseUnconfirmedTx {
	res := ResponseUnconfirmedTx{}
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressTransactions\n", err)
	}

	return res
}
