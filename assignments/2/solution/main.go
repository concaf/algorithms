package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	file, err := os.Open("../problem/IntegerArray.txt")
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

	log.Printf("Finding inversions in input array of size %v!", len(toSort))
	_, inversions :=  FindInversions(toSort)

	log.Printf("Found %v inversions", inversions)
}
