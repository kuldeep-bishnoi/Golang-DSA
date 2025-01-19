package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Input and Output
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("You entered:", text)

	// Conditionals
	num := 10
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Switch statement
	switch num {
	case 10:
		fmt.Println("Number is 10")
	case 20:
		fmt.Println("Number is 20")
	default:
		fmt.Println("Number is neither 10 nor 20")
	}

	// Loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Data Types
	var integer int = 42
	var floatNum float64 = 3.14
	var boolean bool = true
	var str string = "Golang"
	fmt.Println("Integer:", integer)
	fmt.Println("Float:", floatNum)
	fmt.Println("Boolean:", boolean)
	fmt.Println("String:", str)

	// Best Practices
	camelCaseVar := "This is camelCase"
	fmt.Println(camelCaseVar)

	// Using blank identifier
	_, err := fmt.Println("This is a test")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Functions and Scoping
	result := add(5, 3)
	fmt.Println("Sum:", result)
}

// Function to add two numbers
func add(a int, b int) int {
	return a + b
}
