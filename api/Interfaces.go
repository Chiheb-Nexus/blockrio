package Blockrio

import (
	"encoding/json"
	"log"
)

type BlockrioGetCoinInfo interface {
	LoadInfo(url string) ResponseInfo
}

// Load coin Info
func (res ResponseInfo) LoadInfo(url string) ResponseInfo {
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\n", err)
	}

	return res
}

type BlockrioGetExchangeCurrent interface {
	LoadCurrent(url string) ResponseCurrent
}

// Load Exchange Current prices
func (res ResponseCurrent) LoadCurrent(url string) ResponseCurrent {
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadCurrent", err)
	}

	return res
}

type BlockrioGetBlockInfo interface {
	LoadBlock(url string) ResponseBlock
}

// Load coin Block Info
func (res ResponseBlock) LoadBlock(url string) ResponseBlock {
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadBlock", err)
	}

	return res
}

type BlockrioGetBlockTxs interface {
	LoadBlockTxs(url string) ResponseBlockTransactions
}

// Load Block Transactions
func (res ResponseBlockTransactions) LoadBlockTxs(url string) ResponseBlockTransactions {
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed in !\nDebug: LoadBlockTxs", err)
	}

	return res
}

type BlockrioGetTxsInfo interface {
	LoadTxsInfo(url string) ResponseTransactionsInfo
}

// Load Transactions Informations
func (res ResponseTransactionsInfo) LoadTxsInfo(url string) ResponseTransactionsInfo {
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadTxsInfo", err)
	}

	return res
}
