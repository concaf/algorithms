package main

import (
	"fmt"
	"github.com/containscafeine/algorithms/quick-sort/algorithm"
	"log"
	"math/rand"
	"reflect"
	"time"
)

const arrayLength = 20000

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	choosePivot := func(input []int) int {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(len(input))
	}

	toSort := rand.Perm(arrayLength)
	algorithm.QuickSort(toSort, choosePivot)
	log.Printf("The sorted array is: %v", toSort)

	if reflect.DeepEqual(toSort, integerSequence(arrayLength)) {
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
