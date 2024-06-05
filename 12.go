package task

import (
	"fmt"
)

// 1
func countEvenOdd(numbers []int) (int, int) {
	evenCount, oddCount := 0, 0
	for _, num := range numbers {
		if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	return evenCount, oddCount
}

// 2
func reverseArray(arr []int) []int {
	length := len(arr)
	reversed := make([]int, length)
	for i, v := range arr {
		reversed[length-1-i] = v
	}
	return reversed
}

func Task12() {
	// 1
	var numbers [10]int
	fmt.Println("Введите 10 целых чисел:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&numbers[i])
	}

	evenCount, oddCount := countEvenOdd(numbers[:])
	fmt.Printf("Чётных чисел: %d, Нечётных чисел: %d\n", evenCount, oddCount)

	// 2
	reversedNumbers := reverseArray(numbers[:])
	fmt.Println("Реверсированный массив:", reversedNumbers)
}
