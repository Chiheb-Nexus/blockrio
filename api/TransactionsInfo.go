package Blockrio

type ResponseTransactionsInfo struct {
	Status  string        `json:"status"`
	Data    []_DataTxInfo `json:"data"`
	Code    float64       `json:"code"`
	Message string        `json:"message"`
}

type _DataTxInfo struct {
	Tx            string        `json:"tx"`
	Block         float64       `json:"block"`
	Confirmations float64       `json:"confirmations"`
	Time          string        `json:"time_utc"`
	IsCoinbased   float64       `json:"is_coinbased"`
	Trade         _TradeInfo    `json:"trade"`
	Vins          []_VinsInfo2  `json:"vins"`
	Vouts         []_VoutsInfo2 `json:"vout"`
	Fee           string        `json:"fee"`
	DaysDestroyed string        `json:"days_destroyed"`
	IsConfirmed   bool          `json:"is_confirmed"`
	Extras        string        `json:"extras,omitempty"`
}

type _TradeInfo struct {
	Vins  []_VinsInfo  `json:"vins"`
	Vouts []_VoutsInfo `json:"vouts"`
}

type _VinsInfo struct {
	Address       string  `json:"address"`
	IsNonStandard bool    `json:"is_nonstandard"`
	Amount        float64 `json:"amount"`
	N             float64 `json:"n"`
	Type          float64 `json:"type"`
	VoutTx        string  `json:"vout_tx"`
}

type _VoutsInfo struct {
	Address       string  `json:"address"`
	IsNonStandard bool    `json:"is_nonstandard"`
	Amount        float64 `json:"amount"`
	Type          float64 `json:"type"`
	IsSpent       float64 `json:"is_spent"`
}

type _VinsInfo2 struct {
	Address       string  `json:"address"`
	IsNonStandard bool    `json:"is_nonstandard"`
	Amount        string  `json:"amount"`
	N             float64 `json:"n"`
	Type          float64 `json:"type"`
	VoutTx        string  `json:"vout_tx"`
}

type _VoutsInfo2 struct {
	Address       string  `json:"address"`
	IsNonStandard bool    `json:"is_nonstandard"`
	Amount        string  `json:"amount"`
	Type          float64 `json:"type"`
	IsSpent       float64 `json:"is_spent"`
}
