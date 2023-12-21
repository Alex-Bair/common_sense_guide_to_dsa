package main

import "fmt"

func printNumbers(input []interface{}) {
	for _, value := range input {
		// must use a type switch to narrow down input type
		switch v := value.(type) {
		case int:
			fmt.Println(v)
		case []interface{}:
			printNumbers(v)
		}
	}
}

func main() {
	// had to use []interface{} to simulate multidimensional arrays that could hold ints or other arrays
	arr1 := []interface{}{
		1,
		2,
		3,
		[]interface{}{4, 5, 6},
		7,
		[]interface{}{
			8,
			[]interface{}{
				9, 10, 11,
				[]interface{}{
					12, 13, 14,
				},
			},
			[]interface{}{
				15, 16, 17, 18, 19,
				[]interface{}{
					20, 21, 22,
					[]interface{}{
						23, 24, 25,
						[]interface{}{
							26, 27, 29,
						},
					}, 30, 31,
				}, 32,
			}, 33,
		},
	}

	printNumbers(arr1)
}