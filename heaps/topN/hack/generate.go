package main

import (
	"fmt"
	"math/rand"
	"os"
)

const fileName = "input.txt"

func main() {
	generateAndWrite(1000, fileName)
}

func generateAndWrite(n int, filename string) {
	f, err := os.Create(filename)
	checkError(err)
	defer f.Close()

	for i := 0; i < n; i++ {
		data := rand.Int()
		_, err := f.WriteString(fmt.Sprintf("%d\n", data))
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
