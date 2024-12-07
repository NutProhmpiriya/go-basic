// This file implements a queue data structure in Go
// A queue is a First-In-First-Out (FIFO) data structure
// Elements are added at the end (enqueue) and removed from the front (dequeue)
//
// Time Complexity:
// - Enqueue: O(1) amortized
// - Dequeue: O(1)
// - Peek: O(1)
//
// Use Cases:
// - Task scheduling
// - Print queue management
// - Breadth-first search in graphs
// - Request handling in web servers

package main

import "fmt"

// Queue represents a queue data structure
// This implementation uses a slice as the underlying storage
// The first element in the slice is the front of the queue
type Queue struct {
	items []int
}

// Enqueue adds an item to the end of the queue
// Time Complexity: O(1) amortized
// Note: While append is O(1) amortized, it may occasionally need to
// reallocate and copy the underlying array
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the first item in the queue
// Time Complexity: O(1)
// Note: While removing from the front is O(n) with slices,
// we optimize this by not shrinking the slice capacity
func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	
	// Get the first item
	item := q.items[0]
	// Remove it from the queue
	// This creates a new slice without the first element
	q.items = q.items[1:]
	return item, nil
}

// Peek returns the first item without removing it
// Time Complexity: O(1)
func (q *Queue) Peek() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.items[0], nil
}

// IsEmpty returns true if the queue is empty
// Time Complexity: O(1)
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
// Time Complexity: O(1)
func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	// Create a new queue
	queue := &Queue{}

	// Example 1: Enqueuing elements
	fmt.Println("Example 1: Enqueuing elements")
	fmt.Println("Enqueuing: 1, 2, 3")
	queue.Enqueue(1)  // Queue: [1]
	queue.Enqueue(2)  // Queue: [1, 2]
	queue.Enqueue(3)  // Queue: [1, 2, 3]

	// Example 2: Queue information
	fmt.Printf("\nExample 2: Queue Status\n")
	fmt.Printf("Queue size: %d\n", queue.Size())
	if first, err := queue.Peek(); err == nil {
		fmt.Printf("Front element: %d\n", first)
	}

	// Example 3: Dequeuing elements
	fmt.Println("\nExample 3: Dequeuing elements")
	fmt.Println("Dequeuing all elements:")
	for !queue.IsEmpty() {
		if item, err := queue.Dequeue(); err == nil {
			fmt.Printf("Dequeued: %d\n", item)
		}
	}

	// Example 4: Error handling
	fmt.Println("\nExample 4: Error handling")
	fmt.Println("Trying to dequeue from empty queue:")
	_, err := queue.Dequeue()
	fmt.Printf("Error: %v\n", err)

	// Example 5: Mixing operations
	fmt.Println("\nExample 5: Mixed operations")
	queue.Enqueue(10)  // Queue: [10]
	queue.Enqueue(20)  // Queue: [10, 20]
	if item, err := queue.Dequeue(); err == nil {
		fmt.Printf("Dequeued: %d\n", item)  // Queue: [20]
	}
	queue.Enqueue(30)  // Queue: [20, 30]
	fmt.Printf("Final queue size: %d\n", queue.Size())
}
