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

type T = interface{}

type WorkerPool interface {
	Run()
	AddTask(task func())
}

type workerPool struct {
	numWorkers int
	queuedTask chan func()
}

func NewWorkerPool(numWorkers int) WorkerPool {
	return &workerPool {
		numWorkers: numWorkers,
		queuedTask: make(chan func(), numWorkers),
	}
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.numWorkers; i++ {
		go func(workerID int) {
			for task := range wp.queuedTask {
				task()
			}
		}(i + 1)
	} 
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTask <- task
}

type apiCall func()

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

func GetBitfinexPrice() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api-pub.bitfinex.com/v2/ticker/tETHUSD", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := url.Values{}

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

func GetDeribitPrice() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://test.deribit.com/api/v2/public/get_order_book", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := url.Values{}
	q.Add("depth", "1")
	q.Add("instrument_name", "ETH-PERPETUAL")
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
