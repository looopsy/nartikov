package task

import "fmt"

// CalculateFormula calculates the formula S = 2 * x + y^2 - 3 / z
func CalculateFormula(x int16, y uint8, z float32) float32 {
	return 2*float32(x) + float32(y)*float32(y) - 3/z
}

// WrapperFunction takes a function and calls it using defer
func WrapperFunction(f func(int, int) int, a int, b int) int {
	defer f(a, b)
	return f(a, b)
}

func Task15() {
	x := int16(3)
	y := uint8(4)
	z := float32(5.0)
	result := CalculateFormula(x, y, z)
	fmt.Println("Calculated result:", result)

	f1 := func(a int, b int) int {
		fmt.Println("Function 1 executed")
		return a + b
	}

	f2 := func(a int, b int) int {
		fmt.Println("Function 2 executed")
		return a - b
	}

	f3 := func(a int, b int) int {
		fmt.Println("Function 3 executed")
		return a * b
	}

	fmt.Println("Wrapper Function 1:", WrapperFunction(f1, 5, 3))
	fmt.Println("Wrapper Function 2:", WrapperFunction(f2, 5, 3))
	fmt.Println("Wrapper Function 3:", WrapperFunction(f3, 5, 3))
}
