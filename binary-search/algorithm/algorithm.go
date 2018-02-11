package algorithm

func BinarySearch(array []int, element int) int {
    length := len(array)
    if element > array[length-1] || element < array[0] {
        return -1
    }
    if element < array[length/2] {
        return BinarySearch(array[:length/2], element)
    } else if element > array[length/2] {
        return length/2 + BinarySearch(array[length/2:], element)
    } else if element == array[length/2] {
        return length/2
    }
    return -1
}