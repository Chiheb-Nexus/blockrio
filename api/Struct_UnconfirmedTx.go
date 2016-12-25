package blockrio

// ResponseUnconfirmedTx: Return all uncorfimed transactions of coin's transactions
type ResponseUnconfirmedTx struct {
	Status  string       `json:"status"`
	Data    []DataUnconf `json:"data"`
	Code    float64      `json:"code"`
	Message string       `json:"message"`
}

// DataUnconf: Inside ResponseUnconfirmedTx struct
type DataUnconf struct {
	Address     string        `json:"address"`
	Unconfirmed []Unconfirmed `json:"unconfirmed"`
}

// Unconfirmed: Inside DataUnconf struct
type Unconfirmed struct {
	Tx     string  `json:"tx"`
	Time   string  `json:"time_utc"`
	Amount float64 `json:"amount"`
	N      float64 `json:"n"`
}
