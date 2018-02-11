package main

import "fmt"

type listNode struct {
	value int
	next  *listNode
}

func (l *listNode) print() string {
	if l.next == nil {
		return fmt.Sprintf("%v -> nil", l.value)
	}
	return fmt.Sprintf("%v -> %v", l.value, l.next.print())
}

/**
 * @input A : Head pointer of linked list
 *
 * @Output head pointer of list.
 */
func reverseList(A *listNode) *listNode {
	if A.next == nil {
		return A
	}

	if A.next.next == nil {
		toReturn := A.next
		A.next.next = A
		A.next = nil
		return toReturn
	}

	current := A.next.next
	last := A.next
	penultimate := A
	penultimate.next = nil
	for {
		fmt.Printf("Current: %v\n", current.value)
		fmt.Printf("Last: %v\n", last.value)
		fmt.Printf("Penultimate: %v\n", penultimate.value)

		currentLast := last

		last.next = penultimate
		fmt.Printf("%v now points to %v\n", last.value, penultimate.value)

		if current.next == nil {
			current.next = last
			fmt.Printf("%v now points to %v\n", current.value, last.value)
			return current
		} else {
			last = current
			current = current.next
		}

		penultimate = currentLast
	}
	return nil
}

func main() {
	fifth := listNode{
		value: 5,
		next:  nil,
	}
	fourth := listNode{
		value: 4,
		next:  &fifth,
	}
	third := listNode{
		value: 3,
		next:  &fourth,
	}
	second := listNode{
		value: 2,
		next:  &third,
	}
	first := listNode{
		value: 1,
		next:  &second,
	}

	//test := listNode{
	//	value: 1,
	//	next:  &fifth,
	//}

	fmt.Printf("Input list is:\n%v\n", first.print())
	reversed := reverseList(&first)
	fmt.Printf("Reversed list is:\n%v\n", reversed.print())
}
