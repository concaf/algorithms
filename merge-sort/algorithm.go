package main

import "log"

func mergeSort(input []int) []int {
	if len(input) == 0 {
		log.Fatal("The input array has no elements")
	}
	log.Printf("Input array: %v\n", input)
	if len(input) == 1 {
		return input
	}

	array1 := mergeSort(input[:len(input)/2])
	array2 := mergeSort(input[len(input)/2:])
	return merge(array1, array2)
}

func merge(array1, array2 []int) []int {
	log.Printf("Merging arrays %v and %v", array1, array2)
	sorted := make([]int, len(array1)+len(array2))
	posArr1 := 0
	posArr2 := 0

	for sortedPos := 0; sortedPos < len(sorted); sortedPos++ {
		if posArr1 == len(array1) {
			for r := posArr2; r < len(array2); r++ {
				sorted[sortedPos] = array2[r]
				sortedPos++
			}
			break
		}
		if posArr2 == len(array2) {
			for r := posArr1; r < len(array1); r++ {
				sorted[sortedPos] = array1[r]
				sortedPos++
			}
			break
		}

		if array1[posArr1] < array2[posArr2] {
			sorted[sortedPos] = array1[posArr1]
			posArr1++
		} else if array2[posArr2] < array1[posArr1] {
			sorted[sortedPos] = array2[posArr2]
			posArr2++
		}
	}

	log.Printf("Merged arrays %v and %v to %v", array1, array2, sorted)
	return sorted
}
