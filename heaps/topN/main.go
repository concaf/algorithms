/*
Write a program, topN, that given an arbitrarily large file and a number, N, containing individual numbers on each line (e.g. 200Gb file), will output the largest N numbers, highest first.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/containscafeine/algorithms/heaps/topN/algorithm"
)

var (
	file    string
	numbers int
)

func main() {
	parse()

	f, err := os.Open(file)
	checkError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	heap := algorithm.New()
	var count int
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		checkError(err)

		if count < numbers {
			heap.Insert(input)
		} else if input > heap.GetMin() {
			heap.ExtractMin()
			heap.Insert(input)
		}

		count++
	}

	var minFirstArray []int
	for heap.GetMin() != -1 {
		minFirstArray = append(minFirstArray, heap.ExtractMin())
	}

	fmt.Printf("The maximum %v numbers in %v are:\n", numbers, file)
	for i := len(minFirstArray) - 1; i >= 0; i-- {
		fmt.Println(minFirstArray[i])
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parse() {
	flag.StringVar(&file, "file", "input.txt", "file to read")
	flag.IntVar(&numbers, "numbers", 8, "how many largest numbers to return")
	flag.Parse()
}
