package main

import "fmt"

func addUntil100(arr []int) int {
	arrLength := len(arr)

	if arrLength == 0 { return 0 }
	prelimSum := addUntil100(arr[1:])

	if arr[0] + prelimSum > 100 {
		return prelimSum
	} else {
		return arr[0] + prelimSum
	}

}

func main() {
	fmt.Println(addUntil100([]int{43, 239058123, 5}))
}