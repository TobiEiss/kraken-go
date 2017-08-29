package krakenGo

import "context"
import "net/http"

// Session represent a kraken session
type Session struct {
}

// The httpDo function runs the HTTP request and processes its response in a new goroutine.
func httpDo(ctx context.Context, request *http.Request, processResponse func(*http.Response, error) error) error {
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
