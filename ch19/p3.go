package main

import "fmt"

// We can have a space complexity of O(1) by mutating the input slice
func reverse(slice []string) []string {
	sliceLength := len(slice)

	for i := 0; i < sliceLength / 2; i++ {
		slice[i], slice[sliceLength - 1 - i] = slice[sliceLength - 1 - i], slice[i]
	}

	return slice
}

func main() {
	arr := []string{"a", "b", "c", "d", "e", "f"}

	reverse(arr)

	fmt.Println(arr)
}