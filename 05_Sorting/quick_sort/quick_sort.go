package main

import (
	"fmt"
	"math/rand"
)

// Quick Sort
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr)
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quickSort(arr[:left])
	quickSort(arr[left+1:])
	return arr
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Quick Sorted:", quickSort(append([]int{}, arr...)))
}
