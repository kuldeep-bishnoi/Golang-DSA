package main

import "fmt"

// Interpolation search function
func interpolationSearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high && target >= arr[low] && target <= arr[high] {
		if low == high {
			if arr[low] == target {
				return low
			}
			return -1
		}
		pos := low + int(float64(high-low)/(float64(arr[high]-arr[low]))*float64(target-arr[low]))
		if arr[pos] == target {
			return pos
		}
		if arr[pos] < target {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	fmt.Println("Interpolation Search Index:", interpolationSearch(arr, target))
}
