package blockrio

// ResponseAddressUnspent: Return unspent coin from address
type ResponseAddressUnspent struct {
	Status  string        `json:"status"`
	Data    []DataUnspent `json:"data"`
	Code    float64       `json:"code"`
	Message string        `json:"message"`
}

// DataUnuspent: Inside ResponseAddressUnspent struct
type DataUnspent struct {
	Address string     `json:"address"`
	Unspent []_Unspent `json:"unspent"`
}

// _Unspent: Inside DataUnspent struct
type _Unspent struct {
	Tx            string  `json:"tx"`
	Amount        string  `json:"amount"`
	N             float64 `json:"n"`
	Confirmations float64 `json:"confirmations"`
	Script        string  `json:"script"`
}
