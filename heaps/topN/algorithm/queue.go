package algorithm

// represents a queue
type queue struct {
	first *element
	last  *element
}

// represents an element of a queue
type element struct {
	value *Node
	next  *element
}

func (q *queue) enqueue(value *Node) {
	if q.isEmpty() {
		q.first = &element{
			value: value,
		}
		q.last = q.first
	} else {
		q.last.next = &element{
			value: value,
		}
	}
}

func (q *queue) dequeue() *Node {
	toDequeue := q.first.value
	q.first = q.first.next
	return toDequeue
}

func (q *queue) isEmpty() bool {
	return q.first == nil
}
