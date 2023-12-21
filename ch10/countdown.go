package main

import "fmt"

func countdown(num int) {
	if num < 0 {return}
	fmt.Println(num)
	countdown(num - 1)
}

func main() {
	countdown(10)
}