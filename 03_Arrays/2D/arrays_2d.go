package main

import "fmt"

// Function to traverse a 2D array
func traverse2DArray(matrix [][]int) {
	for _, row := range matrix {
		for _, value := range row {
			fmt.Print(value, " ")
		}
		fmt.Println()
	}
}

// Function to perform matrix addition
func addMatrices(a, b [][]int) [][]int {
	result := make([][]int, len(a))
	for i := range a {
		result[i] = make([]int, len(a[i]))
		for j := range a[i] {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D Array Traversal:")
	traverse2DArray(matrix)

	matrixA := [][]int{
		{1, 2},
		{3, 4},
	}
	matrixB := [][]int{
		{5, 6},
		{7, 8},
	}
	fmt.Println("Matrix Addition:")
	result := addMatrices(matrixA, matrixB)
	traverse2DArray(result)
}
