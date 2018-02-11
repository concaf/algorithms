package main

import (
	"fmt"
	"log"

	"github.com/containscafeine/algorithms/binary-search/algorithm"
)

const arrayLength = 900

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	//input := integerSequence(arrayLength)
	input := []int{0, 2, 3, 4, 10, 40, 44}
	log.Printf("Input is: %v", input)
	element := 44
	log.Printf("Searching for: %v", element)
	position := algorithm.BinarySearch(input, element)
	log.Printf("Found element at position: %v", position)
}

func integerSequence(n int) []int {
	var sequence []int
	for i := 0; i < n; i++ {
		sequence = append(sequence, i)
	}
	return sequence
}
