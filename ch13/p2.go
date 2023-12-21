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

func findMissingNumber(arr []int) (product int) {
	quickSort(arr, 0, len(arr) - 1)

	for index, val := range arr {
		if val != index {
			return index
		}
	}

	return -1
}

func main() {
	arr1 := []int{5, 2, 4, 1, 0}
	arr2 := []int{9, 3, 2, 5, 6, 7, 1, 0, 4}

	fmt.Println(findMissingNumber(arr1))
	fmt.Println(findMissingNumber(arr2))
}