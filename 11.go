package task

import (
	"fmt"
	"math/rand"
	"time"
)

// З1
func isEven(n int) bool {
	return n%2 == 0
}

// 2
func generateRandomPoints() (int, int, int, int, int, int) {
	rand.Seed(time.Now().UnixNano())
	x1, y1 := rand.Intn(100), rand.Intn(100)
	x2, y2 := rand.Intn(100), rand.Intn(100)
	x3, y3 := rand.Intn(100), rand.Intn(100)
	return x1, y1, x2, y2, x3, y3
}

// 3
func transformCoordinates(x, y int) (int, int) {
	newX := 2*x + 10
	newY := -3*y - 5
	return newX, newY
}

// 3
func multiply(n int) (result int) {
	result = n * 2
	return
}

func add(n int) (result int) {
	result = n + 10
	return
}

// 4
var globalVar1 int = 5
var globalVar2 int = 10
var globalVar3 int = 15

func addGlobal1(n int) int {
	return n + globalVar1
}

func addGlobal2(n int) int {
	return n + globalVar2
}

func addGlobal3(n int) int {
	return n + globalVar3
}

// Функция для запуска всех заданий
func Task11() {
	// 1
	number := rand.Intn(100)
	fmt.Printf("Число %d четное? %v\n", number, isEven(number))

	// 2
	x1, y1, x2, y2, x3, y3 := generateRandomPoints()
	fmt.Printf("Оригинальные точки: (%d, %d), (%d, %d), (%d, %d)\n", x1, y1, x2, y2, x3, y3)
	newX1, newY1 := transformCoordinates(x1, y1)
	newX2, newY2 := transformCoordinates(x2, y2)
	newX3, newY3 := transformCoordinates(x3, y3)
	fmt.Printf("Трансформированные точки: (%d, %d), (%d, %d), (%d, %d)\n", newX1, newY1, newX2, newY2, newX3, newY3)

	// 3
	num := 7
	multiplied := multiply(num)
	added := add(multiplied)
	fmt.Printf("Результат умножения: %d, Результат прибавления: %d\n", multiplied, added)

	// 4
	initial := 10
	result1 := addGlobal1(initial)
	result2 := addGlobal2(result1)
	result3 := addGlobal3(result2)
	fmt.Printf("Итоговый результат: %d\n", result3)
}
