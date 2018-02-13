package main

import "fmt"

//import "fmt"

type listNode struct {
	value int
	next  *listNode
}

func (l *listNode) print() string {
	if l == nil {
		return ""
	}
	if l.next == nil {
		return fmt.Sprintf("%v -> nil", l.value)
	}
	return fmt.Sprintf("%v -> %v", l.value, l.next.print())
}

/**
 * @input A : Head pointer of linked list
 * @input B : Integer
 *
 * @Output head pointer of list.
 */
func reverseList(A *listNode, B int) *listNode {
	current := A
	var toReturn *listNode
	var firstDone bool
	var last *listNode
	var newReversedHead *listNode
	var oldLast = &listNode{
		value: 100,
		next:  nil,
	}

	for {
		if current == nil {
			return toReturn
		}
		fmt.Printf("Will now reverse: %v at intervals of %v\n ", current.print(), B)
		newReversedHead, last, current = reverseLinkedList(current, B)
		fmt.Printf("Reversed is: %v\n", newReversedHead.print())
		if last == nil {
			fmt.Printf("The last element of reversed list is: nil\n")
		} else {
			fmt.Printf("The last element of reversed list is: %v\n", last.value)
		}
		oldLast.next = newReversedHead
		oldLast = last
		if !firstDone {
			toReturn = newReversedHead
		}
		firstDone = true
		fmt.Printf("Will return: %v\n", toReturn.print())
	}
	return toReturn
}

// returns first, last, next elements
func reverseLinkedList(A *listNode, B int) (*listNode, *listNode, *listNode) {
	if A == nil {
		return nil, nil, nil
	}
	if A.next == nil {
		return A, nil, nil
	}

	firstElement := A

	if A.next.next == nil {
		if B == 1 {
			return A, A, A.next
		}

		toReturn := A.next
		A.next.next = A
		A.next = nil
		return toReturn, nil, nil
	}

	current := A.next.next
	last := A.next
	penultimate := A
	penultimate.next = nil
	for {
		if B == 1 {
			firstElement.next = nil
			fmt.Printf("%v now points to nil\n", firstElement.value)
			return penultimate, firstElement, last
		}
		fmt.Printf("Current: %v\n", current.value)
		fmt.Printf("Last: %v\n", last.value)
		fmt.Printf("Penultimate: %v\n", penultimate.value)

		currentLast := last

		last.next = penultimate
		fmt.Printf("%v now points to %v\n", last.value, penultimate.value)

		if current.next == nil {
			current.next = last
			fmt.Printf("Next element is nil, so %v now points to %v\n", current.value, last.value)
			return current, firstElement, nil
		} else {
			last = current
			current = current.next
		}

		penultimate = currentLast
		B--
	}
	return nil, nil, nil
}

func main() {
	eighth := listNode{
		value: 8,
		next:  nil,
	}
	seventh := listNode{
		value: 7,
		next:  &eighth,
	}
	sixth := listNode{
		value: 6,
		next:  &seventh,
	}
	fifth := listNode{
		value: 5,
		next:  &sixth,
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

	fmt.Printf("Input list is:\n%v\n", first.print())
	reversed := reverseList(&first, 3)
	fmt.Printf("Reversed list is:\n%v\n", reversed.print())
}
