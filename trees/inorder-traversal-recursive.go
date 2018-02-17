package main

import "fmt"

// Definition for binary tree
type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

/**
 * @input A : Root pointer of the tree
 *
 * @Output Integer array.
 */
func inorderTraversal(A *treeNode) []int {
	//fmt.Printf("Now traversing: %v\n", A.value)
	var arr []int

	if A.left != nil {
		arr = append(arr, inorderTraversal(A.left)...)
	}

	arr = append(arr, A.value)

	if A.right != nil {
		arr = append(arr, inorderTraversal(A.right)...)
	}
	return arr
}

func main() {
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

	fmt.Printf("Output: %v\n", inorderTraversal(one))
}
