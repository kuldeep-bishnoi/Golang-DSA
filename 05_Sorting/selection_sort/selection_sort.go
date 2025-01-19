package main

import "fmt"

// Selection Sort
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Selection Sorted:", selectionSort(append([]int{}, arr...)))
}
