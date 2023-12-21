package main

import "fmt"

func numberUniquePaths(rows, columns int) int {
 if rows == 1 || columns == 1 {
	return 1
 }

 return numberUniquePaths(rows - 1, columns) + numberUniquePaths(rows, columns - 1)
}

func main() {
	rows := 3
	columns := 7

	fmt.Println(numberUniquePaths(rows, columns))
}