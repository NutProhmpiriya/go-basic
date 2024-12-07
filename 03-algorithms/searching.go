// This file implements common searching algorithms in Go
// Different searching algorithms are suitable for different scenarios
// based on the data structure and whether the data is sorted
//
// Common Searching Algorithms:
// 1. Linear Search: Simple search through unsorted data
// 2. Binary Search: Efficient search in sorted data
// 3. Jump Search: Balance between linear and binary search
// 4. Interpolation Search: Improved binary search for uniformly distributed data

package main

import (
	"fmt"
	"math"
)

// LinearSearch implements the linear search algorithm
// Time Complexity: O(n)
// Space Complexity: O(1)
// Advantages:
// - Simple to implement
// - Works on unsorted arrays
// - Good for small arrays
// Disadvantages:
// - Slow for large arrays
func LinearSearch(arr []int, target int) int {
	// Iterate through each element
	for i, value := range arr {
		if value == target {
			return i // Found the target
		}
	}
	return -1 // Target not found
}

// BinarySearch implements the binary search algorithm
// Time Complexity: O(log n)
// Space Complexity: O(1)
// Requirements:
// - Array must be sorted
// Advantages:
// - Very efficient for large sorted arrays
// - Logarithmic time complexity
func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		// Calculate middle point
		mid := left + (right-left)/2

		// Check if target is present at mid
		if arr[mid] == target {
			return mid
		}

		// If target greater, ignore left half
		if arr[mid] < target {
			left = mid + 1
		} else {
			// If target smaller, ignore right half
			right = mid - 1
		}
	}

	return -1 // Target not found
}

// JumpSearch implements the jump search algorithm
// Time Complexity: O(√n)
// Space Complexity: O(1)
// Requirements:
// - Array must be sorted
// Advantages:
// - Better than linear search
// - Doesn't require as many comparisons as binary search
// Best for:
// - Searching in arrays where jumping back is expensive
func JumpSearch(arr []int, target int) int {
	n := len(arr)
	// Finding optimal jump size
	step := int(math.Sqrt(float64(n)))

	// Finding the block where element is present (if exists)
	prev := 0
	for minStep := min(step, n)-1; arr[minStep] < target; minStep = min(step, n)-1 {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1
		}
	}

	// Doing linear search for target in block beginning with prev
	for arr[prev] < target {
		prev++
		if prev == min(step, n) {
			return -1
		}
	}

	// If element is found
	if arr[prev] == target {
		return prev
	}

	return -1
}

// InterpolationSearch implements the interpolation search algorithm
// Time Complexity: O(log log n) average case, O(n) worst case
// Space Complexity: O(1)
// Requirements:
// - Array must be sorted
// - Values should be uniformly distributed
// Advantages:
// - Much better than binary search for uniformly distributed data
// - Can outperform binary search in practice
func InterpolationSearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high && target >= arr[low] && target <= arr[high] {
		if low == high {
			if arr[low] == target {
				return low
			}
			return -1
		}

		// Probing the position with keeping uniform distribution in mind
		pos := low + ((high-low)*(target-arr[low]))/(arr[high]-arr[low])

		// Target found
		if arr[pos] == target {
			return pos
		}

		// If target is larger, target is in right sub array
		if arr[pos] < target {
			low = pos + 1
		} else {
			// If target is smaller, target is in left sub array
			high = pos - 1
		}
	}
	return -1
}

// Helper function for min value
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test array (sorted for binary, jump, and interpolation search)
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 13
	fmt.Printf("Searching for %d in array: %v\n\n", target, arr)

	// Example 1: Linear Search
	fmt.Println("Example 1: Linear Search")
	result := LinearSearch(arr, target)
	if result != -1 {
		fmt.Printf("Element found at index: %d\n", result)
	} else {
		fmt.Println("Element not found")
	}

	// Example 2: Binary Search
	fmt.Println("\nExample 2: Binary Search")
	result = BinarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Element found at index: %d\n", result)
	} else {
		fmt.Println("Element not found")
	}

	// Example 3: Jump Search
	fmt.Println("\nExample 3: Jump Search")
	result = JumpSearch(arr, target)
	if result != -1 {
		fmt.Printf("Element found at index: %d\n", result)
	} else {
		fmt.Println("Element not found")
	}

	// Example 4: Interpolation Search
	fmt.Println("\nExample 4: Interpolation Search")
	result = InterpolationSearch(arr, target)
	if result != -1 {
		fmt.Printf("Element found at index: %d\n", result)
	} else {
		fmt.Println("Element not found")
	}

	// Example 5: Searching for non-existent element
	target = 10
	fmt.Printf("\nSearching for non-existent element %d:\n", target)
	fmt.Printf("Linear Search: %d\n", LinearSearch(arr, target))
	fmt.Printf("Binary Search: %d\n", BinarySearch(arr, target))
	fmt.Printf("Jump Search: %d\n", JumpSearch(arr, target))
	fmt.Printf("Interpolation Search: %d\n", InterpolationSearch(arr, target))
}
