package krakenGo

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"net/url"
	"strconv"
	"time"
)

// See https://www.kraken.com/help/api#general-usage for more information
func (session *SessionContext) queryPrivate(typ interface{}, route string, values url.Values) error {
	httpURL := session.Version + route

	if values == nil {
		values = url.Values{}
	}

	secret, _ := base64.StdEncoding.DecodeString(session.APISecret)
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	values.Set("nonce", nonce)

	// shaSum with nonce and values
	sha := sha256.New()
	sha.Write([]byte(nonce + values.Encode()))
	shaSum := sha.Sum(nil)

	// calc hmac
	mac := hmac.New(sha512.New, secret)
	mac.Write(append([]byte(httpURL), shaSum...))
	macSum := mac.Sum(nil)

	// create signature
	signature := string(base64.StdEncoding.EncodeToString(macSum))

	// add to header
	header := map[string]string{"API-Sign": signature, "API-Key": session.APIKey}

	// do request
	return session.query(typ, route, values, header)
}

// AccountBalance array of asset names and balance amount
func (session *SessionContext) AccountBalance() (map[Currency]string, error) {
	var balance map[Currency]string
	err := session.queryPrivate(&balance, RouteAccountBalance, nil)
	return balance, err
}