package main

import "fmt"

func charCount(arr []string) int {
	if len(arr) == 1 {
		return len(arr[0])
	}

	return len(arr[0]) + charCount(arr[1:])
}

func main() {
	arr := []string{
		"ab",
		"c",
		"def",
		"ghij",
	}

	fmt.Println(charCount(arr))
}