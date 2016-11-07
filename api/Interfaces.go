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
		log.Fatal("Unmarchal failed in !\nDebug: LoadBlockTxs\n", err)
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
		log.Fatal("Unmarchal failed !\nDebug: LoadTxsInfo\n", err)
	}

	return res
}

type BlockrioGetBlockRaw interface {
	LoadBlockRaw(url string) ResponseBlockRaw
}

// Load Raws from a Block
func (res ResponseBlockRaw) LoadBlockRaw(url string) ResponseBlockRaw {
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadBlockRaw", err)
	}

	return res
}

type BlockrioGetAddressInfo interface {
	LoadAddressInfo(url string) ResponseAddress
}

// Load Multiple Address Info
func (res ResponseAddress) LoadAddressInfo(url string) ResponseAddress {
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressInfo\n", err)
	}

	return res
}

type BlockrioGetAddressBalance interface {
	LoadAddressBalance(url string) ResponseAddressBalance
}

// Load Multiple Address Balance
func (res ResponseAddressBalance) LoadAddressBalance(url string) ResponseAddressBalance {
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressBalance\n", err)
	}

	return res
}

type BlockrioGetAddressTransactions interface {
	LoadAddressTransactions(url string) ResponseAddressTransactionss
}

// Load Multiple Address Balance
func (res ResponseAddressTransactionss) LoadAddressTransactions(url string) ResponseAddressTransactionss {
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressTransactions\n", err)
	}

	return res
}

type BlockrioGetAddressUnspent interface {
	LoadAddressUnspent(url string) ResponseAddressUnspent
}

// Load Multiple Address Balance
func (res ResponseAddressUnspent) LoadAddressUnspent(url string) ResponseAddressUnspent {
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressTransactions\n", err)
	}

	return res
}

type BlockrioGetUnconfirmedBalance interface {
	LoadUnconfirmedBalance(url string) ResponseUnconfirmedTx
}

// Load Multiple Address Unconfirmed Balance
func (res ResponseUnconfirmedTx) LoadUnconfirmedBalance(url string) ResponseUnconfirmedTx {
	body := FetchUrlByte(url, GetUserAgent())
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Unmarchal failed !\nDebug: LoadAddressTransactions\n", err)
	}

	return res
}
