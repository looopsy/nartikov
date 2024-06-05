package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomArray(size int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100) // Генерация случайных чисел от 0 до 99
	}
	return arr
}

func countNumbersAfter(arr []int, num int) int {
	count := 0
	found := false
	for _, value := range arr {
		if found {
			count++
		}
		if value == num {
			found = true
		}
	}
	if !found {
		return 0
	}
	return count
}

func findFirstIndex(arr []int, num int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == num && (mid == 0 || arr[mid-1] < num) {
			return mid
		} else if arr[mid] < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	// Задание 1: Подсчет чисел после заданного числа
	arr := generateRandomArray(10)
	fmt.Println("Generated Array:", arr)

	var num1 int
	fmt.Print("Введите номер: ")
	fmt.Scan(&num1)

	count := countNumbersAfter(arr, num1)
	fmt.Printf("Количество элементов после %d: %d\n", num1, count)

	// Задание 2: Поиск первого вхождения числа в упорядоченном массиве
	arr2 := []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var num2 int
	fmt.Print("Enter a number for task 2: ")
	fmt.Scan(&num2)

	index := findFirstIndex(arr2, num2)
	if index != -1 {
		fmt.Printf("Индекс первого появления %d: %d\n", num2, index)
	} else {
		fmt.Println("Номер не найден в массиве.")
	}
}
