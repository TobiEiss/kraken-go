package krakenGo

import (
	"context"
	"encoding/json"
	"net/http"
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
)

// SessionContext represent a kraken session
type SessionContext struct {
	context.Context
}

// CreateKrakenSession creates a session with host-kraken
func CreateKrakenSession() *SessionContext {
	krakenHost := "https://api.kraken.com/0/"
	return CreateSession(krakenHost)
}

// CreateSession and set your own host. For example for your tests
func CreateSession(host string) *SessionContext {
	return &SessionContext{context.WithValue(context.Background(), HostSessionContextKey, host)}
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

// impl httpdo with GET-methode
func (session *SessionContext) getHTTPDo(wrapper interface{}, route string) error {
	// create http-Context
	httpContext, cancelFunc := context.WithTimeout(session, 15*time.Second)
	defer cancelFunc()

	// build request
	request, err := http.NewRequest("GET", session.Value(HostSessionContextKey).(string)+route, nil)
	if err != nil {
		return err
	}

	// fire up request and unmarshal serverTime
	err = HTTPDo(httpContext, request, func(response *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer response.Body.Close()
		if err := json.NewDecoder(response.Body).Decode(&wrapper); err != nil {
			return err
		}
		return nil
	})
	return err
}
