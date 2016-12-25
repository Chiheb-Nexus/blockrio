/*
	Shows current exchange rate that is used on blockr.io
*/
package blockrio

// ResponseCurrent: Return current values in JSON
type ResponseCurrent struct {
	Status  string   `json:"status"`
	Data    []Prices `json:"data"`
	Code    float64  `json:"code"`
	Message string   `json:"message"`
}

// Prices: Inside ResponseCurrent struct
type Prices struct {
	Base    string `json:"base"`
	Updated string `json:"updated_utc"`
	Rates   _Rates `json:"rates"`
}

// _Rates: Inside Prices struct
type _Rates struct {
	BTC string `json:"BTC"`
	EUR string `json:"EUR"`
	ZAR string `json:"ZAR"`
	THB string `json:"THB"`
	SGD string `json:"SGD"`
	PHP string `json:"PHP"`
	NZD string `json:"NZD"`
	MYR string `json:"MYR"`
	MXN string `json:"MXN"`
	KRW string `json:"KRW"`
	INR string `json:"INR"`
	ILS string `json:"ILS"`
	IDR string `json:"IDR"`
	HKD string `json:"HKD"`
	CNY string `json:"CNY"`
	CAD string `json:"CAD"`
	BRL string `json:"BRL"`
	AUD string `json:"AUD"`
	TRY string `json:"TRY"`
	RUB string `json:"RUB"`
	HRK string `json:"HRK"`
	NOK string `json:"NOK"`
	CHF string `json:"CHF"`
	SEK string `json:"SEk"`
	RON string `json:"RON"`
	PLN string `json:"PLN"`
	HUF string `json:"HUF"`
	GBP string `json:"GBP"`
	DKK string `json:"DKK"`
	CZK string `json:"CZK"`
	BGN string `json:"BGN"`
	JPY string `json:"JPY"`
	USD string `json:"USD"`
}
