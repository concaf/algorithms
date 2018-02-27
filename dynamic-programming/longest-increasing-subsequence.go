package main

import "fmt"

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func lis(A []int) int {
	length := len(A)
	if length == 1 {
		return 1
	}
	cache := make(map[int]int)
	max := 0
	for i := 0; i <= len(A)-2; i++ {
		currentMax := lisElement(i, A, length, cache)
		if currentMax > max {
			max = currentMax
		}
	}
	return max
}

func lisElement(pos int, A []int, length int, cache map[int]int) int {
	if _, ok := cache[pos]; ok {
		return cache[pos]
	}

	if length == 0 {
		return 1
	}
	max := 1
	for i := pos + 1; i < length; i++ {
		if A[pos] < A[i] {
			newMax := lisElement(i, A, length, cache) + 1
			if newMax > max {
				max = newMax
			}
		}
	}
	cache[pos] = max
	return max
}

func main() {
	input := []int{69, 54, 19, 51, 16, 54, 64, 89, 72, 40, 31, 43, 1, 11, 82, 65, 75, 67, 25, 98, 31, 77, 55, 88, 85, 76, 35, 101, 44, 74, 29, 94, 72, 39, 20, 24, 23, 66, 16, 95, 5, 17, 54, 89, 93, 10, 7, 88, 68, 10, 11, 22, 25, 50, 18, 59, 79, 87, 7, 49, 26, 96, 27, 19, 67, 35, 50, 10, 6, 48, 38, 28, 66, 94, 60, 27, 76, 4, 43, 66, 14, 8, 78, 72, 21, 56, 34, 90, 89}
	fmt.Printf("Input is: %v\n", input)
	fmt.Printf("The longest length is: %v\n", lis(input))
}
