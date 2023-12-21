package main

import "fmt"

func isEven(num int) bool {
	return num % 2 == 0
}

func justEvens(arr []int) []int {
	arrLength := len(arr)

	if arrLength == 0 {
		return []int{}
	}

	if isEven(arr[arrLength - 1]) {
		return append(justEvens(arr[:arrLength - 1]), arr[arrLength - 1])
	} else {
		return justEvens(arr[:arrLength - 1])
	}
}

func main() {
	arr := []int{ 1, 2, 3, 4, 5, 6,	7, 8,	9, 10 }

	fmt.Println(justEvens(arr))
}