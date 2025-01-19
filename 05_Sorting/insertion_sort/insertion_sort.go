package main

import "fmt"

// Insertion Sort
func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Insertion Sorted:", insertionSort(append([]int{}, arr...)))
}
