package main

import "fmt"

// Function to traverse a 1D array
func traverseArray(arr []int) {
	for _, value := range arr {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

// Function to find the maximum element in a 1D array
func findMax(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

// Two-pointer technique to find a pair with a given sum
func twoSum(arr []int, target int) (int, int) {
	left, right := 0, len(arr)-1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return left, right
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return -1, -1
}

// Sliding window technique to find the maximum sum of a subarray of size k
func maxSumSubarray(arr []int, k int) int {
	maxSum, windowSum := 0, 0
	for i := 0; i < len(arr); i++ {
		windowSum += arr[i]
		if i >= k-1 {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			windowSum -= arr[i-(k-1)]
		}
	}
	return maxSum
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Print("Array Traversal: ")
	traverseArray(arr)
	fmt.Println("Maximum Element:", findMax(arr))
	left, right := twoSum(arr, 5)
	fmt.Println("Two Sum Indices:", left, right)
	fmt.Println("Max Sum Subarray of size 2:", maxSumSubarray(arr, 2))
}
