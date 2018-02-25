//Find largest distance
//Given an arbitrary unweighted rooted tree which consists of N (2 <= N <= 40000) nodes. The goal of the problem is to find largest distance between two nodes in a tree. Distance between two nodes is a number of edges on a path between the nodes (there will be a unique path between any pair of nodes since it is a tree). The nodes will be numbered 0 through N - 1.
//
//The tree is given as an array P, there is an edge between nodes P[i] and i (0 <= i < N). Exactly one of the i’s will have P[i] equal to -1, it will be root node.
//
// Example:
//If given P is [-1, 0, 0, 0, 3], then node 0 is the root and the whole tree looks like this:
//          0
//       /  |  \
//      1   2   3
//               \
//                4
// One of the longest path is 1 -> 0 -> 3 -> 4 and its length is 3, thus the answer is 3. Note that there are other paths with maximal distance.

package main

import "fmt"

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func solve(A []int) int {
	if len(A) == 0 {
		return 0
	}

	farthestNodeFromRoot, _ := dfs(0, A, -1, map[int]bool{})
	_, maxCount := dfs(farthestNodeFromRoot, A, -1, map[int]bool{})
	return maxCount
}

// Returns the farthest leaf and its distance
func dfs(node int, input []int, distance int, visited map[int]bool) (int, int) {

	fmt.Printf("Marking node as visited: %v\n", node)

	markVisited(node, visited)
	fmt.Printf("Current visited nodes: %v\n", visited)

	distance++
	fmt.Printf("Distance is: %v\n", distance)

	neighbors := getUnvisitedNeighbors(node, input, visited)
	if len(neighbors) == 0 {
		return node, distance
	}

	var currentLeaf int
	var currentMax int
	for _, neighbor := range neighbors {
		leaf, max := dfs(neighbor, input, distance, visited)
		if max > currentMax {
			currentMax = max
			currentLeaf = leaf
		}
	}
	fmt.Printf("Current maximum distance is %v to node %v\n\n", currentMax, currentLeaf)
	return currentLeaf, currentMax
}

func getUnvisitedNeighbors(node int, input []int, visited map[int]bool) []int {
	var neighbors []int
	if input[node] != -1 && !isVisited(input[node], visited) {
		neighbors = append(neighbors, input[node])
	}
	for i, n := range input {
		if n == node && !isVisited(i, visited) {
			neighbors = append(neighbors, i)
		}
	}
	fmt.Printf("Neighbors of %v are: %v\n", node, neighbors)
	return neighbors
}

func markVisited(node int, visited map[int]bool) {
	visited[node] = true
}

func isVisited(node int, visited map[int]bool) bool {
	if _, ok := visited[node]; ok {
		return true
	}
	return false
}

func main() {
	input := []int{-1, 0, 0, 0, 3}
	fmt.Printf("Input is: %v\n", input)

	length := solve(input)
	fmt.Printf("The longest length is: %v\n", length)
}
