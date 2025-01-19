# Recursion and Backtracking in Golang

## Introduction to Recursion
Recursion is a programming technique where a function calls itself to solve a problem. It is a powerful tool for solving problems that can be broken down into smaller, similar subproblems. In Golang, a function can call itself by using the function's name.

### Examples and use cases
- Factorial calculation
- Fibonacci sequence
- Tower of Hanoi
- Binary search

### Coding Examples
#### Factorial Calculation
```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```
#### Fibonacci Sequence
```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```
#### Tower of Hanoi
```go
func hanoi(n int, from, to, aux string) {
    if n == 1 {
        fmt.Printf("Move disk 1 from rod %s to rod %s\n", from, to)
        return
    }
    hanoi(n-1, from, aux, to)
    fmt.Printf("Move disk %d from rod %s to rod %s\n", n, from, to)
    hanoi(n-1, aux, to, from)
}
```
#### Binary Search
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
    return -1
}
```

## Recursion with Arrays in Golang
Recursion can be used with arrays to solve problems that involve iterating over the elements of an array. In Golang, arrays are fixed-size sequences of elements of the same type. They are declared using the `var` keyword followed by the array name, the type of the elements, and the size of the array.

### Techniques and examples
- Array traversal
- Array manipulation
- Array sorting
- Array searching

### Coding Examples
#### Array Traversal
```go
func traverseArray(arr []int, index int) {
    if index >= len(arr) {
        return
    }
    fmt.Println(arr[index])
    traverseArray(arr, index+1)
}
```
#### Array Manipulation
```go
func sumArray(arr []int, index int) int {
    if index >= len(arr) {
        return 0
    }
    return arr[index] + sumArray(arr, index+1)
}
```
#### Array Sorting
```go
func quickSort(arr []int, low, high int) {
    if low < high {
        pi := partition(arr, low, high)
        quickSort(arr, low, pi-1)
        quickSort(arr, pi+1, high)
    }
}
```
#### Array Searching
```go
func searchArray(arr []int, target int, index int) int {
    if index >= len(arr) {
        return -1
    }
    if arr[index] == target {
        return index
    }
    return searchArray(arr, target, index+1)
}
```

## Backtracking in Golang
Backtracking is a general algorithmic technique that considers searching every possible combination in order to solve a computational problem. It is often used to solve problems that involve making a sequence of decisions. In Golang, backtracking can be implemented using recursion.

### Introduction and problem-solving
- N-Queens problem
- Sudoku solver
- Knights tour problem
- Rat in a maze problem

### Array and matrix problems 
- Subset sum problem
- Permutations of a string
- Combinations of a string
- Word search in a matrix

### Coding Examples
#### N-Queens Problem
```go
func solveNQueens(n int) [][]string {
    var result [][]string
    var board [][]string
    for i := 0; i < n; i++ {
        board = append(board, make([]string, n))
        for j := 0; j < n; j++ {
            board[i][j] = "."
        }
    }
    placeQueens(&result, board, 0, n)
    return result
}
```

#### Sudoku Solver
```go
func isValid(board [][]byte, row, col int, num byte) bool {
    for x := 0; x < 9; x++ {
        if board[row][x] == num {
            return false
        }
    }
    for x := 0; x < 9; x++ {
        if board[x][col] == num {
            return false
        }
    }
    startRow, startCol := row-row%3, col-col%3
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i+startRow][j+startCol] == num {
                return false
            }
        }
    }
    return true
}
```

#### Knight's Tour Problem
```go
func solveKTUtil(board [][]int, currX, currY, moveX, moveY []int, pos int) bool {
    var N = 8
    if pos == N*N {
        return true
    }
    for i := 0; i < 8; i++ {
        newX := currX + moveX[i]
        newY := currY + moveY[i]
        if isValid(newX, newY, N) && board[newX][newY] == -1 {
            board[newX][newY] = pos
            if solveKTUtil(board, newX, newY, moveX, moveY, pos+1) {
                return true
            } else {
                board[newX][newY] = -1
            }
        }
    }
    return false
}
```
#### Rat in a Maze Problem
```go
func solveMazeUtil(maze [][]int, x, y int, sol [][]int) bool {
    if x == len(maze)-1 && y == len(maze[0])-1 {
        sol[x][y] = 1
        return true
    }
    if isValid(maze, x, y) {
        sol[x][y] = 1
        if solveMazeUtil(maze, x+1, y, sol) {
            return true
        }
        if solveMazeUtil(maze, x, y+1, sol) {
            return true
        }
        sol[x][y] = 0
        return false
    }
    return false
}
```