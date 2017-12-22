package main

import "fmt"

func (q queue) enqueue(e int) queue {
	return append(q, e)
}

func (q queue) dequeue() queue {
	return q[1:]
}

type queue []int

func main() {
	var q1 queue
	q1 = q1.enqueue(1)
	fmt.Println(q1)
	q1 = q1.enqueue(2)
	fmt.Println(q1)
	q1 = q1.enqueue(4)
	fmt.Println(q1)
	q1 = q1.enqueue(9)
	fmt.Println(q1)
	q1 = q1.dequeue()
	fmt.Println(q1)
	q1 = q1.dequeue()
	fmt.Println(q1)
}
