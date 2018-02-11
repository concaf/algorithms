package algorithm

func BinarySearchIterative(array []int, element int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		middle := (right + left) / 2
		if element > array[middle] {
			left = middle + 1
		} else if element < array[middle] {
			right = middle - 1
		} else if element == array[middle] {
			return middle
		}
	}
	return -1
}
