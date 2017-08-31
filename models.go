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
