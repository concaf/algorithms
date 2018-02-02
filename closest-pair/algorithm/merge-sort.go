package algorithm

import (
	"log"
	"strings"
)

func MergeSort(input []*Point, axis string) []*Point {
	if len(input) == 0 {
		log.Fatal("The input array has no elements")
	}
	log.Printf("Input array: %v\n", getPoint(input))
	if len(input) == 1 {
		return input
	}

	array1 := MergeSort(input[:len(input)/2], axis)
	array2 := MergeSort(input[len(input)/2:], axis)
	return merge(array1, array2, axis)
}

func merge(array1, array2 []*Point, axis string) []*Point {
	log.Printf("Merging arrays %v and %v", getPoint(array1), getPoint(array2))
	sorted := make([]*Point, len(array1)+len(array2))
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

		switch strings.ToLower(axis) {
		case "x":
			if array1[posArr1].X < array2[posArr2].X {
				sorted[sortedPos] = array1[posArr1]
				posArr1++
				//} else if array2[posArr2].X <= array1[posArr1].X {
			} else {
				sorted[sortedPos] = array2[posArr2]
				posArr2++
			}
		case "y":
			if array1[posArr1].Y < array2[posArr2].Y {
				sorted[sortedPos] = array1[posArr1]
				posArr1++
				//} else if array2[posArr2].Y <= array1[posArr1].Y {
			} else {
				sorted[sortedPos] = array2[posArr2]
				posArr2++
			}
		default:
			log.Fatalf("Invalid axis passed: %v", axis)
		}
	}

	log.Printf("Merged arrays %v and %v to %v", getPoint(array1), getPoint(array2), getPoint(sorted))
	return sorted
}
