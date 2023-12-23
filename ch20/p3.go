package main

import "fmt"

func greatestProfit(slice []int) (maxProfit int) {
	low := slice[0]
	high := -1
	maxProfit = 0

	for _, price := range slice {
		if price < low {
			low = price
			high = price
		} else if price > high {
			high = price
			profit := high - low
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return
}

func main() {
	prices := []int{10, 7, 5, 8, 11, 2, 6}

	fmt.Println(greatestProfit(prices))
}