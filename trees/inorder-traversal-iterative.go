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

	//if r == nil {
	//	return []int{}
	//}
	//
	//st := make([]*treeNode, 0)
	//node := r
	//for node != nil {
	//	st = append(st, node)
	//	node = node.left
	//}
	//
	//result := make([]int, 0)
	//for len(st) > 0 {
	//	node = st[len(st)-1]
	//	st = st[:len(st)-1]
	//
	//	result = append(result, node.value)
	//	if node.right != nil {
	//
	//		node = node.right
	//		for node != nil {
	//			st = append(st, node)
	//			node = node.left
	//		}
	//	}
	//}
	//
	//return result
}

func main() {
	//eight := &treeNode{
	//	value: 8,
	//}
	//seven := &treeNode{
	//	value: 7,
	//}
	//six := &treeNode{
	//	value: 6,
	//	left:  seven,
	//	right: eight,
	//}
	five := &treeNode{
		value: 5,
	}
	four := &treeNode{
		value: 4,
		//right: six,
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
