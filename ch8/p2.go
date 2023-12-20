package main

import "fmt"

func firstDuplicate(arr []string) string {
	// use hash map seenStrings to remember what strings we've already seen as we iterate over the input slice
	seenStrings := make(map[string]bool)

	for _, s := range arr {
		if seenStrings[s] {
			return s
		}

		seenStrings[s] = true
	}

	return ""
}

func main() {
	arr := []string{"a", "b", "c", "d", "c", "e", "f"}

	fmt.Println(firstDuplicate(arr))
}