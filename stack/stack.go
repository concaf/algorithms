package main

import "fmt"

type Stack []string

func (stack *Stack) push(item string) {
	*stack = append(*stack, item)
}

func (stack *Stack) pop() (popped string) {
	*stack, popped = (*stack)[:len(*stack)-1], (*stack)[len(*stack)-1]
	return
}

func (stack *Stack) peek() string {
	return (*stack)[len(*stack)-1]
}

func (stack *Stack) isEmpty() bool {
    return len(*stack) == 0
}

func (stack *Stack) size() int {
    return len(*stack)
}

func main() {
	s := make(Stack, 0)
	s.push("I")
    s.push("Am")
    s.push("batman")
	fmt.Printf("The stack looks like-\n%v\n", s)
    fmt.Printf("Popping the last element - %v\n", s.pop())
    fmt.Printf("Peeked at the top of the stack be like - %v\n", s.peek())
    fmt.Printf("Is the stack empty yet? - %v\n", s.isEmpty())
    fmt.Printf("The current size of the stack is - %v\n", s.size())
    s.push("BATMAN")
    fmt.Printf("Final stack after pushing looks like -\n%v\n", s)
}
