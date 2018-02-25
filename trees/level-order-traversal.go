package main

import "fmt"

//Definition for binary tree
type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

/**
 * @input A : Root pointer of the tree
 *
 * @Output 2D int array.
 */
func levelOrder(A *treeNode) [][]int {
	levelOrderArray := make([][]int, 0)
	levelOrderArray = levelOrderLevel(A, 0, levelOrderArray)
	return levelOrderArray
}

func levelOrderLevel(A *treeNode, level int, levelOrderArray [][]int) [][]int {
	if A == nil {
		return levelOrderArray
	}

	isLevel := false
	for i := range levelOrderArray {
		if i == level {
			isLevel = true
			break
		}
	}

	if !isLevel {
		levelOrderArray = append(levelOrderArray, nil)
	}

	levelOrderArray[level] = append(levelOrderArray[level], A.value)
	levelOrderArray = levelOrderLevel(A.left, level+1, levelOrderArray)
	levelOrderArray = levelOrderLevel(A.right, level+1, levelOrderArray)
	return levelOrderArray
}

func main() {
	eight := &treeNode{
		value: 8,
	}
	seven := &treeNode{
		value: 7,
	}
	six := &treeNode{
		value: 6,
	}
	five := &treeNode{
		value: 5,
	}
	four := &treeNode{
		value: 4,
		right: six,
	}
	three := &treeNode{
		value: 3,
		left:  seven,
		right: eight,
	}
	two := &treeNode{
		value: 2,
		left:  four,
		right: five,
	}
	one := &treeNode{
		value: 1,
		left:  two,
		right: three,
	}

	fmt.Printf("Output: %v\n", levelOrder(one))
}
