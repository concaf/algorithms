package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const arrayLength = 6

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	rand.Seed(time.Now().UnixNano())
	toSort := rand.Perm(arrayLength)
	log.Printf("The input array is: %v", toSort)
	_, inversions := FindInversions(toSort)
	log.Printf("Inversions are %v", inversions)
	log.Printf("Found %v inversions", len(inversions))
}
