package main

import (
	"fmt"
	"log"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	input := []int{2, 7, 6, 65, 3, 14, 1}
	selectionSort(input)
}

func getSmallestElementIndex(input []int) int {
	// [65 6 14 7]
	smallest := input[0]
	pos := 0
	for i := 1; i < len(input); i++ {
		if input[i] < smallest {
			pos = i
		}
	}
	log.Printf("Found smallest element in %v to be %v", input, input[pos])
	return pos
}

func selectionSort(input []int) {
	if len(input) == 0 {
		log.Fatal("The input array has no elements")
	}
	log.Printf("Input array: %v\n\n", input)
	for pos := 0; pos < len(input); pos++ {
		// We need to add pos to the returned index since the index is relative to
		// the sliced array, not the original input array
		smallestPos := getSmallestElementIndex(input[pos:]) + pos
		log.Printf("Replacing first element %v with the smallest element %v", input[pos], input[smallestPos])
		input[pos], input[smallestPos] = input[smallestPos], input[pos]
		log.Printf("Array after replacing: %v", input)
		log.Printf("\n")
	}
}
