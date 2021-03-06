package blockrio

// ResponseAddressTransactions: Return all the transaction of a coin's address
type ResponseAddressTransactions struct {
	Status  string             `json:"status"`
	Data    []_DataAddressTxss `json:"data"`
	Code    int32              `json:"code"`
	Message string             `json:"message"`
}

// _DataAddressTxss: Inside ResponseAddressTransactions struct
type _DataAddressTxss struct {
	Address       string           `json:"address"`
	LimitTxs      int32            `json:"limit_txs"`
	NbTxs         int32            `json:"nb_txs"`
	NbTxDisplayed int32            `json:"nb_txs_displayed"`
	Txs           []_TxAddressTxss `json:"txs"`
}

// _TxAddressTxss: Inside _DataAddressTxss
type _TxAddressTxss struct {
	Tx             string  `json:"tx"`
	Time           string  `json:"time_utc"`
	Confirmations  int64   `json:"confirmations"`
	Amount         float64 `json:"amount"`
	AmountMultiSig float64 `json:"amount_multisig"`
}
