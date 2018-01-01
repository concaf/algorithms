package main

import "log"

func findInversions(input []int) [][]int {
	var inversions [][]int
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				log.Printf("Found inversion: %v", []int{input[i], input[j]})
				inversions = append(inversions, []int{input[i], input[j]})
			}
		}
	}
	return inversions
}
