package main

import (
	"fmt"
	"github.com/containscafeine/algorithms/heaps/topN/algorithm"
)

func main() {
	heap := algorithm.New()
	heap.Insert(100)

	heap.ExtractMin()
	fmt.Println(heap.GetMin())
}
