package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	API_KEY              string
	API_URL              string
	EXCHANGE_API_URL     string
	CURRENCY_PAIR_ARG    string
	
	Crypto_Exchanges 	 CryptoExchanges // Keeps track of current exchanges from API
)

func init () {
	// Load env variables
	err := godotenv.Load()
	if err != nil { log.Fatalln("Failed to load API info from environment")}

	API_KEY 		 = os.Getenv("API_KEY") 	
	API_URL 		 = os.Getenv("API_URL")
	EXCHANGE_API_URL = os.Getenv("EXCHANGE_API_URL")

	go parseCLIArgs()
}

func main() {
	go GetExchangesFromEndpoint(API_KEY, EXCHANGE_API_URL)
	ConnectToEndpoint(API_KEY)
}

func parseCLIArgs() {
	if len(os.Args) < 2 {
		fmt.Println(`
		Usage:   go run *.go <crypto pair>

		Example: go run *.go btc-usd
		`)
		os.Exit(1)
	}

	args := os.Args[1]
	CURRENCY_PAIR_ARG = strings.ToUpper(args)
}



