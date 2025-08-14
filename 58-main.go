package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const apiKey = "d2a37fe2-3152-4915-8f08-070acfbdd5e6" // Replace with your actual API key
const apiUrl = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest"

type Quote struct {
	Price float64 `json:"price"`
}

type CryptoData struct {
	Quote map[string]Quote `json:"quote"`
}

type Response struct {
	Data map[string]CryptoData `json:"data"`
}

func getTokenPrice(symbol string) (float64, error) {
	url := fmt.Sprintf("%s?symbol=%s", apiUrl, symbol)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, err
	}

	if data, exists := response.Data[symbol]; exists {
		return data.Quote["USD"].Price, nil
	}

	return 0, fmt.Errorf("the token %s not found", symbol)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <token_symbol>")
		return
	}

	symbol := os.Args[1]
	price, err := getTokenPrice(symbol)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The price of %s is $%.2f\n", symbol, price)
}
