# Sorting Algorithms in Golang

## Basic Sorting Algorithms
### Bubble Sort
Bubble sort is a simple sorting algorithm that repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order. The pass through the list is repeated until the list is sorted.

**Golang Implementation:**
```go
func bubbleSort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr
}
```

### Selection Sort
Selection sort is a sorting algorithm that selects the smallest (or largest, depending on the ordering) element from the unsorted portion of the list and swaps it with the first element of the unsorted portion.

**Golang Implementation:**
```go
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
```

### Insertion Sort
Insertion sort is a simple sorting algorithm that builds the final sorted array (or list) one item at a time. It is much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort.

**Golang Implementation:**
```go
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
```

## Advanced Sorting Algorithms
### Merge Sort
Merge sort is a divide-and-conquer algorithm that divides the input array into two halves, calls itself for the two halves, and then merges the two sorted halves.

**Golang Implementation:**
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
    merged := []int{}
    for len(left) > 0 && len(right) > 0 {
        if left[0] <= right[0] {
            merged = append(merged, left[0])
            left = left[1:]
        } else {
            merged = append(merged, right[0])
            right = right[1:]
        }
    }
    merged = append(merged, left...)
    merged = append(merged, right...)
    return merged
}
```

### Quick Sort
Quick sort is a divide-and-conquer algorithm that selects a 'pivot' element from the array and partitions the other elements into two sub-arrays, according to whether they are less than or greater than the pivot.

**Golang Implementation:**
```go
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
```

### Heap Sort
Heap sort is a comparison-based sorting technique based on a Binary Heap data structure. It is similar to selection sort where we first find the maximum element and place the maximum element at the end. We repeat the same process for remaining element except the last element.

**Golang Implementation:**
```go
func heapSort(arr []int) []int {
    n := len(arr)
    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }
    for i := n - 1; i >= 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        heapify(arr, i, 0)
    }
    return arr
}

func heapify(arr []int, n int, i int) {
    largest := i
    left := 2*i + 1
    right := 2*i + 2
    if left < n && arr[left] > arr[largest] {
        largest = left
    }
    if right < n && arr[right] > arr[largest] {
        largest = right
    }
    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)
    }
}
```

### Radix Sort
Radix sort is a non-comparative integer sorting algorithm that sorts data with integer keys by grouping keys by the individual digits which share the same significant position and value.

**Golang Implementation:**
```go
func radixSort(arr []int) []int {
    max := getMax(arr)
    exp := 1
    for max/exp > 0 {
        countingSort(arr, exp)
        exp *= 10
    }
    return arr
}

func getMax(arr []int) int {
    max := arr[0]
    for _, v := range arr {
        if v > max {
            max = v
        }
    }
    return max
}

func countingSort(arr []int, exp int) {
    n := len(arr)
    output := make([]int, n)
    count := make([]int, 10)

    for i := 0; i < n; i++ {
        index := (arr[i] / exp) % 10
        count[index]++
    }

    for i := 1; i < 10; i++ {
        count[i] += count[i-1]
    }

    for i := n - 1; i >= 0; i-- {
        index := (arr[i] / exp) % 10
        output[count[index]-1] = arr[i]
        count[index]--
    }

    for i := 0; i < n; i++ {
        arr[i] = output[i]
    }
}
```

### Bucket Sort
Bucket sort is a comparison sort algorithm that operates on elements by dividing them into different buckets and then sorting these buckets individually.

**Golang Implementation:**
```go
func bucketSort(arr []int) []int {
    n := len(arr)
    if n <= 1 {
        return arr
    }
    buckets := make([][]int, n)
    for i := 0; i < n; i++ {
        index := arr[i] * n / (arr[n-1] + 1)
        buckets[index] = append(buckets[index], arr[i])
    }
    for i := 0; i < n; i++ {
        buckets[i] = insertionSort(buckets[i])
    }
    index := 0
    for i := 0; i < n; i++ {
        for j := 0; j < len(buckets[i]); j++ {
            arr[index] = buckets[i][j]
            index++
        }
    }
    return arr
}
```

### Shell Sort
Shell sort is a comparison sort algorithm that generalizes insertion sort by allowing the comparison and exchange of far-apart elements.

**Golang Implementation:**
```go
func shellSort(arr []int) []int {
    n := len(arr)
    gap := n / 2
    for gap > 0 {
        for i := gap; i < n; i++ {
            temp := arr[i]
            j := i
            for j >= gap && arr[j-gap] > temp {
                arr[j] = arr[j-gap]
                j -= gap
            }
            arr[j] = temp
        }
        gap /= 2
    }
    return arr
}
```

## Comparison of Sorting Algorithms
### Time and Space Complexity
| Algorithm      | Time Complexity | Space Complexity |
| Bubble Sort    | O(n^2)          | O(1)             |
| Selection Sort | O(n^2)          | O(1)             |
| Insertion Sort | O(n^2)          | O(1)             |
| Merge Sort     | O(n log n)      | O(n)             |
| Quick Sort     | O(n log n)      | O(log n)         |
| Heap Sort      | O(n log n)      | O(1)             |
| Radix Sort     | O(nk)           | O(nk)            |
| Bucket Sort    | O(n + k)        | O(n + k)         |
| Shell Sort     | O(n log n)      | O(1)             |

### Use Cases
- Bubble sort is suitable for small data sets where performance is not a critical factor.
- Selection sort is suitable for small data sets where performance is not a critical factor.
- Insertion sort is suitable for small data sets where performance is not a critical factor.
- Merge sort is suitable for large data sets where stability is important.
- Quick sort is suitable for large data sets where performance is critical.
- Heap sort is suitable for large data sets where performance is critical.
- Radix sort is suitable for sorting large datasets of integers or strings.
- Bucket sort is suitable for sorting floating-point numbers or strings.
- Shell sort is suitable for sorting large datasets where performance is critical.