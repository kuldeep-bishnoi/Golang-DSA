# Searching Algorithms in Golang

## Linear Search
- **Basic Concept**: Linear search is a simple searching algorithm that sequentially checks each element of a list until a match is found or until all the elements have been searched.
- **Implementation in Golang**:
```go
func linearSearch(arr []int, target int) int {
    for i, v := range arr {
        if v == target {
            return i
        }
    }
    return -1 // Return -1 if the target is not found
}
```
### Example Usage
```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5}
    target := 3
    index := linearSearch(arr, target)
    fmt.Println("Index of", target, ":", index)
}
```

## Binary Search
- **Basics**: Binary search is a fast search algorithm that finds the position of a target value within a sorted array. It compares the target value to the middle element of the array and eliminates half of the array in each iteration.
- **Variations**: Binary search can be used for both sorted arrays and linked lists. It can also be used for searching in rotated sorted arrays.
- **Searching in 2D arrays**: Binary search can be extended to search in 2D arrays by treating the 2D array as a 1D array and applying the standard binary search algorithm.
- **Implementation in Golang**:
```go
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
```
### Example Usage
```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5}
    target := 3
    index := binarySearch(arr, target)
    fmt.Println("Index of", target, ":", index)
}
```

## Advanced Searching Techniques
- **Jump Search**: Jump search is a searching algorithm for ordered lists. It creates a block and tries to find the element in that block. If the item is not in the block, it shifts the entire block. The block size is based on the size of the list.
- **Interpolation Search**: Interpolation search is an improved variant of binary search for instances where the values in a sorted array are uniformly distributed. It estimates the probable position of the target value based on the values at the start and end of the array.
- **Implementation in Golang**:
```go
func jumpSearch(arr []int, target int) int {
    step := int(math.Sqrt(float64(len(arr))))
    prev := 0
    for arr[min(step, len(arr)-1)] < target {
        prev = step
        step += int(math.Sqrt(float64(len(arr))))
        if prev >= len(arr) {
            return -1
        }
    }
    while arr[prev] < target {
        prev++
        if prev == min(step, len(arr)) {
            return -1
        }
    }
    if arr[prev] == target {
        return prev
    }
    return -1
}

func interpolationSearch(arr []int, target int) int {
    low, high := 0, len(arr)-1
    for low <= high && target >= arr[low] && target <= arr[high] {
        if low == high {
            if arr[low] == target {
                return low
            }
            return -1
        }
        pos := low + ((high-low)/(arr[high]-arr[low]))*(target-arr[low])
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
```
### Example Usage
```go
package main

import (
    "fmt"
    "math"
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    arr := []int{1, 2, 3, 4, 5}
    target := 3
    indexJump := jumpSearch(arr, target)
    indexInterpolation := interpolationSearch(arr, target)
    fmt.Println("Index of", target, "using Jump Search:", indexJump)
    fmt.Println("Index of", target, "using Interpolation Search:", indexInterpolation)
}
```
## Time and Space Complexity
| Algorithm      | Time Complexity | Space Complexity |
| Linear Search  | O(n)            | O(1)             |
| Binary Search  | O(log n)        | O(1)             |
| Jump Search    | O(sqrt(n))      | O(1)             |
| Interpolation Search | O(log log n) | O(1)             |
