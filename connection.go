package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	socket "github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var File *os.File

func init () {
	createOrderBookCSV("orderBook.csv")
}

func ConnectToEndpoint(api_key string) {
	conn, _, err := socket.DefaultDialer.Dial(API_URL, nil)

	if err != nil { log.Fatalln("Failed to connect to socket endpoint") }
	defer conn.Close()

	conn.WriteMessage(socket.TextMessage, QueryInBytes(api_key, true))
	conn.WriteMessage(socket.TextMessage, QueryInBytes(CURRENCY_PAIR_ARG, false))

	responseChan := make(chan []byte, 1000)
	
	for {
		_, response, err := conn.ReadMessage()
		if err != nil { log.Fatalln(err) }

		responseChan <- response
		go parseResponse(responseChan)
	}
}

// Parses response from socket to JSON and writes to struct 
func parseResponse(responseChan chan[]byte){
	for respBytes := range responseChan { // Wait for response on chan..
		parsedResponse := &FetchedCurrencyPair{}
		err := json.Unmarshal(respBytes, parsedResponse)
		if err != nil { log.Fatalln(err) }

		go organiseDataFromResp(*parsedResponse)
	}
}

// Processes data received into seperate structs for asks and bids
func organiseDataFromResp(data FetchedCurrencyPair) {
	for _, order := range data {
		if len(order.Bids) > 0 {
			// Organise order data into separate entities.
			bidPrice := order.Bids[0][0]
			bid := OrderData{order.ExchangeID, order.Bids[0][1], ResolveExchangeIDtoName(order.ExchangeID)}

			askPrice := order.Asks[0][0]
			ask := OrderData{order.ExchangeID, order.Asks[0][1], ResolveExchangeIDtoName(order.ExchangeID)}
			
			bidOutput := OrderOutput{
				orderType: "BID",
				Price: bidPrice,
				Available: bid,
			}

			askOutput := OrderOutput{ 
				orderType: "ASK",
				Price: askPrice,
				Available: ask,
			}

			go sendToOutput(askOutput)
			go sendToOutput(bidOutput)
			go writeToCSV(askOutput, File)
			go writeToCSV(bidOutput, File)
		}	
	}	
}

// Prints output to CLI
func sendToOutput(order OrderOutput) {
	logrus.SetFormatter(&logrus.JSONFormatter{})

    logrus.WithFields(logrus.Fields{
		"ExchangeID": order.Available.ExchangeID,
		"Price": order.Price,
		"Type": order.orderType,
		"Amount": order.Available.Amount, 
		"ExchangeName": order.Available.ExchangeName,
	}).Infoln("OK")
}

// Automates query and returns as bytes to be plugged into socket
func QueryInBytes(parameter string, isAuth bool) []byte {
	if isAuth == true {
		return []byte(fmt.Sprintf("{\"action\":\"auth\",\"params\":\"%s\"}", parameter))	
	}

	return []byte(fmt.Sprintf("{\"action\":\"subscribe\",\"params\":\"XL2.%s\"}", parameter))
}
// ExchangeID, Price, OrderType, Amount, ExchangeName
func writeToCSV(order OrderOutput, file *os.File) {
	exchangeID := strconv.Itoa(order.Available.ExchangeID)
	price := strconv.Itoa(int(order.Price))
	amount := fmt.Sprintf("%f", order.Available.Amount)

	data := []string{
		exchangeID,
		price,
		order.orderType,
		amount,
		order.Available.ExchangeName,
	} 

	writer := csv.NewWriter(file)
	defer writer.Flush()

	errW := writer.Write(data)
	
	if errW != nil { log.Println("Failed to write orders to orderBook.csv")}
}

func createOrderBookCSV(fName string) {
	f, _ := os.Create(fName)
	File = f
}