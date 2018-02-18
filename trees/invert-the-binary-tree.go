package main

// Definition for binary tree
type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

/**
 * @input A : Root pointer of the tree
 *
 * @Output root pointer of tree.
 */
func invertTree(A *treeNode) *treeNode {
	if A == nil {
		return nil
	}
	A.left, A.right = invertTree(A.right), invertTree(A.left)
	return A
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

	invertTree(one)
}
