# Sorting Algorithms in Golang

## 📖 Simple Explanation

Think of sorting like arranging a deck of numbered cards in order. Each sorting algorithm has a different strategy, similar to how different people might arrange the cards:

- **Bubble Sort**: Compare adjacent cards, swapping them if needed, making the largest "bubble up" to the end.
- **Selection Sort**: Find the smallest card in the unsorted pile and move it to your sorted pile.
- **Insertion Sort**: Like sorting cards in your hand, taking one card at a time and inserting it in the right position.
- **Merge Sort**: Split the deck in half, sort each half, then merge them back together.
- **Quick Sort**: Pick a "pivot" card, put smaller cards to the left, larger to the right, then sort the left and right piles.

```
Unsorted: [5] [2] [9] [1] [5] [6]
Sorted:   [1] [2] [5] [5] [6] [9]
```

## 🔑 Key Concepts

- **Stability**: Whether equal elements maintain their original order
- **In-Place vs. Out-of-Place**: Whether sorting happens with minimal extra memory
- **Comparison vs. Non-Comparison**: Whether sorting relies on comparing elements
- **Adaptive**: Whether the algorithm performs better on partially sorted data
- **Time and Space Complexity**: Efficiency considerations for different scenarios

## 💻 Basic Implementation in Go

### Bubble Sort

```go
func bubbleSort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        swapped := false
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = true
            }
        }
        if !swapped {
            break  // Array is sorted
        }
    }
    return arr
}
```

### Merge Sort

```go
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    
    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])
    
    return merge(left, right)
}

func merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))
    i, j := 0, 0
    
    for i < len(left) && j < len(right) {
        if left[i] <= right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }
    
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)
    
    return result
}
```

### Quick Sort

```go
func quickSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    
    pivot := arr[len(arr)/2]
    var left, middle, right []int
    
    for _, x := range arr {
        if x < pivot {
            left = append(left, x)
        } else if x == pivot {
            middle = append(middle, x)
        } else {
            right = append(right, x)
        }
    }
    
    left = quickSort(left)
    right = quickSort(right)
    
    return append(append(left, middle...), right...)
}
```

## ⏱️ Time and Space Complexity

| Algorithm | Best Time | Average Time | Worst Time | Space | Stable | In-Place | Adaptive | Notes |
|-----------|-----------|--------------|------------|-------|--------|----------|----------|-------|
| Bubble Sort | O(n) | O(n²) | O(n²) | O(1) | Yes | Yes | Yes | Simple, inefficient for large datasets |
| Selection Sort | O(n²) | O(n²) | O(n²) | O(1) | No | Yes | No | Simple, performs well on small datasets |
| Insertion Sort | O(n) | O(n²) | O(n²) | O(1) | Yes | Yes | Yes | Efficient for small or nearly sorted data |
| Merge Sort | O(n log n) | O(n log n) | O(n log n) | O(n) | Yes | No | No | Consistent performance, uses extra space |
| Quick Sort | O(n log n) | O(n log n) | O(n²) | O(log n) | No | Yes* | No | Fast in practice, pivot selection is crucial |
| Heap Sort | O(n log n) | O(n log n) | O(n log n) | O(1) | No | Yes | No | In-place but slower than quicksort in practice |
| Radix Sort | O(nk) | O(nk) | O(nk) | O(n+k) | Yes | No | No | Non-comparison sort, good for integers |

*Quick sort's space complexity is O(log n) for the recursive call stack.

## 🧩 Common Patterns and Techniques

### Divide and Conquer
- Breaks the problem into smaller subproblems
- Solves each subproblem recursively
- Combines the solutions
- Examples: Merge Sort, Quick Sort

```go
// Divide and conquer pattern
func divideAndConquerSort(arr []int) []int {
    // Base case
    if len(arr) <= 1 {
        return arr
    }
    
    // Divide
    mid := len(arr) / 2
    left := divideAndConquerSort(arr[:mid])
    right := divideAndConquerSort(arr[mid:])
    
    // Conquer (combine)
    return combine(left, right)
}
```

### In-Place Sorting
- Minimizes memory usage by modifying the input array directly
- Uses swapping to rearrange elements
- Examples: Bubble Sort, Selection Sort, Insertion Sort, Heap Sort

```go
// In-place sorting pattern (Selection Sort example)
func selectionSort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        // Swap in-place
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
    return arr
}
```

## 🔍 Real-World Applications

1. **Database Systems**: Sorting records for efficient retrieval
2. **Search Algorithms**: Sorting data before applying binary search
3. **File Systems**: Organizing files by name, date, or size
4. **Compression Algorithms**: Sorting symbols by frequency
5. **Graphics Rendering**: Sorting objects by depth (Painter's algorithm)
6. **Priority Scheduling**: Operating system task scheduling
7. **Data Analysis**: Sorting datasets for statistical operations

## 🎯 Common Interview Questions

### Question 1: Sort Colors (Dutch National Flag Problem)

**Problem**: Given an array with n objects colored red, white, or blue (represented as 0, 1, and 2), sort them in-place so that objects of the same color are adjacent.

**Solution Approach**:
1. Use the Dutch National Flag algorithm with three pointers
2. Maintain low, mid, and high pointers
3. Elements < 1 go before low, elements = 1 go between low and high, elements > 1 go after high

```go
func sortColors(nums []int) {
    low, mid, high := 0, 0, len(nums)-1
    
    for mid <= high {
        switch nums[mid] {
        case 0: // Red
            nums[low], nums[mid] = nums[mid], nums[low]
            low++
            mid++
        case 1: // White
            mid++
        case 2: // Blue
            nums[mid], nums[high] = nums[high], nums[mid]
            high--
        }
    }
}
```

**Time Complexity**: O(n)
**Space Complexity**: O(1)

### Question 2: Merge Intervals

**Problem**: Given an array of intervals, merge all overlapping intervals and return the non-overlapping intervals.

**Solution Approach**:
1. Sort the intervals by their start times
2. Iterate through the sorted intervals
3. Merge overlapping intervals by updating the end time

```go
func mergeIntervals(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }
    
    // Sort intervals by start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    result := [][]int{intervals[0]}
    
    for i := 1; i < len(intervals); i++ {
        current := intervals[i]
        last := result[len(result)-1]
        
        // Check if current interval overlaps with the last one in result
        if current[0] <= last[1] {
            // Merge by updating the end time
            if current[1] > last[1] {
                last[1] = current[1]
            }
        } else {
            // No overlap, add current interval to result
            result = append(result, current)
        }
    }
    
    return result
}
```

**Time Complexity**: O(n log n) due to sorting
**Space Complexity**: O(n) for the result array

## 💪 Practice Problems

1. **Easy**: [Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/) - Sort the squares of integers
2. **Medium**: [Kth Largest Element](https://leetcode.com/problems/kth-largest-element-in-an-array/) - Find the kth largest element without sorting
3. **Medium**: [Sort List](https://leetcode.com/problems/sort-list/) - Sort a linked list in O(n log n) time
4. **Hard**: [Count of Smaller Numbers After Self](https://leetcode.com/problems/count-of-smaller-numbers-after-self/) - Count smaller elements on the right
5. **Hard**: [First Missing Positive](https://leetcode.com/problems/first-missing-positive/) - Find the smallest missing positive integer

## 📚 Additional Resources

- [Visualgo Sorting](https://visualgo.net/en/sorting) - Visual animations of sorting algorithms
- [Go Sort Package](https://pkg.go.dev/sort) - Go's standard library sort package
- [Big-O Cheat Sheet](https://www.bigocheatsheet.com/) - Time and space complexity reference
- [Sorting Algorithms Animations](https://www.toptal.com/developers/sorting-algorithms) - Interactive animations
- [Tech Interview Handbook: Sorting](https://www.techinterviewhandbook.org/algorithms/sorting/) - Interview preparation for sorting questions

## 📝 Exercises

1. Implement a custom sorting function in Go to sort structures by multiple fields
2. Implement a hybrid sorting algorithm that uses insertion sort for small arrays and merge sort for larger ones
3. Create a visualization of different sorting algorithms using Go's standard library
4. Implement an external sorting algorithm for data that doesn't fit in memory
5. Write a function to sort an almost sorted array (where each element is at most k positions away from its sorted position) in O(n log k) time