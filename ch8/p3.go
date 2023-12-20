package main

import (
	"fmt"
	"strings"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

func missingLetter(str string) string {
	// use a hash map to keep track of which runes appear in the input string
	letterTracker := make(map[rune]bool)
	for _, runeValue := range strings.ToLower(str) {
		letterTracker[runeValue] = true
	}

	// iterate over all letters in the alphabet and return the first one that does not exist in the letterTracker hash map
	for _, runeValue := range alphabet {
		if !letterTracker[runeValue] {
			return string(runeValue)
		}
	}

	return ""
}

func main() {
	str := "the quick brown box jumps over a lazy dog"
	fmt.Println(missingLetter(str))
}