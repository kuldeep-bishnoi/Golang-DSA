package main

import "fmt"

// Linear search function
func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1 // Return -1 if the target is not found
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	index := linearSearch(arr, target)
	fmt.Println("Index of", target, ":", index)
}
