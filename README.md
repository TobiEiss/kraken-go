kraken-go is a library for the kraken.com - API.
Writen in golang.

[![Build Status](https://travis-ci.org/TobiEiss/kraken-go.svg?branch=master)](https://travis-ci.org/TobiEiss/kraken-go)
[![Coverage Status](https://coveralls.io/repos/github/TobiEiss/kraken-go/badge.svg?branch=master)](https://coveralls.io/github/TobiEiss/kraken-go?branch=master)


### Currently implemented:
#### Public:
:heavy_check_mark: Get server time  
:heavy_check_mark: Get asset info  
:heavy_check_mark: Get tradable asset pairs  
:heavy_check_mark: Get ticker information

#### Private:
:heavy_check_mark: Get Balance  

## Getting started

Install kraken-go:
`go get github.com/TobiEiss/kraken-go`

Write your first kraken-go-program:
```golang
package main

import (
	"log"

	"github.com/TobiEiss/kraken-go"
)

func main() {
	// new kraken.com - session
	session := krakenGo.CreateKrakenSession()

	// get server-time
	serverTime, err := session.GetServerTime()
	if err != nil {
		log.Println(err)
	}
	log.Printf("Unix-ServerTime on kraken.com: %d", serverTime.Unixtime)
}
```