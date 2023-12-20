package main

import (
	"fmt"
)

func firstNonDuplicateChar(str string) string {
	count := make(map[rune]int)
	for _, runeVal := range str {
		count[runeVal] += 1
	}

	for _, runeVal := range str {
		if count[runeVal] == 1 {
			return string(runeVal)
		}
	}

	return ""
}

func main() {
	str := "minimum"
	
	fmt.Println(firstNonDuplicateChar(str))
}