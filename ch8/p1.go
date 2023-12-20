package main

import "fmt"

func intersection(arr1, arr2 []int) []int {
	// hash table for values present in arr1
	arr1Values := make(map[int]bool)

	for _, val := range arr1 {
		arr1Values[val] = true
	}

	// create slice to contain the values at the intersection of arr1 and arr2
	intersection := make([]int, 0)

	// iterate through arr2 and append the values that are also in the arr1Values hash table to the intersection slice
	for _, val := range arr2 {
		if arr1Values[val] {
			intersection = append(intersection, val)
		}
	}

	return intersection
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{0, 2, 4, 6, 8}

	fmt.Println(intersection(arr1, arr2))
}