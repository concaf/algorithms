/*
Write a program, topN, that given an arbitrarily large file and a number, N, containing individual numbers on each line (e.g. 200Gb file), will output the largest N numbers, highest first.
*/

package main

import (
	"bufio"
	"fmt"
	"github.com/containscafeine/algorithms/heaps/topN/algorithm"
	"os"
	"strconv"
)

const fileName = "input.txt"

func main() {

	f, err := os.Open(fileName)
	checkError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	heap := algorithm.New()
	var count int
	for scanner.Scan() {
		count++
		fmt.Println(count)
		input, err := strconv.Atoi(scanner.Text())
		checkError(err)
		heap.Insert(input)
	}
	fmt.Println(heap.GetMin())

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
