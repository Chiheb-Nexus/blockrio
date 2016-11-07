package Blockrio

type ResponseAddressUnspent struct {
	Status  string        `json:"status"`
	Data    []DataUnspent `json:"data"`
	Code    float64       `json:"code"`
	Message string        `json:"message"`
}

type DataUnspent struct {
	More MoreUnspentInfo
	Less LessunspentInfo
}

type MoreUnspentInfo struct {
	Address string        `json:"address"`
	Unspent []MoreUnspent `json:"unspent"`
}

type LessunspentInfo struct {
	Address string        `json:"address"`
	Unspent []LessUnspent `json:"unspent"`
}

type MoreUnspent struct {
	Tx            string  `json:"tx"`
	Amount        float64 `json:"amount"`
	N             float64 `json:"n"`
	Confirmations float64 `json:"confirmations"`
	Script        string  `jsont:script`
}

type LessUnspent struct {
	Address string      `json:'address'`
	Unspent MoreUnspent `json:"unspent"`
}
