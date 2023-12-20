package main

import "fmt"

func greatestNumber(arr []int) (currentGreatest int) {
	currentGreatest = arr[0]

	for _, num := range arr[1:] {
		if num > currentGreatest {
			currentGreatest = num
		}
	}

	return 
}

func main() {
	fmt.Println(greatestNumber([]int{1, 43, 2, 6, 15, 2}))
}