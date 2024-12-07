// This file implements common sorting algorithms in Go
// Each algorithm has different characteristics making them suitable for different scenarios
//
// Common Sorting Algorithms:
// 1. Bubble Sort: Simple but inefficient
// 2. Quick Sort: Efficient and widely used
// 3. Merge Sort: Stable and predictable performance
// 4. Insertion Sort: Efficient for small or nearly sorted data

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// BubbleSort implements the bubble sort algorithm
// Time Complexity: O(n²) for all cases
// Space Complexity: O(1)
// Stable: Yes
// Best for: Small datasets or nearly sorted arrays
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Flag to optimize if array is already sorted
		swapped := false
		
		// Compare adjacent elements
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap if they are in wrong order
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		
		// If no swapping occurred, array is sorted
		if !swapped {
			break
		}
	}
}

// QuickSort implements the quicksort algorithm
// Time Complexity: O(n log n) average, O(n²) worst case
// Space Complexity: O(log n)
// Stable: No
// Best for: General purpose sorting, works well with random data
func QuickSort(arr []int) {
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		// Find the partition index
		pi := partition(arr, low, high)
		
		// Recursively sort the left part
		quickSortHelper(arr, low, pi-1)
		// Recursively sort the right part
		quickSortHelper(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// Choose the rightmost element as pivot
	pivot := arr[high]
	i := low - 1 // Index of smaller element
	
	// Move elements smaller than pivot to the left
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	
	// Place pivot in its correct position
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// MergeSort implements the merge sort algorithm
// Time Complexity: O(n log n) for all cases
// Space Complexity: O(n)
// Stable: Yes
// Best for: When stable sorting is needed and memory is not a constraint
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Split array into two halves
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Compare and merge elements
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Add remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// InsertionSort implements the insertion sort algorithm
// Time Complexity: O(n²) worst/average case, O(n) best case
// Space Complexity: O(1)
// Stable: Yes
// Best for: Small datasets or nearly sorted arrays
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Move elements greater than key ahead
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Helper function to generate random array
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	return arr
}

// Helper function to check if array is sorted
func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func main() {
	// Example 1: Bubble Sort
	fmt.Println("Example 1: Bubble Sort")
	arr1 := generateRandomArray(10)
	fmt.Printf("Original array: %v\n", arr1)
	BubbleSort(arr1)
	fmt.Printf("Sorted array: %v\n", arr1)
	fmt.Printf("Is sorted? %v\n\n", isSorted(arr1))

	// Example 2: Quick Sort
	fmt.Println("Example 2: Quick Sort")
	arr2 := generateRandomArray(10)
	fmt.Printf("Original array: %v\n", arr2)
	QuickSort(arr2)
	fmt.Printf("Sorted array: %v\n", arr2)
	fmt.Printf("Is sorted? %v\n\n", isSorted(arr2))

	// Example 3: Merge Sort
	fmt.Println("Example 3: Merge Sort")
	arr3 := generateRandomArray(10)
	fmt.Printf("Original array: %v\n", arr3)
	arr3 = MergeSort(arr3)
	fmt.Printf("Sorted array: %v\n", arr3)
	fmt.Printf("Is sorted? %v\n\n", isSorted(arr3))

	// Example 4: Insertion Sort
	fmt.Println("Example 4: Insertion Sort")
	arr4 := generateRandomArray(10)
	fmt.Printf("Original array: %v\n", arr4)
	InsertionSort(arr4)
	fmt.Printf("Sorted array: %v\n", arr4)
	fmt.Printf("Is sorted? %v\n", isSorted(arr4))
}
