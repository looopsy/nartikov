package main

import (
	"fmt"
	"strings"
)

func separateEvenOdd(numbers []int) ([]int, []int) {
	even := []int{}
	odd := []int{}

	for _, num := range numbers {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}

	return even, odd
}
func parseTest(sentences []string, chars []rune) [][]int {
	result := make([][]int, len(sentences))

	for i, sentence := range sentences {
		words := strings.Fields(sentence)
		lastWord := words[len(words)-1]

		indices := make([]int, len(chars))
		for j, char := range chars {
			index := strings.LastIndex(lastWord, string(char))
			if index != -1 {
				indices[j] = index
			}
		}

		result[i] = indices
	}

	return result
}

func main() {
	// Задание 1: Разделение чётных и нечётных чисел
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even, odd := separateEvenOdd(numbers)
	fmt.Println("Even numbers:", even)
	fmt.Println("Odd numbers:", odd)

	// Задание 2: Поиск символов в нескольких строках
	sentences := []string{
		"Hello world!",
		"This is a test .",
	}
	chars := []rune{'l', 'e'}

	result := parseTest(sentences, chars)
	fmt.Println("Indices of characters in the last words:")
	for _, indices := range result {
		fmt.Println(indices)
	}
}
