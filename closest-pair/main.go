package main

import (
	"fmt"
	"log"

	"github.com/containscafeine/algorithms/closest-pair/algorithm"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	input := []*algorithm.Point{
		{76, 23},
		{4, 5},
		{3, 5},
		{12, 44},
		{9, 23},
		{21,64},
		{5, 0},
		{6, 3},
		{29, 12},
		{1, 1},
		{12,54},
		{5, 7},
		{4, 30},
		{2, 3},
		{6, 85},
		{32, 64},
		{1, 4},
		{10, 1},
		{76, 3},
		{11, 2},
		{8, 12},
		{13, 4},
		{21, 2},
		{5, 20},
		{10, 42},
	}

	pointsSortedOnX := algorithm.MergeSort(input, "x")
	pointsSortedOnY := algorithm.MergeSort(input, "y")
	minimumDistance := algorithm.FindClosestPair(pointsSortedOnX, pointsSortedOnY)
	log.Printf("The minimum distance between points is: %v", minimumDistance)
}
