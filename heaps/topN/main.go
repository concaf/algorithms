package main

import (
	"fmt"
	"github.com/containscafeine/algorithms/heaps/topN/algorithm"
)

func main() {
	var heap algorithm.Node
	heap.Insert(1)
	fmt.Println(heap.GetMin())
}
