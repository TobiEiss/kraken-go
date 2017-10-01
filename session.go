package krakenGo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// SessionContextKey is the key-type vor the context
type SessionContextKey string

const (
	// HostSessionContextKey represent the key for the host
	HostSessionContextKey = SessionContextKey("host")

	RouteServerTime = "public/Time"
	RouteAssets     = "public/Assets"
	RouteAssetPairs = "public/AssetPairs"
	RouteTickerInfo = "public/Ticker"

	RouteAccountBalance = "private/Balance"
	RouteOpenOrders     = "private/OpenOrders"
	RouteClosedOrders   = "private/ClosedOrders"
	RouteTradesHistory  = "private/TradesHistory"
)

// SessionContext represent a kraken session
type SessionContext struct {
	context.Context
	Host      string
	Version   string
	APIKey    string
	APISecret string
}

// CreateKrakenSession creates a session with host-kraken
func CreateKrakenSession() *SessionContext {
	krakenHost := "https://api.kraken.com"
	krakenVersion := "/0/"
	return CreateSession(krakenHost, krakenVersion)
}

// CreateSession and set your own host. For example for your tests
func CreateSession(host string, version string) *SessionContext {
	return &SessionContext{
		Context: context.Background(),
		Host:    host,
		Version: version,
	}
}

// UsePrivateAPI for private-API you need the API-Key and the API-Secret
func (session *SessionContext) UsePrivateAPI(apiKey string, apiSecret string) {
	session.APIKey = apiKey
	session.APISecret = apiSecret
}

// KrakenResponse is the response from kraken.com
type KrakenResponse struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}

// HTTPDo function runs the HTTP request and processes its response in a new goroutine.
func HTTPDo(ctx context.Context, request *http.Request, processResponse func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to processResponse.
	transport := &http.Transport{}
	client := &http.Client{Transport: transport}
	errorChannel := make(chan error, 1)

	// do request
	go func() { errorChannel <- processResponse(client.Do(request)) }()
	select {
	case <-ctx.Done():
		transport.CancelRequest(request)
		<-errorChannel // wait for processResponse function
		return ctx.Err()
	case err := <-errorChannel:
		return err
	}
}

// query the api
func (session *SessionContext) query(typ interface{}, route string, values url.Values, header map[string]string) error {
	var krakenResponse KrakenResponse
	krakenResponse.Result = typ

	// create httpURL
	httpURL := session.Host + session.Version + route

	// create http-Context
	httpContext, cancelFunc := context.WithTimeout(session, 15*time.Second)
	defer cancelFunc()

	// build request
	request, err := func() (*http.Request, error) {
		if values != nil {
			return http.NewRequest("POST", httpURL, strings.NewReader(values.Encode()))
		}
		return http.NewRequest("GET", httpURL, nil)
	}()
	if err != nil {
		return err
	}

	// add header if necessary
	for key, value := range header {
		request.Header.Add(key, value)
	}

	// fire up request and unmarshal serverTime
	err = HTTPDo(httpContext, request, func(response *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer response.Body.Close()
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		if err := decoder.Decode(&krakenResponse); err != nil {
			return err
		}
		return nil
	})
	return err
}
