package blockrio

type ResponseUnconfirmedTx struct {
	Status  string       `json:"status"`
	Data    []DataUnconf `json:"data"`
	Code    float64      `json:"code"`
	Message string       `json:"message"`
}

type DataUnconf struct {
	Address     string        `json:"address"`
	Unconfirmed []Unconfirmed `json:"unconfirmed"`
}

type Unconfirmed struct {
	Tx     string  `json:"tx"`
	Time   string  `json:"time_utc"`
	Amount float64 `json:"amount"`
	N      float64 `json:"n"`
}
