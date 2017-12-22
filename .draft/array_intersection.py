def array_intersection(arr1, arr2):
    intersection = []
    for element in arr1:
        if element in arr2:
            intersection.append(element)
    return intersection

arr1 = [1,2,3,4,5,6,7,8,9]
arr2 = [2,3,5,7,12,32,563]

print(array_intersection(arr1, arr2))