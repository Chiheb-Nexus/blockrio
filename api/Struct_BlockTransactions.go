package blockrio

// ResponseBlockTransactions: Return all the transactions of a coin's block
type ResponseBlockTransactions struct {
	Status  string    `json:"status"`
	Data    []_DataTx `json:"data"`
	Code    float64   `json:"code"`
	Message string    `json:"message"`
}

// _DataTx: Inside ResponseBlockTransactions struct
type _DataTx struct {
	Nb             float64 `json:"nb"`
	LimitTx        float64 `json:"limit_txs"`
	NbTxs          float64 `json:"nb_txs"`
	NbTxsDisplayed float64 `json:"nb_txs_displayed"`
	Txs            []_Txs  `json:"txs"`
}

// _Txs: Inside _DataTx struct
type _Txs struct {
	Tx            string  `json:"tx"`
	Trade         _Trade  `json:"trade"`
	DaysDestroyed float64 `json:"days_destroyed"`
	IsCoinbased   float64 `json:"is_coinbased"`
	Fee           string  `json:"fee"`
	Extras        bool    `json:"extras"`
}

// _Trade: Inside _Txs struct
type _Trade struct {
	Vins  []_Vins `json:"vins"`
	Vouts []_Vout `json:"vouts"`
}

// _Vins: Inside _Trade struct
type _Vins struct {
	Address       string  `json:"address"`
	IsNonStandard bool    `json:"is_nonstandard"`
	Amount        float64 `json:"amount"`
	N             float64 `json:"n"`
	Type          float64 `json:"type"`
	VoutTx        string  `json:"vout_tx"`
}

// _Vout: Inside _Trade struct
type _Vout struct {
	Address       string  `json:"address"`
	IsNonStandard bool    `json:"is_nonstandard"`
	Amount        float64 `json:"amount"`
	N             float64 `json:"n"`
	Type          float64 `json:"type"`
	IsSpent       float64 `json:"is_spent"`
}
