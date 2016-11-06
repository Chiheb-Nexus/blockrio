package Blockrio

/*
{
	"status":"success",
	"data":{
		"tx":"60c1f1a3160042152114e2bba45600a5045711c3a8a458016248acec59653471",
		"block":233958,
		"confirmations":203589,
		"time_utc":"2013-04-30T20:58:41Z",
		"is_coinbased":0,
		"trade":{
			"vins":[
				{
					"address":"1CRZRBwfuwUaVSPJtd6DBuezbm7XPBHLa1",
					"is_nonstandard":false,
					"amount":-0.015,
					"n":1,
					"type":0,
					"vout_tx":"6b040cd7a4676b5c7b11f144e73c1958c177fcd79e934f6be8ce02c8cd12546d"
				}
			],
			"vouts":[
				{
					"address":"16T3RPZLmxtXQCgWi1S8kef5Ca6jqXhoeT",
					"is_nonstandard":false,
					"amount":0.015,
					"n":1,
					"type":1,
					"is_spent":1
				}
			]
		},
		"vins":[
			{
				"address":"1CRZRBwfuwUaVSPJtd6DBuezbm7XPBHLa1",
				"is_nonstandard":false,
				"amount":"-380.43749285",
				"n":1,
				"type":0,
				"vout_tx":"6b040cd7a4676b5c7b11f144e73c1958c177fcd79e934f6be8ce02c8cd12546d"
			}
		],
		"vouts":[
			{
				"address":"1CRZRBwfuwUaVSPJtd6DBuezbm7XPBHLa1",
				"is_nonstandard":false,
				"amount":"380.42249285",
				"n":0,
				"type":1,
				"is_spent":1
			},
			{
				"address":"16T3RPZLmxtXQCgWi1S8kef5Ca6jqXhoeT",
				"is_nonstandard":false,
				"amount":"0.01500000",
				"n":1,
				"type":1,
				"is_spent":1
			}
		],
		"fee":"0.00000000",
		"days_destroyed":"0.00",
		"is_unconfirmed":false,
		"extras":null
	},
	"code":200,
	"message":""
}

*/

type ResponseTransactionsInfo struct {
	Status  string      `json:"status"`
	Data    _DataTxInfo `json:"data"`
	Code    float64     `json:"code"`
	Message string      `json:"message"`
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
