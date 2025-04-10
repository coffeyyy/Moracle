package main

import (
	"fmt"
	"time"
)

func main() {

	api_calls := [5]apiCall{
	GetCMCPrice,
	GetCoinbasePrice,
	GetOKXPrice,
	GetBitfinexPrice,
	GetDeribitPrice,
}

	workerCount := 5
	wp := NewWorkerPool(workerCount)
	wp.Run()


	for i := 0; i < len(api_calls); i++ {
		taskID := i
		wp.AddTask(func(taskID int) func() {
			return func() {
				fmt.Printf("Processing API Num: %d\n", taskID)
				api_calls[taskID]()
			}
		} (taskID))
	}

	time.Sleep(2 * time.Second)
}


