package krakenGo_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/TobiEiss/kraken-go"
)

// This test simulate a server witch sleeps for 110 millis.
// Also a request with a timeout of 100 millis.
// If the error is not nil, there is a failure!
func TestHttpDo(t *testing.T) {
	// simulate a server
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(110 * time.Millisecond)
		fmt.Fprintln(writer, "hello")
	}))
	defer testServer.Close()

	// build context with timeout of 100 millis
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// build request
	request, err := http.NewRequest("GET", testServer.URL, nil)
	if err != nil {
		t.Fail()
	}

	// do http-request
	err = krakenGo.HTTPDo(ctx, request, func(response *http.Response, err error) error {
		if err != nil {
			return err
		}
		return nil
	})

	// error have to be not nil cause of timeout
	log.Println(err.Error())
	if err == nil {
		t.Fail()
	}
}
