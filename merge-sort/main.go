package main

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
)

const arrayLength = 20

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	toSort := rand.Perm(arrayLength)
	sorted := mergeSort(toSort)
	log.Printf("The sorted array is: %v", sorted)

	if reflect.DeepEqual(sorted, integerSequence(arrayLength)) {
		log.Println("The array has been successfully sorted")
	} else {
		log.Println("The array is still unsorted")
	}
}

func integerSequence(n int) []int {
	var sequence []int
	for i := 0; i < n; i++ {
		sequence = append(sequence, i)
	}
	return sequence
}
