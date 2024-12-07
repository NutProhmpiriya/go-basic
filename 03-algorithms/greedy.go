// This file demonstrates common greedy algorithms in Go
// Greedy algorithms make locally optimal choices at each step
// with the hope of finding a global optimum solution
//
// Common characteristics of Greedy Algorithms:
// 1. Makes the locally optimal choice at each step
// 2. Never reconsiders its choices
// 3. Works well for optimization problems

package main

import (
	"fmt"
	"sort"
)

// ActivitySelection solves the activity selection problem
// Given a set of activities with start and end times,
// find the maximum number of activities that can be performed
// Time Complexity: O(n log n) due to sorting
// Space Complexity: O(1)
type Activity struct {
	Start int
	End   int
}

func ActivitySelection(activities []Activity) []Activity {
	if len(activities) == 0 {
		return []Activity{}
	}

	// Sort activities by end time
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].End < activities[j].End
	})

	selected := []Activity{activities[0]}
	lastEnd := activities[0].End

	// Select activities that don't overlap
	for i := 1; i < len(activities); i++ {
		if activities[i].Start >= lastEnd {
			selected = append(selected, activities[i])
			lastEnd = activities[i].End
		}
	}

	return selected
}

// FractionalKnapsack solves the fractional knapsack problem
// Unlike 0/1 knapsack, we can take fractions of items
// Time Complexity: O(n log n) due to sorting
// Space Complexity: O(1)
type Item struct {
	Value  float64
	Weight float64
}

func FractionalKnapsack(items []Item, capacity float64) float64 {
	// Calculate value per unit weight and sort
	type ItemRatio struct {
		Item
		Ratio float64
	}

	ratios := make([]ItemRatio, len(items))
	for i, item := range items {
		ratios[i] = ItemRatio{item, item.Value / item.Weight}
	}

	// Sort by value/weight ratio in descending order
	sort.Slice(ratios, func(i, j int) bool {
		return ratios[i].Ratio > ratios[j].Ratio
	})

	totalValue := 0.0
	currentWeight := 0.0

	for _, item := range ratios {
		if currentWeight+item.Weight <= capacity {
			// Take whole item
			currentWeight += item.Weight
			totalValue += item.Value
		} else {
			// Take fraction of item
			remainingCapacity := capacity - currentWeight
			totalValue += item.Ratio * remainingCapacity
			break
		}
	}

	return totalValue
}

// HuffmanCoding implements Huffman coding for data compression
// Time Complexity: O(n log n)
// Space Complexity: O(n)
type HuffmanNode struct {
	Char       rune
	Freq       int
	Left, Right *HuffmanNode
}

type HuffmanHeap []*HuffmanNode

func (h HuffmanHeap) Len() int           { return len(h) }
func (h HuffmanHeap) Less(i, j int) bool { return h[i].Freq < h[j].Freq }
func (h HuffmanHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *HuffmanHeap) Push(x interface{}) {
	*h = append(*h, x.(*HuffmanNode))
}
func (h *HuffmanHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BuildHuffmanTree(text string) *HuffmanNode {
	// Count frequency of characters
	freq := make(map[rune]int)
	for _, c := range text {
		freq[c]++
	}

	// Create heap
	var heap HuffmanHeap
	for char, f := range freq {
		heap = append(heap, &HuffmanNode{Char: char, Freq: f})
	}
	sort.Sort(heap)

	// Build Huffman tree
	for len(heap) > 1 {
		left := heap[0]
		heap = heap[1:]
		right := heap[0]
		heap = heap[1:]

		internal := &HuffmanNode{
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}
		
		// Insert the new internal node
		i := 0
		for i < len(heap) && heap[i].Freq < internal.Freq {
			i++
		}
		heap = append(heap[:i], append([]*HuffmanNode{internal}, heap[i:]...)...)
	}

	return heap[0]
}

// DijkstraShortestPath implements Dijkstra's shortest path algorithm
// Time Complexity: O((V + E) log V) with binary heap
// Space Complexity: O(V)
type Edge struct {
	To     int
	Weight int
}

func DijkstraShortestPath(graph [][]Edge, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = int(1e9) // Initialize with infinity
	}
	dist[start] = 0

	// Priority queue entries: [vertex, distance]
	pq := make([][2]int, 0)
	pq = append(pq, [2]int{start, 0})

	for len(pq) > 0 {
		// Pop vertex with minimum distance
		curr := pq[0]
		pq = pq[1:]
		u := curr[0]
		d := curr[1]

		// Skip if we've found a better path
		if d > dist[u] {
			continue
		}

		// Check all neighbors
		for _, edge := range graph[u] {
			v := edge.To
			newDist := dist[u] + edge.Weight

			if newDist < dist[v] {
				dist[v] = newDist
				// Add to priority queue
				pq = append(pq, [2]int{v, newDist})
				// Keep queue sorted by distance
				sort.Slice(pq, func(i, j int) bool {
					return pq[i][1] < pq[j][1]
				})
			}
		}
	}

	return dist
}

func main() {
	// Example 1: Activity Selection
	activities := []Activity{
		{1, 4}, {3, 5}, {0, 6}, {5, 7},
		{3, 9}, {5, 9}, {6, 10}, {8, 11},
		{8, 12}, {2, 14}, {12, 16},
	}

	selected := ActivitySelection(activities)
	fmt.Println("Activity Selection Problem:")
	fmt.Printf("Selected activities: %+v\n\n", selected)

	// Example 2: Fractional Knapsack
	items := []Item{
		{60, 10},  // Value: 60, Weight: 10
		{100, 20}, // Value: 100, Weight: 20
		{120, 30}, // Value: 120, Weight: 30
	}
	capacity := 50.0

	maxValue := FractionalKnapsack(items, capacity)
	fmt.Println("Fractional Knapsack Problem:")
	fmt.Printf("Maximum value: %.2f\n\n", maxValue)

	// Example 3: Huffman Coding
	text := "this is an example for huffman encoding"
	huffmanTree := BuildHuffmanTree(text)
	fmt.Println("Huffman Coding:")
	fmt.Printf("Huffman tree root frequency: %d\n\n", huffmanTree.Freq)

	// Example 4: Dijkstra's Shortest Path
	graph := [][]Edge{
		{{1, 4}, {2, 1}},           // Edges from vertex 0
		{{3, 1}},                   // Edges from vertex 1
		{{1, 2}, {3, 5}},          // Edges from vertex 2
		{{4, 3}},                   // Edges from vertex 3
		{},                         // Edges from vertex 4
	}

	distances := DijkstraShortestPath(graph, 0)
	fmt.Println("Dijkstra's Shortest Path:")
	fmt.Printf("Shortest distances from vertex 0: %v\n", distances)
}
