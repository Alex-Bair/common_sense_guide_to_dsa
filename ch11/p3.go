package main

import "fmt"

func triangleNumber(n int) int {
	if n == 1 {
		return 1
	}

	return n + triangleNumber(n - 1)
}

func main() {
	for i := 1; i < 8; i++ {
		fmt.Println(triangleNumber(i))
	}
}