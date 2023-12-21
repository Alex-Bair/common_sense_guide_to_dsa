package main

import "fmt"

func partition(arr []int, leftPointer, rightPointer int) int {
	// set up pivot as the rightmost element and adjust right pointer 
	pivotIndex := rightPointer
	pivot := arr[pivotIndex]
	rightPointer -= 1

	for {
		// move left pointer to the right until it encounters a value greater than or equal to the pivot
		for {
			if arr[leftPointer] >= pivot { break }
			leftPointer += 1
		}

		// move right pointer to the left until it encounters a value less than or equal to the pivot
		for {
			if arr[rightPointer] <= pivot || rightPointer == leftPointer { break }
			rightPointer -= 1
		}

		// if the left and right pointers have met (or passed each other), break out of the loop
		if leftPointer >= rightPointer {
			break
		} else {
		// otherwise, swap the values of the left and right pointers, shift the left pointer over, and repeat
			arr[leftPointer], arr[rightPointer] = arr[rightPointer], arr[leftPointer]
			leftPointer += 1
		}
	}

	// once the left and right pointers have met/passed each other, swap the values at the left pointer and the pivot
	arr[leftPointer], arr[pivotIndex] = arr[pivotIndex], arr[leftPointer]

	// return the value of the left pointer to use in the quickSort algorithm
	return leftPointer
}

func quickSort(arr []int , leftIndex, rightIndex int) {
	if rightIndex - leftIndex <= 0 { return }

	pivotIndex := partition(arr, leftIndex, rightIndex)

	quickSort(arr, leftIndex, pivotIndex - 1)
	quickSort(arr, pivotIndex + 1, rightIndex)
}

// find max number with efficiency of O(N^2)
func maxN2(arr []int) int {
	for i := 0; i < len(arr); i++ {
		isMax := true

		for j := 0; j < len(arr); j++ {
			if arr[j] > arr[i] {
				isMax = false
			}
		}

		if isMax {
			return arr[i]
		}
	}

	return -1
}

// find max number with efficiency of O(N log N)

func maxNLogN(arr []int) int {
	arrLength := len(arr)

	quickSort(arr, 0, arrLength - 1)

	return arr[arrLength - 1]
}

// find max number with efficiency of O(N)
func maxN(arr []int) (max int) {
	max = arr[0]

	for _, num1 := range arr[1:] {
		if num1 > max {
			max = num1
		}
	}

	return
}

func main() {
	arr := []int{5, 2, 4, 1, 0, 43, 23, 6, 34, 21, -3}

	fmt.Println(maxN2(arr))
	fmt.Println(maxN(arr))
	fmt.Println(maxNLogN(arr))
}