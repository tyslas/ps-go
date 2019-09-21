package main

import (
	"net/http" //make http calls
	"io/ioutil" //read responses
	"encoding/xml" //unmarshall response body into xml
	"fmt"
	"time"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(4)

	start := time.Now()

	stockSymbols := []string {
		"gib",
		"ctsh",
		"vmw",
		"googl",
		"msft",
		"aapl",
		"hpq",
		"twlo",
	}

	numComplete := 0

	for _, symbol := range stockSymbols {
		go func(symbol string) {
			resp, _ := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + symbol)
			defer resp.Body.Close() //close connection
			body, _ := ioutil.ReadAll(resp.Body) //read response
	
			quote := new(QuoteResponse) //create reponse object
			xml.Unmarshal(body, &quote) //assign reponse to object
	
			fmt.Printf("%s: $%.2f\n", quote.Name, quote.LastPrice) //print to console
			numComplete++
		}(symbol)
	}

	for numComplete < len(stockSymbols) {
		 time.Sleep(10 * time.Millisecond)
	}

	elaspedTime := time.Since(start)
	fmt.Printf("\nexecution time: %s\n ", elaspedTime)
}

type QuoteResponse struct {
	Status string
	Name string
	LastPrice float32
	Change float32
	ChangePercent float32
	TimeStamp string
	MSDate float32
	MarketCap int
	Volume int
	ChangeYTD float32
	ChangePercentYTD float32
	High float32
	Low float32
	Open float32
}