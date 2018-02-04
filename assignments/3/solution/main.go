package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"sort"

	"github.com/containscafeine/algorithms/quick-sort/algorithm"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	firstElementPivot := func(input []int) int {
		return 0
	}

	lastElementPivot := func(input []int) int {
		return len(input) - 1
	}

	medianOfThreePivot := func(input []int) int {

		type pivot struct {
			element  int
			position int
		}

		first := pivot{
			input[0],
			0,
		}
		last := pivot{
			input[len(input)-1],
			len(input) - 1,
		}

		var middle pivot
		if len(input)%2 == 0 {
			middle = pivot{
				input[len(input)/2-1],
				len(input)/2 - 1,
			}
		} else {
			middle = pivot{
				input[len(input)/2],
				len(input) / 2,
			}
		}

		points := []int{first.element, middle.element, last.element}
		sort.Ints(points)
		switch points[1] {
		case first.element:
			return first.position
		case middle.element:
			return middle.position
		case last.element:
			return last.position
		default:
			log.Fatal("invalid case")
		}
		return 0
	}

	log.Printf("Total %v comparisons were made with first element as a pivot!", algorithm.QuickSort(NewArray(), firstElementPivot))
	log.Printf("Total %v comparisons were made with last element as a pivot!", algorithm.QuickSort(NewArray(), lastElementPivot))
	log.Printf("Total %v comparisons were made with median element as a pivot!", algorithm.QuickSort(NewArray(), medianOfThreePivot))
}

func NewArray() []int {
	file, err := os.Open("../problem/QuickSort.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	toSort := make([]int, 0)

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		toSort = append(toSort, val)
	}
	return toSort
}
