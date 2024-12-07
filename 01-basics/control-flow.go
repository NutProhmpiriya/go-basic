package main

import "fmt"

func main() {
	// If-else statement
	fmt.Println("=== If-Else Examples ===")
	age := 18
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// If with initialization
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// For loop examples
	fmt.Println("\n=== Loop Examples ===")
	
	// Basic for loop
	fmt.Println("Basic for loop:")
	for i := 0; i < 3; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// For loop as while
	fmt.Println("\nWhile-style loop:")
	count := 0
	for count < 3 {
		fmt.Printf("Count: %d\n", count)
		count++
	}

	// Range-based for loop
	fmt.Println("\nRange-based loop:")
	fruits := []string{"apple", "banana", "orange"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}

	// Switch statement
	fmt.Println("\n=== Switch Examples ===")
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Regular work day")
	}

	// Switch without expression
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Excellent!")
	case score >= 80:
		fmt.Println("Good job!")
	default:
		fmt.Println("Keep practicing!")
	}

	// Break and continue
	fmt.Println("\n=== Break and Continue Examples ===")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Skip iteration when i is 2
		}
		if i == 4 {
			break // Exit loop when i is 4
		}
		fmt.Printf("Current number: %d\n", i)
	}
}
