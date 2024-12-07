package main

import (
	"fmt"
	"strings"
)

// Basic function with parameters and return value
func add(a, b int) int {
	return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Named return values
func rectangle(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // naked return
}

// Variadic function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function as parameter
func calculate(operation func(int, int) int, a, b int) int {
	return operation(a, b)
}

// Closure (anonymous function)
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Method with receiver
type Person struct {
	FirstName string
	LastName  string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Pointer receiver method
func (p *Person) ChangeFirstName(newName string) {
	p.FirstName = newName
}

func main() {
	fmt.Println("=== Basic Function ===")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	fmt.Println("\n=== Multiple Return Values ===")
	if result, err := divide(10, 2); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	fmt.Println("\n=== Named Return Values ===")
	area, perimeter := rectangle(5, 3)
	fmt.Printf("Rectangle 5x3 - Area: %.2f, Perimeter: %.2f\n", area, perimeter)

	fmt.Println("\n=== Variadic Function ===")
	fmt.Printf("Sum of 1,2,3: %d\n", sum(1, 2, 3))
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of slice: %d\n", sum(numbers...))

	fmt.Println("\n=== Function as Parameter ===")
	multiply := func(a, b int) int { return a * b }
	fmt.Printf("Calculate multiply 4 * 5: %d\n", calculate(multiply, 4, 5))

	fmt.Println("\n=== Closure ===")
	increment := counter()
	fmt.Printf("Count: %d\n", increment())
	fmt.Printf("Count: %d\n", increment())
	fmt.Printf("Count: %d\n", increment())

	fmt.Println("\n=== Methods ===")
	person := Person{FirstName: "John", LastName: "Doe"}
	fmt.Printf("Full name: %s\n", person.FullName())
	
	person.ChangeFirstName("Jane")
	fmt.Printf("After name change: %s\n", person.FullName())

	// Defer example
	fmt.Println("\n=== Defer Example ===")
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")
}
