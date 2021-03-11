package main

import (
	"encoding/json"
	"io/ioutil" // Go 1.14 - deprecated after 1.15 I believe.
	"log"
	"net/http"
)

type exchange struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CryptoExchanges struct {
	list [] exchange
}

func GetExchangesFromEndpoint(API_KEY, endpoint_url string) {
	resp, err := http.Get(endpoint_url+API_KEY)
	if err != nil { log.Println(err) }
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	// Decode bytes response of crpyto exchange data into memory.
	cryptoExch := CryptoExchanges{}
	if err := json.Unmarshal(body, &cryptoExch.list); err != nil { log.Println(err) }

	writeExchangesToGlobal(cryptoExch)
}

func writeExchangesToGlobal(dataToWrite CryptoExchanges) {
	Crypto_Exchanges = dataToWrite
}

func ResolveExchangeIDtoName(id int) string {
	for _, exchange := range Crypto_Exchanges.list {
		if exchange.ID == id {
			return exchange.Name
		}
	}
	return ""
}
