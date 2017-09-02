package krakenGo

// ServerTime represent the kraken-Server-Time
type ServerTime struct {
	Unixtime int `json:"unixtime,omitempty"`
}

// Asset represent an asset
type Asset struct {
	Aclass          string `json:"aclass,omitempty"`
	Altname         string `json:"altname,omitempty"`
	Decimals        int    `json:"decimals,omitempty"`
	DisplayDecimals int    `json:"display_decimals,omitempty"`
}

// AssetPair are all tradeable asset pairs
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
