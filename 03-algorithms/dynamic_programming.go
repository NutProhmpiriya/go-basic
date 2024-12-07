// This file demonstrates common dynamic programming problems and solutions in Go
// Dynamic Programming (DP) is a method for solving complex problems by breaking them
// down into simpler subproblems. It is applicable when:
// 1. The problem can be broken down into smaller subproblems
// 2. The solution to the problem can be constructed from solutions to its subproblems
// 3. The subproblems overlap

package main

import (
	"fmt"
)

// FibonacciRecursive calculates the nth Fibonacci number using recursion
// Time Complexity: O(2^n)
// Space Complexity: O(n) due to recursion stack
// This is inefficient due to repeated calculations
func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

// FibonacciDP calculates the nth Fibonacci number using dynamic programming
// Time Complexity: O(n)
// Space Complexity: O(n)
// Uses memoization to store previously calculated values
func FibonacciDP(n int) int {
	if n <= 1 {
		return n
	}

	// Create DP table
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	// Fill the table
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// LongestCommonSubsequence finds the length of the longest common subsequence
// between two strings using dynamic programming
// Time Complexity: O(m*n)
// Space Complexity: O(m*n)
// where m and n are the lengths of the input strings
func LongestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

// KnapsackProblem solves the 0/1 knapsack problem using dynamic programming
// Time Complexity: O(n*W)
// Space Complexity: O(n*W)
// where n is the number of items and W is the knapsack capacity
func KnapsackProblem(values []int, weights []int, capacity int) int {
	n := len(values)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Fill the DP table
	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if weights[i-1] <= w {
				dp[i][w] = max(values[i-1]+dp[i-1][w-weights[i-1]], dp[i-1][w])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][capacity]
}

// CoinChange finds the minimum number of coins needed to make a given amount
// Time Complexity: O(amount * len(coins))
// Space Complexity: O(amount)
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1 // Initialize with impossible value
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1 // No solution found
	}
	return dp[amount]
}

// Helper function for max value
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function for min value
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1: Fibonacci Numbers
	n := 10
	fmt.Printf("Fibonacci(%d) using recursion: %d\n", n, FibonacciRecursive(n))
	fmt.Printf("Fibonacci(%d) using DP: %d\n\n", n, FibonacciDP(n))

	// Example 2: Longest Common Subsequence
	text1 := "abcde"
	text2 := "ace"
	lcs := LongestCommonSubsequence(text1, text2)
	fmt.Printf("Length of Longest Common Subsequence between '%s' and '%s': %d\n\n", 
		text1, text2, lcs)

	// Example 3: 0/1 Knapsack Problem
	values := []int{60, 100, 120}    // Values of items
	weights := []int{10, 20, 30}     // Weights of items
	capacity := 50                    // Knapsack capacity
	maxValue := KnapsackProblem(values, weights, capacity)
	fmt.Printf("Maximum value in Knapsack: %d\n\n", maxValue)

	// Example 4: Coin Change Problem
	coins := []int{1, 2, 5}
	amount := 11
	minCoins := CoinChange(coins, amount)
	if minCoins != -1 {
		fmt.Printf("Minimum coins needed for amount %d: %d\n", amount, minCoins)
	} else {
		fmt.Printf("Cannot make amount %d with given coins\n", amount)
	}
}
