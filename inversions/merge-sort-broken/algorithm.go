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

			for i := range array2 {
				log.Printf("%v in an inversion", []int{array1[posArr1], array2[i]})
				inversions = append(inversions, []int{array1[posArr1], array2[i]})
			}

			posArr1++
			continue
		}

		if array1[posArr1] < array2[posArr2] {
			sorted[sortedPos] = array1[posArr1]
			posArr1++
		} else if array2[posArr2] < array1[posArr1] {
			sorted[sortedPos] = array2[posArr2]
			log.Printf("%v in an inversion", []int{array1[posArr1], array2[posArr2]})
			inversions = append(inversions, []int{array1[posArr1], array2[posArr2]})
			posArr2++
		}
	}

	log.Printf("Merged arrays %v and %v to %v", array1, array2, sorted)
	return sorted, inversions
}
