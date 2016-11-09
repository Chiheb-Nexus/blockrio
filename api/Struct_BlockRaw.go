package Blockrio

type ResponseBlockRaw struct {
	Status  string     `json:"status"`
	Data    []_DataRow `json:"data"`
	Code    float64    `json:"code"`
	Message string     `json:"message"`
}

type _DataRow struct {
	Time              float64  `json:"time"`
	MedianTime        float64  `json:"mediantime"`
	Nonce             float64  `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	ChainWork         string   `json:"chainwork"`
	PreviousBlockHach string   `json:"previousblockhash"`
	Hash              string   `json:"hash"`
	Confirmations     float64  `json:"confirmations"`
	Size              float64  `json:"size"`
	Height            float64  `json:"height"`
	Version           float64  `json:"Version"`
	MerkleRoot        string   `json:"merkleroot"`
	Tx                []string `json:"tx"`
}
