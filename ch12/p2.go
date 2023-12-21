package main

import "fmt"

func golomb(n int) int {
	memo := map[int]int{
		1: 1,
	}

	return golombHelper(n, memo)
}

func golombHelper(n int, memo map[int]int) int {
	_, present := memo[n]

	if !present {
		memo[n] = 1 + golombHelper(n - golombHelper(golombHelper(n - 1, memo), memo), memo)
	}

	return memo[n]
}

func main() {
	fmt.Println(golomb(43))
}