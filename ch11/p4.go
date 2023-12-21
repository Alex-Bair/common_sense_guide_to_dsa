package main

import "fmt"

func findXIndex(str string) int {
	if str[0] == 'x' {
		return 0
	}

	return 1 + findXIndex(str[1:])
}

func main() {
	str := "abcdefghijklmnopqrstuvwxyz"

	fmt.Println(findXIndex(str))
}