def reverse_array(array):
    array_length = len(array)
    for index in range(array_length/2):
        array[index], array[array_length-index-1] = \
            array[array_length-index-1], array[index]
    return array

if __name__ == "__main__":
    arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    reversed_output = reverse_array(arr)
    print(reversed_output)
