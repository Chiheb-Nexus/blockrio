package Blockrio

/*
{
"status":"success",
"data":{
  "address":"198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi",
  "is_unknown":false,
  "balance":8000.00236957,
  "balance_multisig":0,
  "totalreceived":8000.00236957,
  "nb_txs":33,
  "first_tx":{
    "time_utc":"2009-02-22T10:50:53Z",
    "tx":"0f0fbcc18fd0d090ad3402574df8404cec1176bc000f9aa0dc19f8d832ff94db",
    "block_nb":"5219",
    "value":400,
    "confirmations":432363
    },
    "last_tx":{
      "time_utc":"2016-08-11T16:33:24Z",
      "tx":"c653ac1e7a66117e097dd16cfd122b24b0f92060d096d8754a1e550d4c64f520",
      "block_nb":"424718",
      "value":0.0002,
      "confirmations":12864
    },
    "is_valid":true
  },
  "code":200,
  "message":""
}
*/

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
