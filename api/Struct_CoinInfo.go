/*
	This api presents coin information:
		- coin - Basic coin information with coin name, abbreviation, logo and homepage URL.
		- volume - Volume information: how many coins are in supply and how many coins will there ever be.
		- last_block - Information about the last block in the longest chain.
		- next_difficulty - When will next difficulty be retargeted and how big it will probably be.
		- market - Current price value on markets.
*/

package blockrio

// ResponseInfo: status, data, code, message
type ResponseInfo struct {
	Status  string  `json:"status"`
	Data    Info    `json:"data"` // coin, volume, markets, last_block, next_difficulty, websocket
	Code    float64 `json:"code"`
	Message string  `json:"message,omitempty"`
}

// Info: Inside of ResponseInfo
type Info struct {
	Coin      _Coin      `json:"coin"`            // name, abbr, logo, homepage
	Volume    _Volume    `json:"volume"`          // current, all, perc
	Markets   _Markets   `json:"markets"`         // btce, coinbase
	LastBlock _LastBlock `json:"last_block"`      // nb, time_utc, nb_txs, fee, difficulty
	NextDiff  _NextDiff  `json:"next_difficulty"` // difficulty, retarget_in, retarget_block, perc
	WebSocket _WebSocket `json:"websocket"`       // ws_url, wss_url
}

// _Coin: Inside of Info struct
type _Coin struct {
	Name     string `json:"name"`
	Abbr     string `json:"abbr"`
	Logo     string `json:"logo"`
	HomePage string `json:"homepage,omitempty"`
}

// _Volume: Inside of Info struct
type _Volume struct {
	Current float64 `json:"current"`
	All     float64 `json:"all"`
	Perc    float64 `json:"perc"`
}

// _Markets: Inside of Info struct
type _Markets struct {
	Btce     _BtceInfo     `json:"btce"`     // name, last_update_utc, value, currency, daily_change
	Coinbase _CoinbaseInfo `json:"coinbase"` // name, last_update_utc, value, currency, daily_change
	Cryptsy  _CryptsyInfo  `json:"cryptsy"`
}

// _CryptsyInfo: Inside of Info struct
type _CryptsyInfo struct {
	Name        string     `json:"name"`
	LastUpdate  string     `json:"last_update_utc"`
	Value       float64    `json:"value"`
	Currency    string     `json:"currency"`
	DailyChange _CrypDaily `json:"daily_change,omitempty"` // value, perc, diff
}

// _CryptDaily: Inside of Info struct
type _CrypDaily struct {
	Value string  `json:"value"`
	Prec  float64 `json:"prec"`
	Diff  float64 `json:"diff"`
}

// _BtceInfo: Inside of Info struct
type _BtceInfo struct {
	Name        string     `json:"name"`
	LastUpdate  string     `json:"last_update_utc"`
	Value       float64    `json:"value"`
	Currency    string     `json:"currency"`
	DailyChange _BtceDaily `json:"daily_change,omitempty"` // value, perc, diff
}

// _BtceDaily: Inside of Info struct
type _BtceDaily struct {
	Value string  `json:"value"`
	Prec  float64 `json:"prec"`
	Diff  float64 `json:"diff"`
}

// _CoinbaseInfo: Inside of Info struct
type _CoinbaseInfo struct {
	Name        string         `json:"name"`
	LastUpdate  string         `json:"last_update_utc"`
	Value       float64        `json:"value"`
	Currency    string         `json:"currency"`
	DailyChange _CoinbaseDaily `json:"daily_change"` // value, perc, diff
}

// _CoinbaseDaily: Inside of Info struct
type _CoinbaseDaily struct {
	Value string  `json:"value"`
	Prec  float64 `json:"prec"`
	Diff  float64 `json:"diff"`
}

// _LastBlock: Inside of Info struct
type _LastBlock struct {
	Nb         float64 `json:"nb"`
	Time       string  `json:"time_utc"`
	NbTxs      float64 `json:"nb_txs"`
	Fee        string  `json:"fee"`
	Difficulty string  `json:"difficulty"`
}

// _NextDiff: Inside of Info struct
type _NextDiff struct {
	Difficulty    float64 `json:"difficulty"`
	RetargetIn    float64 `json:"retarget_in"`
	RetargetBlock float64 `json:"retarget_block"`
	Perc          float64 `json:"perc"`
}

// _WebSocket: Inside of Info struct
type _WebSocket struct {
	Wsurl  string `json:"ws_url"`
	WssUrl string `json:"wss_url"`
}
