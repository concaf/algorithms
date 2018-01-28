package main

import "log"

func findInversions(input []int) ([]int, [][]int) {
	var inversions [][]int
	if len(input) == 0 {
		log.Fatal("The input array has no elements")
	}
	log.Printf("Input array: %v\n", input)
	if len(input) == 1 {
		return input, inversions
	}

	array1, inv1 := findInversions(input[:len(input)/2])
	inversions = append(inversions, inv1...)
	array2, inv2 := findInversions(input[len(input)/2:])
	inversions = append(inversions, inv2...)
	return merge(array1, array2, inversions)
}

func merge(array1, array2 []int, inversions [][]int) ([]int, [][]int) {
	log.Printf("Merging arrays %v and %v", array1, array2)
	sorted := make([]int, len(array1)+len(array2))
	posArr1 := 0
	posArr2 := 0

	for sortedPos := 0; sortedPos < len(sorted); sortedPos++ {
		if posArr1 == len(array1) {
			sorted[sortedPos] = array2[posArr2]
			posArr2++
			continue
		}
		if posArr2 == len(array2) {
			sorted[sortedPos] = array1[posArr1]
			posArr1++
			continue
		}

		if array1[posArr1] < array2[posArr2] {
			sorted[sortedPos] = array1[posArr1]
			posArr1++
		} else if array2[posArr2] < array1[posArr1] {
			// This is the case where the split inversions will occur.

			// Since both the arrays are sorted, we can safely say that the
			// remaining elements in the first array from the current element on
			// are all smaller than the current element of the second array.
			for remainingInArr1 := posArr1; remainingInArr1 < len(array1); remainingInArr1++ {
			log.Printf("%v in an inversion", []int{array1[remainingInArr1], array2[posArr2]})
			inversions = append(inversions, []int{array1[remainingInArr1], array2[posArr2]})
			}

			sorted[sortedPos] = array2[posArr2]
			posArr2++
		}
	}

	log.Printf("Merged arrays %v and %v to %v", array1, array2, sorted)
	return sorted, inversions
}
