package main

import "fmt"

func highestProduct(slice []int) (maxProduct int) {
	max := slice[0]
	min := slice[0]

	var product int

	for _, num := range slice[1:] {
		if num > 0 {
			product = num * max

			if num > max {
				max = num
			}
		} else if num < 0 {
			product = num * min

			if num < min {
				min = num
			}
		}

		if product > maxProduct {
			maxProduct = product
		}
	}

	return
}

func main() {
	slice := []int{5, -10, -6, 9, 4}

	fmt.Println(highestProduct(slice))
}