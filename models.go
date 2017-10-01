package krakenGo

import (
	"strconv"
	"strings"
	"time"
)

type Currency string

// Currencies
const (
	BCH  Currency = "BCH"
	DASH Currency = "DASH"
	EOS  Currency = "EOS"
	GNO  Currency = "GNO"
	KFEE Currency = "KFEE"
	USDT Currency = "USDT"
	XDAO Currency = "XDAO"
	XETC Currency = "XETC"
	XETH Currency = "XETH"
	XICN Currency = "XICN"
	XLTC Currency = "XLTC"
	XMLN Currency = "XMLN"
	XNMC Currency = "XNMC"
	XREP Currency = "XREP"
	XXBT Currency = "XXBT"
	XXDG Currency = "XXDG"
	XXLM Currency = "XXLM"
	XXMR Currency = "XXMR"
	XXRP Currency = "XXRP"
	XXVN Currency = "XXVN"
	XZEC Currency = "XZEC"
	ZCAD Currency = "ZCAD"
	ZEUR Currency = "ZEUR"
	ZGBP Currency = "ZGBP"
	ZJPY Currency = "ZJPY"
	ZKRW Currency = "ZKRW"
	ZUSD Currency = "ZUSD"
)

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

// JSONTime represent the time-stuff
type JSONTime time.Time

// MarshalJSON overides the marshal methode of time
func (t JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// UnmarshalJSON overides the unmarshal methode of time
func (t *JSONTime) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	r = strings.Split(string(s), ".")[0]

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q, 0)
	return
}

// OpenOrders wraps orders
type OpenOrders struct {
	Orders map[string]Order `json:"open"`
}

// ClosedOrders wraps orders
type ClosedOrders struct {
	Orders map[string]Order `json:"closed"`
}

// Order represent an order
type Order struct {
	Refid    interface{} `json:"refid"`
	Userref  interface{} `json:"userref"`
	Status   string      `json:"status"`
	Opentm   JSONTime    `json:"opentm"`
	Starttm  JSONTime    `json:"starttm"`
	Expiretm JSONTime    `json:"expiretm"`
	Descr    struct {
		Pair      string `json:"pair"`
		Type      string `json:"type"`
		Ordertype string `json:"ordertype"`
		Price     string `json:"price"`
		Price2    string `json:"price2"`
		Leverage  string `json:"leverage"`
		Order     string `json:"order"`
	} `json:"descr"`
	Vol     string `json:"vol"`
	VolExec string `json:"vol_exec"`
	Cost    string `json:"cost"`
	Fee     string `json:"fee"`
	Price   string `json:"price"`
	Misc    string `json:"misc"`
	Oflags  string `json:"oflags"`
}

// Ledger contains all Ledger-Entries
type Ledger struct {
	LedgerEntries map[string]LedgerEntry `json:"ledger"`
}

// LedgerEntry represent a ledger entry
type LedgerEntry struct {
	Refid   string   `json:"refid"`
	Time    JSONTime `json:"time"`
	Type    string   `json:"type"`
	Aclass  string   `json:"aclass"`
	Asset   string   `json:"asset"`
	Amount  string   `json:"amount"`
	Fee     string   `json:"fee"`
	Balance string   `json:"balance"`
}

// Trades hold a map of id to trade
type Trades struct {
	Trades map[string]Trade `json:"trades"`
}

// Trade represent a trade
type Trade struct {
	Ordertxid string   `json:"ordertxid"`
	Pair      string   `json:"pair"`
	Time      JSONTime `json:"time"`
	Type      string   `json:"type"`
	Ordertype string   `json:"ordertype"`
	Price     string   `json:"price"`
	Cost      string   `json:"cost"`
	Fee       string   `json:"fee"`
	Vol       string   `json:"vol"`
	Margin    string   `json:"margin"`
	Misc      string   `json:"misc"`
}
