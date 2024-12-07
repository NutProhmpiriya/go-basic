// This file demonstrates the usage of arrays and slices in Go
// Arrays are fixed-size sequences of elements of the same type
// Slices are dynamic, flexible views into arrays

package main

import "fmt"

func main() {
	// ==================== Arrays ====================
	fmt.Println("Arrays Examples:")

	// Method 1: Declare a fixed-size array with explicit size
	// Arrays in Go are fixed-size and zero-based indexed
	// This creates an array of size 5 with values [1,2,3,4,5]
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v\n", numbers)

	// Method 2: Array with implicit size using [...]
	// Go compiler will count the elements and set the size
	fruits := [...]string{"apple", "banana", "orange"}
	fmt.Printf("Fruits: %v\n", fruits)

	// ==================== Slices ====================
	fmt.Println("\nSlices Examples:")

	// Method 1: Create slice from array
	// Syntax: array[startIndex:endIndex]
	// This creates a slice that references elements from index 1 to 3 (exclusive)
	numbersSlice := numbers[1:4]
	fmt.Printf("Slice from array: %v\n", numbersSlice)

	// Method 2: Create slice with make function
	// make([]T, length, capacity)
	// - T: type of elements
	// - length: initial length of slice
	// - capacity: maximum size the slice can grow to
	dynamicSlice := make([]int, 3, 5)
	fmt.Printf("Slice with make: %v (len=%d, cap=%d)\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	// Append elements to slice
	// When capacity is exceeded, Go automatically allocates a new array
	// with larger capacity and copies elements over
	dynamicSlice = append(dynamicSlice, 1, 2, 3)
	fmt.Printf("After append: %v (len=%d, cap=%d)\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	// Combining slices using append with ... operator
	// The ... operator unpacks the second slice's elements
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	combined := append(slice1, slice2...)
	fmt.Printf("Combined slices: %v\n", combined)
}
