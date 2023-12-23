package main

import (
	"fmt"
	"math"
)

func roundFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val * ratio) / ratio
}

func sort(slice []float64) (sorted []float64) {
	hash := make(map[float64]int)

	for _, value := range slice {
		hash[value] += 1
	}

	for i := 970; i <= 990; i += 1 {
		t := float64(i) / 10.0

		if hash[t] > 0 {
			for iteration := 1; iteration <= hash[t]; iteration++ {
				sorted = append(sorted, t)
			}
		}
	}

	return
}

func main() {
	temps := []float64{98.6, 98.0, 97.1, 99.0, 98.9, 97.8, 98.5, 98.2, 98.0, 97.1}

	fmt.Println(sort(temps))
}