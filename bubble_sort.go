package main

import "fmt"

func bubbleSort(array []int) []int {
	for i := range array {
		// last i elements are already in current positions
		// -1 => compare array[j] and array[j+1]
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				swap(j, j+1, array)
			}
		}
	}
	return array
}

// stop the algorithm if the array is sorted
func optimizedBubbleSort(array []int) []int {
	isSorted := true
	for i := range array {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				swap(j, j+1, array)
				// If any swap is occured, we need to continue to sort
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
	return array
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	array := []int{8, 2, 5, 4, 1}
	fmt.Println("sorted array: ", bubbleSort(array))
	fmt.Println("sorted array: ", optimizedBubbleSort(array))
}

/* output
sorted array:  [1 2 4 5 8]
sorted array:  [1 2 4 5 8]
*/
