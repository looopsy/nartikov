package task

import "fmt"

// MergeSortedArrays merges two sorted arrays into one sorted array
func MergeSortedArrays(arr1, arr2 []int) []int {
	merged := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged[k] = arr1[i]
			i++
		} else {
			merged[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr1) {
		merged[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		merged[k] = arr2[j]
		j++
		k++
	}
	return merged
}

// BubbleSort sorts an array using bubble sort algorithm
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func Task13() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8, 9}
	merged := MergeSortedArrays(arr1, arr2)
	fmt.Println("Merged array:", merged)

	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", arr)
	BubbleSort(arr)
	fmt.Println("Sorted array:", arr)
}
