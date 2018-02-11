package algorithm

func BinarySearchRecursive(array []int, element int) int {
	return binarySearchLeftRight(array, element, 0, len(array)-1)
}

func binarySearchLeftRight(array []int, element, left, right int) int {
	if left > right {
		return -1
	}
	middle := (left + right) / 2
	if element < array[middle] {
		return binarySearchLeftRight(array, element, left, middle-1)
	} else if element > array[middle] {
		return binarySearchLeftRight(array, element, middle+1, right)
	} else if element == array[middle] {
		return middle
	}
	return -1
}
