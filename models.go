package krakenGo

// ServerTime represent the kraken-Server-Time
type ServerTime struct {
	Unixtime int `json:"unixtime,omitempty"`
}

// Asset represent an asset
// ========================
//
// altname = alternate name
// aclass = asset class
// decimals = scaling decimal places for record keeping
// display_decimals = scaling decimal places for output display
type Asset struct {
	Aclass          string `json:"aclass,omitempty"`
	Altname         string `json:"altname,omitempty"`
	Decimals        int    `json:"decimals,omitempty"`
	DisplayDecimals int    `json:"display_decimals,omitempty"`
}

// AssetPair are all tradeable asset pairs
// =======================================
//
// altname = alternate pair name
// aclass_base = asset class of base component
// base = asset id of base component
// aclass_quote = asset class of quote component
// quote = asset id of quote component
// lot = volume lot size
// pair_decimals = scaling decimal places for pair
// lot_decimals = scaling decimal places for volume
// lot_multiplier = amount to multiply lot volume by to get currency volume
// leverage_buy = array of leverage amounts available when buying
// leverage_sell = array of leverage amounts available when selling
// fees = fee schedule array in [volume, percent fee] tuples
// fees_maker = maker fee schedule array in [volume, percent fee] tuples (if on maker/taker)
// fee_volume_currency = volume discount currency
// margin_call = margin call level
// margin_stop = stop-out/liquidation margin level
type AssetPair struct {
	Altname           string        `json:"altname"`
	AclassBase        string        `json:"aclass_base"`
	Base              string        `json:"base"`
	AclassQuote       string        `json:"aclass_quote"`
	Quote             string        `json:"quote"`
	Lot               string        `json:"lot"`
	PairDecimals      int           `json:"pair_decimals"`
	LotDecimals       int           `json:"lot_decimals"`
	LotMultiplier     int           `json:"lot_multiplier"`
	LeverageBuy       []interface{} `json:"leverage_buy"`
	LeverageSell      []interface{} `json:"leverage_sell"`
	Fees              [][]float64   `json:"fees"`
	FeesMaker         [][]float64   `json:"fees_maker"`
	FeeVolumeCurrency string        `json:"fee_volume_currency"`
	MarginCall        int           `json:"margin_call"`
	MarginStop        int           `json:"margin_stop"`
}

// TickerInfo represent the ticker info
// ====================================
//
// a = ask array(<price>, <whole lot volume>, <lot volume>),
// b = bid array(<price>, <whole lot volume>, <lot volume>),
// c = last trade closed array(<price>, <lot volume>),
// v = volume array(<today>, <last 24 hours>),
// p = volume weighted average price array(<today>, <last 24 hours>),
// t = number of trades array(<today>, <last 24 hours>),
// l = low array(<today>, <last 24 hours>),
// h = high array(<today>, <last 24 hours>),
// o = today's opening price
type TickerInfo struct {
	A []string `json:"a"`
	B []string `json:"b"`
	C []string `json:"c"`
	V []string `json:"v"`
	P []string `json:"p"`
	T []int    `json:"t"`
	L []string `json:"l"`
	H []string `json:"h"`
	O string   `json:"o"`
}
