package krakenGo

// GetServerTime returns serverTime
func (session *SessionContext) GetServerTime() (ServerTime, error) {
	servertime := ServerTime{}
	err := session.getHTTPDo(&servertime, RouteServerTime)
	return servertime, err
}

// GetAssetInfo returns the assets of kraken
func (session *SessionContext) GetAssetInfo() (map[string]Asset, error) {
	assets := map[string]Asset{}
	err := session.getHTTPDo(&assets, RouteAssets)
	return assets, err
}

// GetAssetPairs returns all tradeable asset-pairs
func (session *SessionContext) GetAssetPairs() (map[string]AssetPair, error) {
	assetpairs := map[string]AssetPair{}
	err := session.getHTTPDo(&assetpairs, RouteAssetPairs)
	return assetpairs, err
}
