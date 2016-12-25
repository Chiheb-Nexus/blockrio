package blockrio

// ResponseAddressBalance: Return coin address balance
type ResponseAddressBalance struct {
	Status  string         `json:"status"`
	Data    []_DataBalance `json:"data"`
	Code    float64        `json:"code"`
	Message string         `json:"message"`
}

// _DataBalance: Inside ResponseAddressBalance struct
type _DataBalance struct {
	Address         string  `json:"address"`
	Balance         float64 `json:"balance"`
	BalanceMultiSig float64 `json:"balance_multisig"`
}
