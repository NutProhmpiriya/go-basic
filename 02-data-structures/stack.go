// This file implements a stack data structure in Go
// A stack is a Last-In-First-Out (LIFO) data structure
// Elements are added (pushed) and removed (popped) from the same end
//
// Time Complexity:
// - Push: O(1) amortized
// - Pop: O(1)
// - Peek: O(1)
//
// Use Cases:
// - Function call management (call stack)
// - Expression evaluation
// - Undo/Redo operations
// - Depth-first search implementation
// - Parentheses matching

package main

import "fmt"

// Stack represents a stack data structure
// This implementation uses a slice as the underlying storage
// The last element in the slice is the top of the stack
type Stack struct {
	items []int
}

// Push adds an item to the top of the stack
// Time Complexity: O(1) amortized
// Note: While append is O(1) amortized, it may occasionally need to
// reallocate and copy the underlying array
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
// Time Complexity: O(1)
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	// Get the index of the last item
	index := len(s.items) - 1
	// Get the last item
	item := s.items[index]
	// Remove it from the stack
	s.items = s.items[:index]
	return item, nil
}

// Peek returns the top item without removing it
// Time Complexity: O(1)
func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack is empty
// Time Complexity: O(1)
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
// Time Complexity: O(1)
func (s *Stack) Size() int {
	return len(s.items)
}

// Example application: Check if brackets are balanced
// This is a common use case for stacks
// Time Complexity: O(n) where n is the length of the input string
func isValidBrackets(s string) bool {
	stack := &Stack{}
	
	// Iterate through each character in the string
	for _, ch := range s {
		if ch == '(' {
			// Push for opening bracket
			stack.Push(1)
		} else if ch == ')' {
			// For closing bracket, check if there's a matching opening bracket
			if stack.IsEmpty() {
				// More closing brackets than opening brackets
				return false
			}
			stack.Pop()
		}
	}
	
	// Stack should be empty if brackets are balanced
	return stack.IsEmpty()
}

func main() {
	// Create a new stack
	stack := &Stack{}

	// Example 1: Pushing elements
	fmt.Println("Example 1: Pushing elements")
	fmt.Println("Pushing: 1, 2, 3")
	stack.Push(1)  // Stack: [1]
	stack.Push(2)  // Stack: [1, 2]
	stack.Push(3)  // Stack: [1, 2, 3]

	// Example 2: Stack information
	fmt.Printf("\nExample 2: Stack Status\n")
	fmt.Printf("Stack size: %d\n", stack.Size())
	if top, err := stack.Peek(); err == nil {
		fmt.Printf("Top element: %d\n", top)
	}

	// Example 3: Popping elements
	fmt.Println("\nExample 3: Popping elements")
	fmt.Println("Popping all elements:")
	for !stack.IsEmpty() {
		if item, err := stack.Pop(); err == nil {
			fmt.Printf("Popped: %d\n", item)
		}
	}

	// Example 4: Error handling
	fmt.Println("\nExample 4: Error handling")
	fmt.Println("Trying to pop from empty stack:")
	_, err := stack.Pop()
	fmt.Printf("Error: %v\n", err)

	// Example 5: Bracket matching application
	fmt.Println("\nExample 5: Bracket Matching Example")
	testCases := []string{"((()))", "(()())", "(()", ")("}
	for _, test := range testCases {
		fmt.Printf("Is '%s' valid? %v\n", test, isValidBrackets(test))
	}
}
