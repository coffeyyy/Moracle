package main

import (
	"math"
	"time"
)

func calculate_standard_deviation(float mean, int size, prices map[string]int) float {
	sum_square_diffs := 0
	for _, price := range prices {
		sum_square_diffs += ((prices[i] - mean) ** 2)
	}
	return math.Sqrt(sum_square_diffs/size-1)
}

func calculate_mean(prices map[string]int, int size) float {
	sum := 0
	for _, price := range prices {
		sum += price
	}
	return sum/size
}