# Arrays in Golang

## Simple Explanation

Think of an array like a row of mailboxes in an apartment building. Each mailbox has a unique number (index), and you can put something inside each one (a value). To find something, you need to know which mailbox number to look in.

```
Index:   0    1    2    3    4
       ┌────┬────┬────┬────┬────┐
Values: │ 10 │ 20 │ 30 │ 40 │ 50 │
       └────┴────┴────┴────┴────┘
```

Arrays in Go have a fixed size, which means once you create them, you can't add more mailboxes. If you need a flexible size, you'll want to use slices (which we'll also cover here).

## Key Concepts

- **Fixed Size**: Arrays in Go have a fixed length that cannot change after declaration
- **Zero-Based Indexing**: The first element is at index 0, not 1
- **Homogeneous Elements**: All elements must be of the same type
- **Memory Allocation**: Arrays are stored in contiguous memory locations
- **Slices**: Go's dynamic arrays, built on top of arrays with flexible size

## Basic Implementation in Go

### Arrays

```go
package main

import "fmt"

func main() {
    // Declare and initialize a fixed-size array
    var grades [5]int = [5]int{95, 89, 76, 93, 100}
    
    // Shorter initialization
    scores := [5]int{95, 89, 76, 93, 100}
    
    // Let compiler count the elements
    fibonacci := [...]int{1, 1, 2, 3, 5, 8, 13}
    
    // Access elements
    fmt.Println("First grade:", grades[0])
    
    // Modify elements
    scores[2] = 80
    
    // Print arrays
    fmt.Println("Grades:", grades)
    fmt.Println("Scores:", scores)
    fmt.Println("Fibonacci:", fibonacci)
    fmt.Println("Length of fibonacci:", len(fibonacci))
}
```

### Slices (Dynamic Arrays)

```go
package main

import "fmt"

func main() {
    // Create a slice
    numbers := []int{1, 2, 3, 4, 5}
    
    // Append elements to a slice
    numbers = append(numbers, 6, 7, 8)
    
    // Create a slice with make(type, length, capacity)
    scores := make([]int, 5, 10)
    
    // Slicing operations
    someNumbers := numbers[1:4]  // Elements 1,2,3 (indices 1,2,3)
    
    fmt.Println("Numbers:", numbers)
    fmt.Println("Some numbers:", someNumbers)
    fmt.Println("Length:", len(numbers))
    fmt.Println("Capacity:", cap(numbers))
}
```

## Time and Space Complexity

| Operation | Array | Slice | Time Complexity | Space Complexity | Explanation |
|-----------|-------|-------|----------------|-----------------|-------------|
| Access    | `arr[i]` | `slice[i]` | O(1) | O(1) | Direct memory address calculation |
| Insert at end | N/A (fixed size) | `append(slice, x)` | O(1)* | O(1) | Amortized constant time due to occasional reallocation |
| Insert at index | N/A (fixed size) | `append(slice[:i], append([]T{x}, slice[i:]...)...)` | O(n) | O(n) | Need to shift all elements after i |
| Delete from end | N/A (fixed size) | `slice = slice[:len(slice)-1]` | O(1) | O(1) | Simple slice reslicing |
| Delete from index | N/A (fixed size) | `slice = append(slice[:i], slice[i+1:]...)` | O(n) | O(n) | Need to shift all elements after i |
| Search | Linear search | Linear search | O(n) | O(1) | Need to check each element in worst case |
| Sort | Various algorithms | Various algorithms | O(n log n) | Varies | Typically uses sort.Ints() |

*Amortized O(1) means occasional operations might take longer, but on average it's constant time.

## 🧩 Common Patterns and Techniques

### Two-Pointer Technique
- Uses two pointers to solve array problems efficiently
- Often used for searching pairs in a sorted array
- Common in problems like finding pairs with a given sum, detecting palindromes, etc.

```go
// Find a pair that sums to target (works on sorted array)
func findPair(nums []int, target int) (int, int) {
    left, right := 0, len(nums)-1
    
    for left < right {
        currentSum := nums[left] + nums[right]
        
        if currentSum == target {
            return left, right
        } else if currentSum < target {
            left++
        } else {
            right--
        }
    }
    
    return -1, -1  // Pair not found
}
```

### Sliding Window Technique
- Converts some O(n²) algorithms to O(n)
- Used for problems involving contiguous subarrays
- Common in problems like finding maximum sum subarray of size k, longest substring with distinct chars

```go
// Maximum sum subarray of size k
func maxSumSubarray(arr []int, k int) int {
    n := len(arr)
    if n < k {
        return -1
    }
    
    // Calculate sum of first window
    maxSum := 0
    windowSum := 0
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }
    
    maxSum = windowSum
    
    // Slide window and update max sum
    for i := k; i < n; i++ {
        windowSum = windowSum + arr[i] - arr[i-k]
        if windowSum > maxSum {
            maxSum = windowSum
        }
    }
    
    return maxSum
}
```

## Real-World Applications

1. **Image Processing**: Images are represented as 2D arrays of pixels
2. **Spreadsheet Applications**: Data in Excel-like applications is stored in 2D arrays
3. **Database Systems**: Records are often stored and manipulated as arrays
4. **Audio Processing**: Sound samples are stored in arrays for processing
5. **Gaming**: Game boards (chess, tic-tac-toe) are represented as 2D arrays
6. **Machine Learning**: Feature vectors and matrices for calculations

## 🎯 Common Interview Questions

### Question 1: Two Sum

**Problem**: Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to target.

**Solution Approach**:
1. Use a hash map to store values we've seen and their indices
2. For each element, check if its complement (target - current element) exists in the map
3. If found, return the current index and the stored index

```go
func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)  // value -> index
    
    for i, num := range nums {
        complement := target - num
        if j, found := seen[complement]; found {
            return []int{j, i}
        }
        seen[num] = i
    }
    
    return nil
}
```

**Time Complexity**: O(n)
**Space Complexity**: O(n)

### Question 2: Maximum Subarray

**Problem**: Find the contiguous subarray within an array which has the largest sum.

**Solution Approach**:
1. Use Kadane's algorithm
2. Keep track of maximum sum ending at current position and global maximum
3. For each element, decide whether to start a new subarray or extend the existing one

```go
func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    maxSoFar := nums[0]
    maxEndingHere := nums[0]
    
    for i := 1; i < len(nums); i++ {
        // Either take current element alone or add to previous subarray
        maxEndingHere = max(nums[i], maxEndingHere+nums[i])
        maxSoFar = max(maxSoFar, maxEndingHere)
    }
    
    return maxSoFar
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

**Time Complexity**: O(n)
**Space Complexity**: O(1)

## Practice Problems

1. **Easy**: [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) - Check if array contains duplicates
2. **Easy**: [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) - Find the maximum profit
3. **Medium**: [Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/) - Calculate product of all elements except self
4. **Medium**: [3Sum](https://leetcode.com/problems/3sum/) - Find all unique triplets that sum to zero
5. **Hard**: [Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) - Calculate how much water can be trapped

## Additional Resources

- [Go Tour: Arrays](https://tour.golang.org/moretypes/6) - Official Go tutorial on arrays
- [Go Tour: Slices](https://tour.golang.org/moretypes/7) - Official Go tutorial on slices
- [Effective Go: Slices](https://golang.org/doc/effective_go#slices) - Go best practices for slices
- [Go by Example: Arrays](https://gobyexample.com/arrays) - Practical examples of arrays in Go
- [Go by Example: Slices](https://gobyexample.com/slices) - Practical examples of slices in Go

## Exercises

1. Implement a function that reverses an array in-place
2. Create a function that merges two sorted arrays into a single sorted array
3. Write a program that rotates an array by k positions
4. Implement a function that finds the majority element in an array
5. Create a function that removes duplicates from a sorted array
