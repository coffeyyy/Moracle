package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)


func GetCMCPrice() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cmc_api_key := os.Getenv("CMC_API_KEY")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("slug", "ethereum") 
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", cmc_api_key)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}

	fmt.Println(resp.Status)
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}

func GetCoinbasePrice() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	coinbase_api_key := os.Getenv("COINBASE_API_KEY")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.coinbase.com/v2/prices/ETH-USD/spot", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := url.Values{}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("Bearer", coinbase_api_key)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}

	fmt.Println(resp.Status)
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}

func GetOKXPrice() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.okx.com/api/v5/public/price-limit", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := url.Values{}
	q.Add("instId", "BTC-USDT-SWAP")

	req.Header.Set("Accepts", "application/json")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}

	fmt.Println(resp.Status)
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}
