// This file implements a singly linked list data structure in Go
// A linked list is a linear data structure where elements are stored in nodes,
// and each node points to the next node in the sequence
//
// Time Complexity:
// - Insert at end: O(n)
// - Delete: O(n)
// - Search: O(n)
// where n is the number of nodes in the list
//
// Use Cases:
// - When frequent insertions and deletions are needed
// - When the size of the data is unknown in advance
// - When memory needs to be allocated dynamically

package main

import "fmt"

// Node represents a node in the linked list
// Each node contains:
// - data: the value stored in the node
// - next: pointer to the next node in the sequence
type Node struct {
	data int
	next *Node
}

// LinkedList represents the linked list
// It contains a pointer to the head (first node) of the list
// Head being nil indicates an empty list
type LinkedList struct {
	head *Node
}

// Insert adds a new node at the end of the list
// Time Complexity: O(n) where n is the number of nodes
// This could be optimized to O(1) by maintaining a tail pointer
func (l *LinkedList) Insert(data int) {
	newNode := &Node{data: data}
	
	// Case 1: Empty list
	// Make the new node the head
	if l.head == nil {
		l.head = newNode
		return
	}

	// Case 2: Non-empty list
	// Traverse to the last node and add the new node
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Delete removes the first occurrence of data in the list
// Returns true if the element was found and deleted, false otherwise
// Time Complexity: O(n) where n is the number of nodes
func (l *LinkedList) Delete(data int) bool {
	// Case 1: Empty list
	if l.head == nil {
		return false
	}

	// Case 2: Data is in head node
	if l.head.data == data {
		l.head = l.head.next
		return true
	}

	// Case 3: Data is in other nodes
	// Traverse the list looking for the node to delete
	current := l.head
	for current.next != nil {
		if current.next.data == data {
			// Found the node to delete
			// Update the next pointer to skip over it
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

// Print displays all elements in the list
// Format: value1 -> value2 -> value3 -> nil
func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	// Create a new linked list
	list := &LinkedList{}

	// Example 1: Insert elements
	fmt.Println("Example 1: Inserting elements")
	list.Insert(1)  // List: 1 -> nil
	list.Insert(2)  // List: 1 -> 2 -> nil
	list.Insert(3)  // List: 1 -> 2 -> 3 -> nil
	list.Insert(4)  // List: 1 -> 2 -> 3 -> 4 -> nil
	fmt.Print("Original List: ")
	list.Print()

	// Example 2: Delete an element
	fmt.Println("\nExample 2: Deleting element 2")
	list.Delete(2)  // List: 1 -> 3 -> 4 -> nil
	fmt.Print("After deleting 2: ")
	list.Print()

	// Example 3: Insert after deletion
	fmt.Println("\nExample 3: Inserting element 5")
	list.Insert(5)  // List: 1 -> 3 -> 4 -> 5 -> nil
	fmt.Print("After inserting 5: ")
	list.Print()
}
