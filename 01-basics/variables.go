package main

import "fmt"

func main() {
	// Basic variable declarations
	var name string = "John"
	var age int = 25
	salary := 50000.0 // Type inference

	// Multiple variable declaration
	var (
		isEmployed bool   = true
		company    string = "TechCorp"
	)

	// Constants
	const PI = 3.14159

	// Basic data types demonstration
	var (
		integerNum int     = 42
		floatNum   float64 = 3.14
		text       string  = "Hello, Go!"
		isTrue     bool    = true
	)

	// Print variables
	fmt.Println("=== Basic Variables ===")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Salary: %.2f\n", salary)
	fmt.Printf("Employed: %t\n", isEmployed)
	fmt.Printf("Company: %s\n", company)

	fmt.Println("\n=== Data Types ===")
	fmt.Printf("Integer: %d (Type: %T)\n", integerNum, integerNum)
	fmt.Printf("Float: %f (Type: %T)\n", floatNum, floatNum)
	fmt.Printf("String: %s (Type: %T)\n", text, text)
	fmt.Printf("Boolean: %t (Type: %T)\n", isTrue, isTrue)

	// Zero values
	var (
		defaultInt    int
		defaultFloat  float64
		defaultString string
		defaultBool   bool
	)

	fmt.Println("\n=== Zero Values ===")
	fmt.Printf("Default Integer: %d\n", defaultInt)
	fmt.Printf("Default Float: %f\n", defaultFloat)
	fmt.Printf("Default String: %q\n", defaultString)
	fmt.Printf("Default Boolean: %t\n", defaultBool)
}
