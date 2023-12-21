package main

import "fmt"

func sumFrom(low, high int) (sum int) {
	if low == high {
		return low
	}

	return high + sumFrom(low, high - 1)
}

func main() {
	fmt.Println(sumFrom(1, 10))
}