package main

import (
	"fmt"
	"math"
)

// Definition for binary tree
type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

/**
 * @input A : Root pointer of the tree
 *
 * @Output Integer
 */
func isBalanced(A *treeNode) int {
	if balancedHeight(A) == -1 {
		return 0
	}
	return 1
}

func balancedHeight(A *treeNode) float64 {
	if A == nil {
		fmt.Printf("Current node is nil\n")
		return 0
	}
	fmt.Printf("Current node is: %v\n", A.value)

	leftHeight := balancedHeight(A.left)
	if A.left != nil {
		fmt.Printf("Left height of %v is %v\n", A.left.value, leftHeight)
	}
	rightHeight := balancedHeight(A.right)
	if A.right != nil {
		fmt.Printf("Right height of %v is %v\n", A.right.value, rightHeight)
	}

	if leftHeight == -1 || rightHeight == -1 {
		return -1
	}

	if math.Abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return 1 + math.Max(leftHeight, rightHeight)
}

func main() {
	nine := &treeNode{
		value: 9,
	}
	eight := &treeNode{
		value: 8,
		left:  nine,
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
		left:  eight,
	}
	three := &treeNode{
		value: 3,
		left:  six,
		right: seven,
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

	fmt.Printf("Output: %v\n", isBalanced(one))
}
