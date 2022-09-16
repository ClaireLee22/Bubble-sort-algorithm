def bubbleSort(array):
    for i in range(len(array)):
        # last i elements are already in current positions
        # -1 => compare array[j] and array[j+1]
        for j in range(len(array)-1-i):
            if array[j] > array[j+1]:
                swap(j, j+1, array)
    return array

# stop the algorithm if the array is sorted
def optimizedBubbleSort(array):
    isSorted = True
    for i in range(len(array)):
        for j in range(len(array)-1-i):
            if array[j] > array[j+1]:
                swap(j, j+1, array)
                # If any swap is occured, we need to continue to sort
                isSorted = False
        
        if isSorted:
            break

    return array


def swap(i, j , array):
    array[i], array[j] = array[j], array[i]

if __name__ == "__main__":
    array = [8, 2, 5, 4, 1]
    print("sorted array", bubbleSort(array))
    print("sorted array", optimizedBubbleSort(array))

"""
output:
('sorted array', [1, 2, 4, 5, 8])
('sorted array', [1, 2, 4, 5, 8])
"""