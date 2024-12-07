// This file implements an undirected graph data structure in Go
// A graph is a collection of vertices (nodes) connected by edges
// This implementation uses an adjacency list representation
//
// Time Complexity:
// - Add Vertex: O(1)
// - Add Edge: O(1)
// - Get Neighbors: O(1)
// - BFS: O(V + E) where V is number of vertices and E is number of edges
// - DFS: O(V + E)
//
// Use Cases:
// - Social networks (friends connections)
// - Road networks (city connections)
// - Computer networks (network topology)
// - Recommendation systems
// - Game development (map navigation)

package main

import "fmt"

// Graph represents an adjacency list graph
// vertices is a map where:
// - key: vertex ID
// - value: slice of adjacent vertex IDs
type Graph struct {
	vertices map[int][]int
}

// NewGraph creates a new graph
// Returns a pointer to an initialized Graph structure
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

// AddVertex adds a vertex to the graph
// If vertex already exists, it won't be added again
// Time Complexity: O(1)
func (g *Graph) AddVertex(vertex int) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = []int{}
	}
}

// AddEdge adds an edge between two vertices
// Since this is an undirected graph, we add the edge in both directions
// Time Complexity: O(1)
func (g *Graph) AddEdge(vertex1, vertex2 int) {
	// Add vertex2 to vertex1's adjacency list
	g.vertices[vertex1] = append(g.vertices[vertex1], vertex2)
	// Add vertex1 to vertex2's adjacency list
	g.vertices[vertex2] = append(g.vertices[vertex2], vertex1)
}

// GetNeighbors returns all neighbors of a vertex
// Time Complexity: O(1)
func (g *Graph) GetNeighbors(vertex int) []int {
	return g.vertices[vertex]
}

// BFS performs breadth-first search starting from a vertex
// BFS explores all vertices at current depth before moving to next depth
// Time Complexity: O(V + E)
func (g *Graph) BFS(start int) []int {
	// Track visited vertices
	visited := make(map[int]bool)
	// Queue for BFS traversal
	queue := []int{start}
	// Mark start vertex as visited
	visited[start] = true
	// Store traversal result
	result := []int{}

	// Continue until queue is empty
	for len(queue) > 0 {
		// Dequeue vertex
		vertex := queue[0]
		queue = queue[1:]
		// Add to result
		result = append(result, vertex)

		// Visit all unvisited neighbors
		for _, neighbor := range g.vertices[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

// DFS performs depth-first search starting from a vertex
// DFS explores as far as possible along each branch before backtracking
// Time Complexity: O(V + E)
func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	result := []int{}
	g.dfsUtil(start, visited, &result)
	return result
}

// dfsUtil is a helper function for DFS
// Uses recursion to traverse the graph
func (g *Graph) dfsUtil(vertex int, visited map[int]bool, result *[]int) {
	// Mark current vertex as visited
	visited[vertex] = true
	// Add to result
	*result = append(*result, vertex)

	// Visit all unvisited neighbors recursively
	for _, neighbor := range g.vertices[vertex] {
		if !visited[neighbor] {
			g.dfsUtil(neighbor, visited, result)
		}
	}
}

func main() {
	// Create a new graph
	graph := NewGraph()

	// Example 1: Adding vertices
	fmt.Println("Example 1: Adding vertices 0-5")
	for i := 0; i < 6; i++ {
		graph.AddVertex(i)
	}

	// Example 2: Adding edges to create this graph:
	// 0 -- 1 -- 2
	// |    |    |
	// 3 -- 4 -- 5
	fmt.Println("\nExample 2: Adding edges")
	edges := [][2]int{
		{0, 1}, {1, 2}, // Top row
		{0, 3}, {1, 4}, {2, 5}, // Vertical connections
		{3, 4}, {4, 5}, // Bottom row
	}
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
		fmt.Printf("Added edge: %d -- %d\n", edge[0], edge[1])
	}

	// Example 3: Print adjacency list
	fmt.Println("\nExample 3: Graph Adjacency List:")
	for vertex, neighbors := range graph.vertices {
		fmt.Printf("Vertex %d: %v\n", vertex, neighbors)
	}

	// Example 4: BFS traversal
	fmt.Println("\nExample 4: BFS starting from vertex 0:")
	bfsResult := graph.BFS(0)
	fmt.Printf("BFS path: %v\n", bfsResult)

	// Example 5: DFS traversal
	fmt.Println("\nExample 5: DFS starting from vertex 0:")
	dfsResult := graph.DFS(0)
	fmt.Printf("DFS path: %v\n", dfsResult)

	// Example 6: Finding neighbors
	vertex := 1
	fmt.Printf("\nExample 6: Neighbors of vertex %d:\n", vertex)
	neighbors := graph.GetNeighbors(vertex)
	fmt.Printf("Neighbors: %v\n", neighbors)
}
