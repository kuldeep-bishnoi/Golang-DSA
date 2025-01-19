package main

import "fmt"

// Binary search function
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // Return -1 if the target is not found
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	index := binarySearch(arr, target)
	fmt.Println("Index of", target, ":", index)
}
