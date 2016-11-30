package Blockrio

type ResponseAddress struct {
	Status  string         `json:"status"`
	Data    []_DataAddress `json:"data"`
	Code    float64        `json:"code"`
	Message string         `json:"code"`
}

type _DataAddress struct {
	Address         string   `json:"address"`
	IsKnow          bool     `json:"is_unknown"`
	Balance         float64  `json:"balance"`
	BalanceMultiSig float64  `json:"balance_multisig"`
	TotalRecieved   float64  `json:"totalreceived"`
	NbTxs           float64  `json:"nb_txs"`
	FirstTx         _FirstTx `json:"first_tx"`
	LastTx          _Lasttx  `json:"last_tx"`
	IsValid         bool     `json:"is_valid"`
}

type _FirstTx struct {
	Time          string  `json:"time_utc"`
	Tx            string  `json:"tx"`
	BlockNb       string  `json:"block_nb"`
	Value         float64 `json:"value"`
	Confirmations float64 `json:"confirmations"`
}

type _Lasttx struct {
	Time          string  `json:"time_utc"`
	Tx            string  `json:"tx"`
	BlockNb       string  `json:"block_nb"`
	Value         float64 `json:"value"`
	Confirmations float64 `json:"confirmations"`
}
