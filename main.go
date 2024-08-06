package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type CryptoResponse struct {
	Bitcoin  CryptoData `json:"bitcoin"`
	Ethereum CryptoData `json:"ethereum"`
}

type CryptoData struct {
	USD float64 `json:"usd"`
}

type ExchangeResponse struct {
	Rates struct {
		GBP float64 `json:"GBP"`
		KSH float64 `json:"KES"`
		USD float64 `json:"USD"`
	} `json:"rates"`
}

var (
	cryptoData   CryptoResponse
	exchangeData ExchangeResponse
	mu           sync.Mutex
)

func main() {
	// Start goroutines to fetch data
	go fetchCryptoDataPeriodically()
	go fetchExchangeDataPeriodically()

	// Create a ticker that triggers every second
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		displayData()
	}
}

func fetchCryptoDataPeriodically() {
	for {
		newData := getCryptoData()
		mu.Lock()
		cryptoData = newData
		mu.Unlock()
		time.Sleep(60 * time.Second) // Fetch every 10 seconds
	}
}

func fetchExchangeDataPeriodically() {
	for {
		newData := getExchangeData()
		mu.Lock()
		exchangeData = newData
		mu.Unlock()
		time.Sleep(60 * time.Second) // Fetch every 60 seconds
	}
}

func displayData() {
	mu.Lock()
	defer mu.Unlock()

	fmt.Printf("\033[2J") // Clear screen
	fmt.Printf("\033[H")  // Move cursor to top-left corner

	fmt.Println("Cryptocurrency Prices:")
	fmt.Printf("Bitcoin: $%.2f\n", cryptoData.Bitcoin.USD)
	fmt.Printf("Ethereum: $%.2f\n", cryptoData.Ethereum.USD)

	fmt.Println("\nExchange Rates:")
	fmt.Printf("GBP/USD: %.4f\n", exchangeData.Rates.USD/exchangeData.Rates.GBP)
	fmt.Printf("GBP/KSH: %.4f\n", exchangeData.Rates.KSH/exchangeData.Rates.GBP)
	fmt.Printf("USD/KSH: %.4f\n", exchangeData.Rates.KSH/exchangeData.Rates.USD)

	fmt.Printf("\nLast updated: %s\n", time.Now().Format("15:04:05"))
}

func getCryptoData() CryptoResponse {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd")
	if err != nil {
		fmt.Println("Error fetching crypto data:", err)
		return CryptoResponse{}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var cryptoResp CryptoResponse
	json.Unmarshal(body, &cryptoResp)
	return cryptoResp
}

func getExchangeData() ExchangeResponse {
	resp, err := http.Get("https://api.exchangerate-api.com/v4/latest/USD")
	if err != nil {
		fmt.Println("Error fetching exchange data:", err)
		return ExchangeResponse{}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var exchangeResp ExchangeResponse
	json.Unmarshal(body, &exchangeResp)
	return exchangeResp
}
