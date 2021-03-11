
## Crypto L2 book consolidation (real-time data)

One of the API’s that Polygon.io provides is a WebSocket stream to Level 2 book data for crypto currencies pair(s). Each event emitted from the API represents the L2 data for a single currency pair for a single crypto exchange. Polygon.io supports several crypto exchanges such as Coinbase, Binance, and Bitstamp to name a few.

The task for this problem is to construct a Level 2 book that takes into account all the different exchange listings. This consolidated book would illustrate the availability of orders for each quoted price.

Here is an example output from the WebSocket API for the crypto pair BTC-USD:

Coinbase(1): [Bids: [33712.7, 0.18635], Asks: [33718.23, 3.5527483]]
Binance(2): [Bids: [33712.7, 0.134], Asks: [33718.23, 0.1]]

The desired output from the program you write would be:


Bids: [{price: 33712.7, available: [
{ exchangeID: 1, amount: 0.18635 }, { exchangeID: 2, amount: 0.134 }]
}]

Asks: [{price: 33718.23, available: [
{ exchangeID: 1, amount: 3.5527483 }, { exchangeID: 2, amount: 0.1 }
]}]

### Run / Build Binary

- go build && ./Polygon.io btc-usd

### File structure

- connection.go - Fetches, proccesses and prints data from websocket to CLI

- orders.go - Host structs for data pulled from websocket

- exchanges.go - Connects to seperate API to fetch Exchanges and writes to in file structs

### Test
- go test