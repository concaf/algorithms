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

	if A == nil {
		return nil
	}

	var arr []int
	current := A
	stack := []*treeNode{}

	for {
		if current == nil {
			lastStackElement := len(stack) - 1
			if lastStackElement == -1 {
				break
			}
			arr = append(arr, stack[lastStackElement].value)
			current = stack[lastStackElement].right
			stack = stack[:lastStackElement]
		} else {
			stack = append(stack, current)
			current = current.left
		}
	}
	return arr
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
		left:  seven,
		right: eight,
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
