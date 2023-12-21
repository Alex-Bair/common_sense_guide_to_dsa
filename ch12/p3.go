package main

import "fmt"

func uniquePaths(rows, columns int) int {
	memo := make(map[string]int)

	return uniquePathsHelper(rows, columns, memo)
}

func uniquePathsHelper(rows, columns int, memo map[string]int) int {
	if rows == 1 || columns == 1 { return 1 }

	key := fmt.Sprint(rows, columns)

	_, present := memo[key]

	if !present {
		memo[key] = uniquePathsHelper(rows - 1, columns, memo) + uniquePathsHelper(rows, columns - 1, memo)
	}

	return memo[key]
}

func main() {
	columns := 7
	rows := 3

	fmt.Println(uniquePaths(rows, columns))
}