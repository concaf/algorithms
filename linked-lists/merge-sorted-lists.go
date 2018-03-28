/**
 * @input A : Head pointer of linked list
 * @input B : Head pointer of linked list
 *
 * @Output head pointer of list.
 */

package main

import "fmt"

type listNode struct {
	value int
	next  *listNode
}

func mergeTwoLists(A *listNode, B *listNode) *listNode {
	var mergedList *listNode
	currentA := A
	currentB := B
	var first *listNode

	for {
		if currentA == nil && currentB == nil {
			break
		}

		if currentA == nil {
			mergedList.addNextToList(currentB.value)
			mergedList = mergedList.next
			currentB = currentB.next
			continue
		}

		if currentB == nil {
			mergedList.addNextToList(currentA.value)
			mergedList = mergedList.next
			currentA = currentA.next
			continue
		}

		if mergedList == nil {
			if currentA.value <= currentB.value {
				mergedList = &listNode{
					value: currentA.value,
				}
				currentA = currentA.next
			} else {
				mergedList = &listNode{
					value: currentB.value,
				}
				currentB = currentB.next
			}
			first = mergedList
			continue
		}

		if currentA.value <= currentB.value {
			mergedList.addNextToList(currentA.value)
			mergedList = mergedList.next
			currentA = currentA.next
		} else {
			mergedList.addNextToList(currentB.value)
			mergedList = mergedList.next
			currentB = currentB.next
		}
	}
	return first
}

func (n *listNode) addNextToList(value int) {
	n.next = &listNode{
		value: value,
	}
}

func (n *listNode) print() {
	if n == nil {
		fmt.Println("Empty")
		return
	}
	for n != nil {
		fmt.Printf("%v -> ", n.value)
		n = n.next
	}
}

func main() {
	list1 := &listNode{
		value: 1,
		next: &listNode{
			value: 2,
			next: &listNode{
				value: 3,
			},
		},
	}

	list2 := &listNode{
		value: 4,
		next: &listNode{
			value: 5,
			next: &listNode{
				value: 6,
			},
		},
	}

	mergedList := mergeTwoLists(list1, list2)
	mergedList.print()
}
