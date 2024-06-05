package task

import (
	"fmt"
)

func sumOdds(start int, end int) {
	if start > end {
		start, end = end, start
	}

	sum := 0
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("Сумма четных чисел в диапазоне:", sum)
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func Task10() {
	// 1
	sumOdds(1, 10)

	// 2
	x := 10
	y := 20

	fmt.Printf("До: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("После: x = %d, y = %d\n", x, y)
}
