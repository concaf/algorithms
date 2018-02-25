package main

import "fmt"

/**
 * @input A : array of strings
 *
 * @Output Integer
 */
func black(A []string) int {
	inputLength := len(A)
	var stringLength int
	if len(A) != 0 {
		stringLength = len(A[0])
	}
	fmt.Printf("Input length is: %v\n", inputLength)
	fmt.Printf("String length is: %v\n", stringLength)
	count := 0
	var visited []X
	for num, str := range A {
		for pos, char := range str {
			if string(char) == "X" {
				node := X{
					stringNumber: num,
					position:     pos,
				}
				fmt.Printf("Found X at %v\n", node)
				if !node.isVisited(visited) {
					fmt.Printf("Node was not visited before: %v\n", node)
					fmt.Printf("Starting BFS at: %v\n", node)
					visited = node.bfs(visited, A, stringLength, inputLength)
					count++
					fmt.Printf("Total black count is now: %v\n\n", count)
				}
			}
		}
	}
	return count
}

type X struct {
	stringNumber int
	position     int
}

func (node *X) bfs(visited []X, input []string, stringLength int, inputLength int) []X {
	var queue []X
	queue = append(queue, *node)

	for len(queue) > 0 {
		fmt.Printf("Queue is: %v\n", queue)
		current := queue[0]
		fmt.Printf("Current element: %v\n", current)
		visited = current.markVisited(visited)
		queue = queue[1:]
		fmt.Printf("Visited nodes are: %v\n", visited)

		queue = append(queue, current.getNeighbors(input, visited, stringLength, inputLength)...)
	}
	return visited
}

func (node *X) markVisited(visited []X) []X {
	return append(visited, *node)
}

func (node *X) isVisited(visited []X) bool {
	for _, alreadyVisited := range visited {
		if *node == alreadyVisited {
			return true
		}
	}
	return false
}

func (node *X) isX(input []string) bool {
	if string(input[node.stringNumber][node.position]) == "X" {
		return true
	}
	return false
}

func (node *X) getNeighbors(input []string, visited []X, stringLength int, inputLength int) []X {
	var neighbors []X

	// Add left
	if node.position > 0 {
		potentialNeighbor := X{
			stringNumber: node.stringNumber,
			position:     node.position - 1,
		}
		if potentialNeighbor.isX(input) && !potentialNeighbor.isVisited(visited) {
			neighbors = append(neighbors, potentialNeighbor)
		}
	}

	// Add right
	if node.position < stringLength-1 {
		potentialNeighbor := X{
			stringNumber: node.stringNumber,
			position:     node.position + 1,
		}
		if potentialNeighbor.isX(input) && !potentialNeighbor.isVisited(visited) {
			neighbors = append(neighbors, potentialNeighbor)
		}
	}

	// Add top
	if node.stringNumber > 0 && string(input[node.stringNumber-1][node.position]) == "X" {
		potentialNeighbor := X{
			stringNumber: node.stringNumber - 1,
			position:     node.position,
		}
		if potentialNeighbor.isX(input) && !potentialNeighbor.isVisited(visited) {
			neighbors = append(neighbors, potentialNeighbor)
		}
	}

	// Add bottom
	if node.stringNumber < inputLength-1 {
		potentialNeighbor := X{
			stringNumber: node.stringNumber + 1,
			position:     node.position,
		}
		if potentialNeighbor.isX(input) && !potentialNeighbor.isVisited(visited) {
			neighbors = append(neighbors, potentialNeighbor)
		}
	}

	fmt.Printf("%v has neighbors: %v\n", *node, neighbors)
	return neighbors
}

func main() {
	input := []string{
		"OOOXOOO",
		"OOXXOXO",
		"OXOOOXO",
		"OOXXXOO",
		"OOXXXOO",
		"OOXXXOO",
	}
	fmt.Sprintf("Input is:\n")
	for _, str := range input {
		fmt.Println(str)
	}

	count := black(input)
	fmt.Printf("There are %v black shapes\n", count)
}
