// This file demonstrates error handling patterns in Go
// Go handles errors explicitly through return values rather than exceptions
// The error interface is a built-in type that represents error conditions

package main

import (
	"errors"
	"fmt"
)

// Custom error type
// Implementing the error interface by defining Error() string method
type DivisionError struct {
	dividend int
	divisor  int
	message  string
}

// Error method implementation
// This makes DivisionError implement the error interface
func (e *DivisionError) Error() string {
	return fmt.Sprintf("%s: %d / %d", e.message, e.dividend, e.divisor)
}

// Function that returns an error
// In Go, it's common to return (result, error)
// If error is nil, the operation was successful
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		// Return custom error for division by zero
		return 0, &DivisionError{
			dividend: dividend,
			divisor:  divisor,
			message:  "cannot divide by zero",
		}
	}
	return dividend / divisor, nil
}

// Function demonstrating multiple error cases
// Shows how to handle negative numbers as an error condition
func calculateSquareRoot(x float64) (float64, error) {
	if x < 0 {
		// errors.New creates a simple error with a message
		return 0, errors.New("cannot calculate square root of negative number")
	}
	return x * x, nil
}

func main() {
	// ==================== Basic Error Handling ====================
	// Basic pattern: check error return value
	result, err := divide(10, 0)
	if err != nil {
		// Error handling: print error message
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}

	// ==================== Multiple Error Cases ====================
	// Testing different scenarios that might cause errors
	numbers := []int{10, 0, 5, 2}
	for i := 0; i < len(numbers)-1; i++ {
		result, err := divide(numbers[i], numbers[i+1])
		if err != nil {
			// Type assertion to check if it's our custom error
			if divErr, ok := err.(*DivisionError); ok {
				fmt.Printf("Custom division error: %v\n", divErr)
			} else {
				fmt.Printf("Other error: %v\n", err)
			}
			continue // Skip to next iteration on error
		}
		fmt.Printf("%d / %d = %d\n", numbers[i], numbers[i+1], result)
	}

	// ==================== Using errors.New ====================
	// Testing with both valid and invalid inputs
	root, err := calculateSquareRoot(-4)
	if err != nil {
		fmt.Printf("Square root error: %v\n", err)
	} else {
		fmt.Printf("Square: %f\n", root)
	}

	// ==================== Panic and Recover ====================
	fmt.Println("\nPanic and Recover Example:")
	
	// defer is executed even if a panic occurs
	// recover() catches a panic and returns the value that was passed to panic()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
			// Here you could do cleanup or logging before the program continues
		}
	}()

	// This will cause a panic
	// Panic is used for unrecoverable errors
	// It stops normal execution and runs deferred functions
	panic("something went wrong!")

	// Code after a panic is not executed unless recovered
	fmt.Println("This won't be printed")
}
