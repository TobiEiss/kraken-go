package krakenGo

import (
	"net/url"
	"strings"
)

// GetServerTime returns serverTime
func (session *SessionContext) GetServerTime() (ServerTime, error) {
	servertime := ServerTime{}
	err := session.query(&servertime, RouteServerTime, nil)
	return servertime, err
}

// GetAssetInfo returns the assets of kraken
func (session *SessionContext) GetAssetInfo() (map[string]Asset, error) {
	assets := map[string]Asset{}
	err := session.query(&assets, RouteAssets, nil)
	return assets, err
}

// GetAssetPairs returns all tradeable asset-pairs
func (session *SessionContext) GetAssetPairs() (map[string]AssetPair, error) {
	assetpairs := map[string]AssetPair{}
	err := session.query(&assetpairs, RouteAssetPairs, nil)
	return assetpairs, err
}

// GetTickerInfo returns tickerinfo.
func (session *SessionContext) GetTickerInfo(pairs ...string) (map[string]TickerInfo, error) {
	tickerInfo := map[string]TickerInfo{}
	err := session.query(
		&tickerInfo,
		RouteTickerInfo,
		url.Values{"pair": {strings.Join(pairs, ",")}})
	return tickerInfo, err
}
