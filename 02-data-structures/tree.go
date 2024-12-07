// This file implements a binary search tree (BST) data structure in Go
// A BST is a binary tree where for each node:
// - All nodes in left subtree have values less than the node
// - All nodes in right subtree have values greater than the node
//
// Time Complexity:
// - Insert: O(log n) average, O(n) worst case
// - Search: O(log n) average, O(n) worst case
// - Traversal: O(n)
// where n is the number of nodes
//
// Use Cases:
// - Implementing sets and maps
// - Database indexing
// - File system organization
// - Expression parsing
// - Priority queues

package main

import "fmt"

// TreeNode represents a node in a binary tree
// Each node contains:
// - Value: the data stored in the node
// - Left: pointer to left child (smaller values)
// - Right: pointer to right child (larger values)
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BinaryTree represents a binary search tree
// It contains a pointer to the root node
type BinaryTree struct {
	Root *TreeNode
}

// Insert adds a new value to the binary search tree
// Maintains BST property: left < node < right
// Time Complexity: O(log n) average, O(n) worst case
func (t *BinaryTree) Insert(value int) {
	// Case 1: Empty tree
	if t.Root == nil {
		t.Root = &TreeNode{Value: value}
		return
	}
	// Case 2: Non-empty tree
	t.insertRecursive(t.Root, value)
}

// insertRecursive is a helper function for Insert
// Recursively finds the correct position to insert the new value
func (t *BinaryTree) insertRecursive(node *TreeNode, value int) {
	// Insert to left subtree
	if value < node.Value {
		if node.Left == nil {
			// Found insertion point
			node.Left = &TreeNode{Value: value}
		} else {
			// Continue searching in left subtree
			t.insertRecursive(node.Left, value)
		}
	} else {
		// Insert to right subtree
		if node.Right == nil {
			// Found insertion point
			node.Right = &TreeNode{Value: value}
		} else {
			// Continue searching in right subtree
			t.insertRecursive(node.Right, value)
		}
	}
}

// Search looks for a value in the binary search tree
// Returns true if found, false otherwise
// Time Complexity: O(log n) average, O(n) worst case
func (t *BinaryTree) Search(value int) bool {
	return t.searchRecursive(t.Root, value)
}

// searchRecursive is a helper function for Search
// Recursively searches for the value in the tree
func (t *BinaryTree) searchRecursive(node *TreeNode, value int) bool {
	// Base case: reached nil node (value not found)
	if node == nil {
		return false
	}
	
	// Found the value
	if node.Value == value {
		return true
	}
	
	// Recursively search in appropriate subtree
	if value < node.Value {
		return t.searchRecursive(node.Left, value)
	}
	return t.searchRecursive(node.Right, value)
}

// InorderTraversal performs an inorder traversal of the tree
// Visits: left subtree -> node -> right subtree
// Result is sorted in ascending order for BST
// Time Complexity: O(n)
func (t *BinaryTree) InorderTraversal() []int {
	result := []int{}
	t.inorderRecursive(t.Root, &result)
	return result
}

// inorderRecursive is a helper function for InorderTraversal
func (t *BinaryTree) inorderRecursive(node *TreeNode, result *[]int) {
	if node != nil {
		t.inorderRecursive(node.Left, result)
		*result = append(*result, node.Value)
		t.inorderRecursive(node.Right, result)
	}
}

// PreorderTraversal performs a preorder traversal of the tree
// Visits: node -> left subtree -> right subtree
// Useful for creating a copy of the tree
// Time Complexity: O(n)
func (t *BinaryTree) PreorderTraversal() []int {
	result := []int{}
	t.preorderRecursive(t.Root, &result)
	return result
}

// preorderRecursive is a helper function for PreorderTraversal
func (t *BinaryTree) preorderRecursive(node *TreeNode, result *[]int) {
	if node != nil {
		*result = append(*result, node.Value)
		t.preorderRecursive(node.Left, result)
		t.preorderRecursive(node.Right, result)
	}
}

// PostorderTraversal performs a postorder traversal of the tree
// Visits: left subtree -> right subtree -> node
// Useful for deleting the tree or evaluating expressions
// Time Complexity: O(n)
func (t *BinaryTree) PostorderTraversal() []int {
	result := []int{}
	t.postorderRecursive(t.Root, &result)
	return result
}

// postorderRecursive is a helper function for PostorderTraversal
func (t *BinaryTree) postorderRecursive(node *TreeNode, result *[]int) {
	if node != nil {
		t.postorderRecursive(node.Left, result)
		t.postorderRecursive(node.Right, result)
		*result = append(*result, node.Value)
	}
}

func main() {
	// Create a binary search tree
	tree := &BinaryTree{}

	// Example 1: Inserting values
	fmt.Println("Example 1: Building the tree")
	fmt.Println("Inserting: 5, 3, 7, 1, 4, 6, 8")
	values := []int{5, 3, 7, 1, 4, 6, 8}
	// This creates the following tree:
	//       5
	//      / \
	//     3   7
	//    / \  / \
	//   1   4 6  8
	for _, value := range values {
		tree.Insert(value)
	}

	// Example 2: Different traversals
	fmt.Println("\nExample 2: Tree Traversals")
	fmt.Println("Inorder (sorted):", tree.InorderTraversal())
	fmt.Println("Preorder:", tree.PreorderTraversal())
	fmt.Println("Postorder:", tree.PostorderTraversal())

	// Example 3: Searching
	fmt.Println("\nExample 3: Searching for values")
	searchValues := []int{4, 9}
	for _, value := range searchValues {
		exists := tree.Search(value)
		fmt.Printf("Is %d in the tree? %v\n", value, exists)
	}
}
