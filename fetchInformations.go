package krakenGo

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// GetServerTime returns serverTime
func (session SessionContext) GetServerTime() (ServerTime, error) {
	var servertime struct {
		Time ServerTime `json:"result,omitempty"`
	}

	// create http-Context
	httpContext, cancelFunc := context.WithTimeout(session, 15*time.Second)
	defer cancelFunc()

	// build request
	request, err := http.NewRequest("GET", session.Value(HostSessionContextKey).(string)+RouteServerTime, nil)
	if err != nil {
		return servertime.Time, err
	}

	// fire up request and unmarshal serverTime
	err = HTTPDo(httpContext, request, func(response *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer response.Body.Close()
		if err := json.NewDecoder(response.Body).Decode(&servertime); err != nil {
			return err
		}
		return nil
	})
	return servertime.Time, err
}
