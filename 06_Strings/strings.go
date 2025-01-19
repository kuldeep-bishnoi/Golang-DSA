package main

import (
	"fmt"
)

func main() {
	// String Concatenation
	str1 := "Hello"
	str2 := "World"
	concatenatedStr := str1 + " " + str2
	fmt.Println("Concatenated String:", concatenatedStr)

	// String Slicing
	str := "Hello, World!"
	fmt.Println("Sliced String:", str[0:5])

	// String Reversal
	reversedStr := reverseString(str)
	fmt.Println("Reversed String:", reversedStr)
}

// Function to reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
