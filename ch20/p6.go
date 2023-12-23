package main

import "fmt"

func longestConsecutiveSequence(slice []int) (maxLength int) {
	hash := make(map[int]bool)

	for _, num := range slice {
		hash[num] = true
	}

	for _, num := range slice {
		if hash[num - 1] {
			continue
		}

		length := 0

		for i := num; hash[i] == true; i++ {
			length++
		}

		if length > maxLength {
			maxLength = length
		}
	}

	return
}

func main() {
	slice1 := []int{10, 5, 12, 3, 55, 30, 4, 11, 2}
	slice2 := []int{19, 13, 15, 12, 18, 14, 17, 11}

	fmt.Println(longestConsecutiveSequence(slice1))
	fmt.Println(longestConsecutiveSequence(slice2))
}