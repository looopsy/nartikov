package main

import "fmt"

// Задание 1: Сортировка вставками
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

// Задание 2: Анонимные функции
func bubbleSortReverse(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// Задание 1: Сортировка вставками
	arr1 := []int{5, 3, 8, 1, 2, 7, 4, 10, 9, 6}
	fmt.Println("Исходный массив для сортировки вставкой:", arr1)
	insertionSort(arr1)
	fmt.Println("Отсортированный массив с использованием сортировки вставкой:", arr1)

	// Задание 2: Анонимные функции
	bubbleSortReverseFunc := func(arr ...int) {
		bubbleSortReverse(arr)
		fmt.Println("Отсортированный массив в обратном порядке с использованием пузырьковой сортировки:", arr)
	}

	arr2 := []int{5, 3, 8, 1, 2, 7, 4, 10, 9, 6}
	fmt.Println("Исходный массив для обратной пузырьковой сортировки:", arr2)
	bubbleSortReverseFunc(arr2...)
}
