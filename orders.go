package main

// FetchedCurrencyPair from WebSocket
type FetchedCurrencyPair []struct {
	Bids 		[][]float64 `json:"b"`
	Asks 		[][]float64 `json:"a"`
	ExchangeID 	int 	    `json:"x"`
}

// Used to represent Bids and Asks
type OrderOutput struct {
	orderType 	string 
	Price 		float64   `json:"price"`
	Available 	OrderData `json:"Available"`
}

type OrderData struct {
	ExchangeID 		int     `json:"exchangeID"`
	Amount 			float64 `json:"amount"`
	ExchangeName 	string  `json:"exchangeName"`
}
