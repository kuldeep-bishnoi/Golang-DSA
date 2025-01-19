package main

import "fmt"

// Factorial using recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Fibonacci using recursion
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// N-Queens problem using backtracking
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

func placeQueens(result *[][]string, board [][]string, row, n int) {
	if row == n {
		var solution []string
		for _, r := range board {
			solution = append(solution, fmt.Sprint(r))
		}
		*result = append(*result, solution)
		return
	}
	for col := 0; col < n; col++ {
		if isValid(board, row, col, n) {
			board[row][col] = "Q"
			placeQueens(result, board, row+1, n)
			board[row][col] = "."
		}
	}
}

func isValid(board [][]string, row, col, n int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Factorial of 5:", factorial(5))
	fmt.Println("Fibonacci of 5:", fibonacci(5))
	fmt.Println("N-Queens solutions for 4 queens:", solveNQueens(4))
}
