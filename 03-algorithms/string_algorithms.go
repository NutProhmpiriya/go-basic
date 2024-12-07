// This file implements common string algorithms in Go
// String algorithms are fundamental in text processing, pattern matching,
// and many other applications

package main

import (
	"fmt"
)

// KMPSearch implements the Knuth-Morris-Pratt string matching algorithm
// Time Complexity: O(n + m) where n is text length and m is pattern length
// Space Complexity: O(m) for the LPS array
func KMPSearch(text, pattern string) []int {
	// Compute LPS (Longest Proper Prefix which is also Suffix) array
	lps := computeLPSArray(pattern)
	matches := []int{}

	i, j := 0, 0 // i for text, j for pattern
	for i < len(text) {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == len(pattern) {
			matches = append(matches, i-j)
			j = lps[j-1]
		} else if i < len(text) && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return matches
}

func computeLPSArray(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0 // Length of previous longest prefix suffix
	i := 1

	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

// RabinKarp implements the Rabin-Karp string matching algorithm
// Time Complexity: O(n + m) average case, O(nm) worst case
// Space Complexity: O(1)
func RabinKarp(text, pattern string) []int {
	if len(pattern) > len(text) {
		return []int{}
	}

	// Constants for the rolling hash function
	const base = 256
	const prime = 101
	matches := []int{}

	// Calculate pattern hash
	patternHash := 0
	windowHash := 0
	for i := 0; i < len(pattern); i++ {
		patternHash = (patternHash*base + int(pattern[i])) % prime
		windowHash = (windowHash*base + int(text[i])) % prime
	}

	// Calculate largest polynomial term
	h := 1
	for i := 0; i < len(pattern)-1; i++ {
		h = (h * base) % prime
	}

	// Slide pattern over text
	for i := 0; i <= len(text)-len(pattern); i++ {
		if patternHash == windowHash {
			// Check character by character
			match := true
			for j := 0; j < len(pattern); j++ {
				if text[i+j] != pattern[j] {
					match = false
					break
				}
			}
			if match {
				matches = append(matches, i)
			}
		}

		// Calculate hash for next window
		if i < len(text)-len(pattern) {
			windowHash = (base*(windowHash-int(text[i])*h) + int(text[i+len(pattern)])) % prime
			if windowHash < 0 {
				windowHash += prime
			}
		}
	}

	return matches
}

// LevenshteinDistance calculates the minimum number of single-character edits
// required to change one string into another
// Time Complexity: O(mn)
// Space Complexity: O(mn)
func LevenshteinDistance(s1, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize first row and column
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j],   // deletion
					dp[i][j-1],   // insertion
					dp[i-1][j-1], // substitution
				)
			}
		}
	}

	return dp[m][n]
}

// LongestPalindromicSubstring finds the longest palindromic substring
// using dynamic programming
// Time Complexity: O(n²)
// Space Complexity: O(n²)
func LongestPalindromicSubstring(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	// Table[i][j] will be true if substring s[i..j] is palindrome
	table := make([][]bool, n)
	for i := range table {
		table[i] = make([]bool, n)
		// Single characters are palindromes
		table[i][i] = true
	}

	start := 0
	maxLength := 1

	// Check for substrings of length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			table[i][i+1] = true
			start = i
			maxLength = 2
		}
	}

	// Check for lengths greater than 2
	for length := 3; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1 // ending index
			if table[i+1][j-1] && s[i] == s[j] {
				table[i][j] = true
				if length > maxLength {
					start = i
					maxLength = length
				}
			}
		}
	}

	return s[start : start+maxLength]
}

// Helper function to find minimum of three integers
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func main() {
	// Example 1: KMP String Matching
	text := "AABAACAADAABAAABAA"
	pattern := "AABA"
	fmt.Println("KMP String Matching:")
	fmt.Printf("Text: %s\nPattern: %s\n", text, pattern)
	matches := KMPSearch(text, pattern)
	fmt.Printf("Pattern found at indices: %v\n\n", matches)

	// Example 2: Rabin-Karp String Matching
	text2 := "GEEKS FOR GEEKS"
	pattern2 := "GEEK"
	fmt.Println("Rabin-Karp String Matching:")
	fmt.Printf("Text: %s\nPattern: %s\n", text2, pattern2)
	matches2 := RabinKarp(text2, pattern2)
	fmt.Printf("Pattern found at indices: %v\n\n", matches2)

	// Example 3: Levenshtein Distance
	str1 := "kitten"
	str2 := "sitting"
	fmt.Println("Levenshtein Distance:")
	fmt.Printf("String 1: %s\nString 2: %s\n", str1, str2)
	distance := LevenshteinDistance(str1, str2)
	fmt.Printf("Edit distance: %d\n\n", distance)

	// Example 4: Longest Palindromic Substring
	text3 := "babad"
	fmt.Println("Longest Palindromic Substring:")
	fmt.Printf("Text: %s\n", text3)
	palindrome := LongestPalindromicSubstring(text3)
	fmt.Printf("Longest palindrome: %s\n", palindrome)
}
