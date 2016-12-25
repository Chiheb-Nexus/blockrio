package blockrio

// ResponseAddress: Return coin address informations
type ResponseAddress struct {
	Status  string         `json:"status"`
	Data    []_DataAddress `json:"data"`
	Code    float64        `json:"code"`
	Message string         `json:"code"`
}

// _DataAddress: Inside ResponseAddress struct
type _DataAddress struct {
	Address         string   `json:"address"`
	IsKnow          bool     `json:"is_unknown"`
	Balance         float64  `json:"balance"`
	BalanceMultiSig float64  `json:"balance_multisig"`
	TotalRecieved   float64  `json:"totalreceived"`
	NbTxs           float64  `json:"nb_txs"`
	FirstTx         _FirstTx `json:"first_tx"`
	LastTx          _LastTx  `json:"last_tx"`
	IsValid         bool     `json:"is_valid"`
}

// _FirstTx: Inside _DataAddress struct
type _FirstTx struct {
	Time          string  `json:"time_utc"`
	Tx            string  `json:"tx"`
	BlockNb       string  `json:"block_nb"`
	Value         float64 `json:"value"`
	Confirmations float64 `json:"confirmations"`
}

// _LastTx: Inside _DataAddress struct
type _LastTx struct {
	Time          string  `json:"time_utc"`
	Tx            string  `json:"tx"`
	BlockNb       string  `json:"block_nb"`
	Value         float64 `json:"value"`
	Confirmations float64 `json:"confirmations"`
}
