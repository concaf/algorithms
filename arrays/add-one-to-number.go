//Given a non-negative number represented as an array of digits,
//
//add 1 to the number ( increment the number represented by the digits ).
//
//The digits are stored such that the most significant digit is at the head of the list.
//
//Example:
//
//If the vector has [1, 2, 3]
//
//the returned vector should be [1, 2, 4]
//
//as 123 + 1 = 124.
//
//Q : Can the input have 0’s before the most significant digit. Or in other words, is 0 1 2 3 a valid input?
//A : For the purpose of this question, YES
//Q : Can the output have 0’s before the most significant digit? Or in other words, is 0 1 2 4 a valid output?
//A : For the purpose of this question, NO. Even if the input has zeroes before the most significant digit.

package main

import "fmt"

/**
 * @input A : Integer array
 *
 * @Output Integer array.
 */
func plusOne(A []int) []int {
	A = removeTrailingZeroes(A)

	lastPos := len(A) - 1
	A = addOne(A, lastPos)

	return A
}

func addOne(A []int, pos int) []int {
	if A[pos] == 9 {
		A[pos] = 0
		if len(A) == 1 {
			A = append([]int{1}, A...)
		} else {
			A = append(plusOne(A[:pos]), A[pos])
		}
	} else {
		A[pos]++
	}
	return A
}

func removeTrailingZeroes(array []int) []int {
	if len(array) == 1 {
		return array
	}

	var i int
	for i = 0; array[i] == 0; i++ {
	}
	return array[i:]
}

func main() {
	input1 := []int{9, 9, 9, 9, 9}
	fmt.Println(input1)
	fmt.Println(plusOne(input1))

	input2 := []int{0, 3, 7, 6, 4, 0, 5, 5, 5}
	fmt.Println(input2)
	fmt.Println(plusOne(input2))

	input3 := []int{0}
	fmt.Println(input3)
	fmt.Println(plusOne(input3))

}
