package main

import "fmt"

func findMissing(slice []int) int {
	expectedSum := 0
	actualSum := 0

	for index, value := range slice {
		expectedSum += index + 1
		actualSum += value
	}

	return expectedSum - actualSum
}

func main() {
	slice1 := []int{2, 3, 0, 6, 1, 5}
	slice2 := []int{8, 2, 3, 9, 4, 7, 5, 0, 6}

	fmt.Println(findMissing(slice1))
	fmt.Println(findMissing(slice2))
}