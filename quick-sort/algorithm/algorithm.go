package algorithm

import (
	"math/rand"
	"time"
)

func QuickSort(input []int) {
	if len(input) <= 1 {
		return
	}

	pivotIndex := choosePivot(input)

	// Swap to get pivot at index 0
	input[0], input[pivotIndex] = input[pivotIndex], input[0]
	pivotIndex = 0
	deflection := 1
	for unpartitioned := 1; unpartitioned < len(input); unpartitioned++ {
		if input[unpartitioned] < input[pivotIndex] {
			input[unpartitioned], input[deflection] = input[deflection], input[unpartitioned]
			deflection++
		}
	}
	// Swap to put pivot in its actual position after sorting
	input[pivotIndex], input[deflection-1] = input[deflection-1], input[pivotIndex]

	QuickSort(input[:deflection-1])
	QuickSort(input[deflection:])
}

func choosePivot(input []int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(input))
}
