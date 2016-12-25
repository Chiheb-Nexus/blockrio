package blockrio

// ResponseBlock: Return coin block infirmations
type ResponseBlock struct {
	Status  string  `json:"status"`
	Data    []_Data `json:"data"`
	Code    float64 `json:"code"`
	Message string  `json:"message"`
}

// _Data: Inside ResponseBlock struct
type _Data struct {
	Nb            float64     `json:"nb"`
	Hash          string      `json:"hash"`
	Version       float64     `json:"version"`
	Confirmations float64     `json:"confirmations"`
	Time          string      `json:"time_utc"`
	NbTxs         float64     `json:"nb_txs"`
	MerkleRoot    string      `json:"merkleroot"`
	NextBlockNb   interface{} `json:"next_block_nb"`
	PrevBlockNb   float64     `json:"prev_block_nb"`
	NextBlockHash string      `json:"next_block_hash"`
	PrevBlockHash string      `json:"prev_block_hash"`
	Fee           string      `json:"fee"`
	VoutSum       float64     `json:"vout_sum"`
	Size          string      `json:"size"`
	Difficulty    float64     `json:"difficulty"`
	DaysDestroyed float64     `json:"days_destroyed"`
	Extras        bool        `json:"extras"`
}
