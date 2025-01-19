# Basics of Golang

## Input and Output
- Golang provides various ways to perform input and output operations. The most common ones are fmt package and bufio package.
- The fmt package provides functions like Print, Printf, Println, Scan, Scanf, Scanln for standard input and output operations.
- The bufio package provides buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.

## Conditionals and Loops
- Golang supports if, else, and switch statements for decision-making.
- The for loop is the only loop statement in Golang. It can be used in three ways: with a single condition, with a post statement, and as a while loop.

## Data Types and Best Practices
- Golang has various primitive data types such as int, float, bool, string, etc.
- Best coding practices in Golang include using camelCase for variable names, using short variable declaration, and using the blank identifier (_) to discard values.

## Functions and Scoping
- Functions in Golang are defined using the func keyword. They can take zero or more arguments and return zero or more values.
- Golang has block scope. Variables declared inside a block are only accessible within that block.