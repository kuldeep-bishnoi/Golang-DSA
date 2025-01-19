package inheritance

import "fmt"

// Person struct
type Person struct {
	Name string
	Age  int
}

// Employee struct embedding Person to demonstrate inheritance
type Employee struct {
	Person
	Department string
}

func main() {
	emp := Employee{
		Person:     Person{Name: "John Doe", Age: 30},
		Department: "Engineering",
	}
	fmt.Println("Employee Name:", emp.Name)
	fmt.Println("Employee Age:", emp.Age)
	fmt.Println("Employee Department:", emp.Department)
}
