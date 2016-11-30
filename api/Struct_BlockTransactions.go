package blockrio

type ResponseBlockTransactions struct {
	Status  string    `json:"status"`
	Data    []_DataTx `json:"data"`
	Code    float64   `json:"code"`
	Message string    `json:"message"`
}

type _DataTx struct {
	Nb             float64 `json:"nb"`
	LimitTx        float64 `json:"limit_txs"`
	NbTxs          float64 `json:"nb_txs"`
	NbTxsDisplayed float64 `json:"nb_txs_displayed"`
	Txs            []_Txs  `json:"txs"`
}

type _Txs struct {
	Tx            string  `json:"tx"`
	Trade         _Trade  `json:"trade"`
	DaysDestroyed float64 `json:"days_destroyed"`
	IsCoinbased   float64 `json:"is_coinbased"`
	Fee           string  `json:"fee"`
	Extras        bool    `json:"extras"`
}

type _Trade struct {
	Vins  []_Vins `json:"vins"`
	Vouts []_Vout `json:"vouts"`
}

type _Vins struct {
	Address       string  `json:"address"`
	IsNonStandard bool    `json:"is_nonstandard"`
	Amount        float64 `json:"amount"`
	N             float64 `json:"n"`
	Type          float64 `json:"type"`
	VoutTx        string  `json:vout_tx"`
}

type _Vout struct {
	Address       string  `json:"address"`
	IsNonStandard bool    `json:"is_nonstandard"`
	Amount        float64 `json:"amount"`
	N             float64 `json:"n"`
	Type          float64 `json:"type"`
	IsSpent       float64 `json:"is_spent"`
}
