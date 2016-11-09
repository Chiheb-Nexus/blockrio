/*
	This api presents coin information:
		- coin - Basic coin information with coin name, abbreviation, logo and homepage URL.
		- volume - Volume information: how many coins are in supply and how many coins will there ever be.
		- last_block - Information about the last block in the longest chain.
		- next_difficulty - When will next difficulty be retargeted and how big it will probably be.
		- market - Current price value on markets.
*/

package Blockrio

// status, data, code, message
type ResponseInfo struct {
	Status  string  `json:"status"`
	Data    Info    `json:"data"` // coin, volume, markets, last_block, next_difficulty, websocket
	Code    float64 `json:"code"`
	Message string  `json:"message,omitempty"`
}

type Info struct {
	Coin      _Coin      `json:"coin"`            // name, abbr, logo, homepage
	Volume    _Volume    `json:"volume"`          // current, all, perc
	Markets   _Markets   `json:"markets"`         // btce, coinbase
	LastBlock _LastBlock `json:"last_block"`      // nb, time_utc, nb_txs, fee, difficulty
	NextDiff  _NextDiff  `json:"next_difficulty"` // difficulty, retarget_in, retarget_block, perc
	WebSocket _WebSocket `json:"websocket"`       // ws_url, wss_url
}

type _Coin struct {
	Name     string `json:"name"`
	Abbr     string `json:"abbr"`
	Logo     string `json:"logo"`
	HomePage string `json:"homepage,omitempty"`
}

type _Volume struct {
	Current float64 `json:"current"`
	All     float64 `json:"all"`
	Perc    float64 `json:"perc"`
}

type _Markets struct {
	Btce     _BtceInfo     `json:"btce"`     // name, last_update_utc, value, currency, daily_change
	Coinbase _CoinbaseInfo `json:"coinbase"` // name, last_update_utc, value, currency, daily_change
	Cryptsy  _CryptsyInfo  `json:"cryptsy"`
}

type _CryptsyInfo struct {
	Name        string     `json:"name"`
	LastUpdate  string     `json:"last_update_utc"`
	Value       float64    `json:"value"`
	Currency    string     `json:"currency"`
	DailyChange _CrypDaily `json:"daily_change,omitempty"` // value, perc, diff
}

type _CrypDaily struct {
	Value string  `json:"value"`
	Prec  float64 `json:"prec"`
	Diff  float64 `json:"diff"`
}

type _BtceInfo struct {
	Name        string     `json:"name"`
	LastUpdate  string     `json:"last_update_utc"`
	Value       float64    `json:"value"`
	Currency    string     `json:"currency"`
	DailyChange _BtceDaily `json:"daily_change,omitempty"` // value, perc, diff
}

type _BtceDaily struct {
	Value string  `json:"value"`
	Prec  float64 `json:"prec"`
	Diff  float64 `json:"diff"`
}

type _CoinbaseInfo struct {
	Name        string         `json:"name"`
	LastUpdate  string         `json:"last_update_utc"`
	Value       float64        `json:"value"`
	Currency    string         `json:"currency"`
	DailyChange _CoinbaseDaily `json:"daily_change"` // value, perc, diff
}

type _CoinbaseDaily struct {
	Value string  `json:"value"`
	Prec  float64 `json:"prec"`
	Diff  float64 `json:"diff"`
}

type _LastBlock struct {
	Nb         float64 `json:"nb"`
	Time       string  `json:"time_utc"`
	NbTxs      float64 `json:"nb_txs"`
	Fee        string  `json:"fee"`
	Difficulty string  `json:"difficulty"`
}

type _NextDiff struct {
	Difficulty    float64 `json:"difficulty"`
	RetargetIn    float64 `json:"retarget_in"`
	RetargetBlock float64 `json:"retarget_block"`
	Perc          float64 `json:"perc"`
}

type _WebSocket struct {
	Wsurl  string `json:"ws_url"`
	WssUrl string `json:"wss_url"`
}
