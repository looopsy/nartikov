package task

import "fmt"

// Determinant3x3 calculates the determinant of a 3x3 matrix
func Determinant3x3(matrix [3][3]float64) float64 {
	return matrix[0][0]*(matrix[1][1]*matrix[2][2]-matrix[1][2]*matrix[2][1]) -
		matrix[0][1]*(matrix[1][0]*matrix[2][2]-matrix[1][2]*matrix[2][0]) +
		matrix[0][2]*(matrix[1][0]*matrix[2][1]-matrix[1][1]*matrix[2][0])
}

// MultiplyMatrices multiplies two matrices
func MultiplyMatrices(a [3][5]float64, b [5][4]float64) [3][4]float64 {
	var result [3][4]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = 0
			for k := 0; k < 5; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

func Task14() {
	matrix := [3][3]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	det := Determinant3x3(matrix)
	fmt.Println("Determinant:", det)

	a := [3][5]float64{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}
	b := [5][4]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}
	product := MultiplyMatrices(a, b)
	fmt.Println("Matrix product:", product)
}
