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

// This test simulate a server witch sleeps for x millis.
// Also a request with a timeout of y millis.
// If server-sleep longer than timeout, expects an error.
// If server-sleep is shorter than timeout, expects NO error.
func TestHttpDo(t *testing.T) {
	var testCases = []struct {
		ServerDuration time.Duration
		ClientTimeOut  time.Duration
	}{{
		ServerDuration: 110 * time.Millisecond,
		ClientTimeOut:  130 * time.Millisecond,
	}, {
		ServerDuration: 130 * time.Millisecond,
		ClientTimeOut:  110 * time.Millisecond,
	}}

	for _, testCase := range testCases {
		// simulate a server
		testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			time.Sleep(testCase.ServerDuration)
			fmt.Fprintln(writer, "hello")
		}))
		defer testServer.Close()

		// build context with timeout of 100 millis
		ctx, cancel := context.WithTimeout(context.Background(), testCase.ClientTimeOut)
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
		if (testCase.ServerDuration > testCase.ClientTimeOut && err == nil) ||
			(testCase.ServerDuration < testCase.ClientTimeOut && err != nil) {
			t.Fail()
		}
	}

}

func TestKrakenServerTime(t *testing.T) {
	session := krakenGo.CreateKrakenSession()
	serverTime, err := session.GetServerTime()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	if serverTime.Unixtime == 0 {
		t.Fail()
	}
	log.Println(serverTime)
}

func TestKrakenAssetInfo(t *testing.T) {
	session := krakenGo.CreateKrakenSession()
	assets, err := session.GetAssetInfo()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	if len(assets) == 0 {
		t.Fail()
	}
	log.Println(assets)
}

func TestKrakenAssetPairs(t *testing.T) {
	session := krakenGo.CreateKrakenSession()
	assetpairs, err := session.GetAssetPairs()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	if len(assetpairs) == 0 {
		t.Fail()
	}
}
