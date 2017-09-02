package krakenGo

// GetServerTime returns serverTime
func (session *SessionContext) GetServerTime() (ServerTime, error) {
	var servertimeWrapper struct {
		ServerTime ServerTime `json:"result,omitempty"`
	}
	err := session.getHTTPDo(&servertimeWrapper, RouteServerTime)

	return servertimeWrapper.ServerTime, err
}

// GetAssetInfo returns the assets of kraken
func (session *SessionContext) GetAssetInfo() (map[string]Asset, error) {
	var assetsWrapper struct {
		Assets map[string]Asset `json:"result,omitempty"`
	}
	err := session.getHTTPDo(&assetsWrapper, RouteAssets)
	return assetsWrapper.Assets, err
}

// GetAssetPairs returns all tradeable asset-pairs
func (session *SessionContext) GetAssetPairs() (map[string]AssetPair, error) {
	var assetpairsWrapper struct {
		AssetPair map[string]AssetPair `json:"result,omitempty"`
	}
	err := session.getHTTPDo(&assetpairsWrapper, RouteAssetPairs)
	return assetpairsWrapper.AssetPair, err
}
