package main

import "fmt"

// Function to traverse a 3D array
func traverse3DArray(cube [][][]int) {
	for _, matrix := range cube {
		for _, row := range matrix {
			for _, value := range row {
				fmt.Print(value, " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func main() {
	cube := [][][]int{
		{
			{1, 2, 3},
			{4, 5, 6},
		},
		{
			{7, 8, 9},
			{10, 11, 12},
		},
	}
	fmt.Println("3D Array Traversal:")
	traverse3DArray(cube)
}
