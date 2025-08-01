package main

import (
	"slices"
	"strings"
	"time"
)

// SlowSort sorts a slice of integers using a very inefficient algorithm (bubble sort)
// TODO: Optimize this function to be more efficient
func SlowSort(data []int) []int {
	// Make a copy to avoid modifying the original
	result := make([]int, len(data))
	copy(result, data)

	// Bubble sort implementation
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}

// OptimizedSort is your optimized version of SlowSort
// It should produce identical results but perform better
func OptimizedSort(data []int) []int {
	// TODO: Implement a more efficient sorting algorithm
	// Hint: Consider using sort package or a more efficient algorithm
	sorted := make([]int, len(data))
	copy(sorted, data)
	slices.Sort(sorted)
	return sorted
}

// InefficientStringBuilder builds a string by repeatedly concatenating
// TODO: Optimize this function to be more efficient
func InefficientStringBuilder(parts []string, repeatCount int) string {
	result := ""

	for i := 0; i < repeatCount; i++ {
		for _, part := range parts {
			result += part
		}
	}

	return result
}

// OptimizedStringBuilder is your optimized version of InefficientStringBuilder
// It should produce identical results but perform better
func OptimizedStringBuilder(parts []string, repeatCount int) string {
	// TODO: Implement a more efficient string building method
	// Hint: Consider using strings.Builder or bytes.Buffer
	var builder strings.Builder

	for i := 0; i < repeatCount; i++ {
		for _, part := range parts {
			builder.WriteString(part)
		}
	}

	return builder.String()
}

// ExpensiveCalculation performs a computation with redundant work
// It computes the sum of all fibonacci numbers up to n
// TODO: Optimize this function to be more efficient
func ExpensiveCalculation(n int) int {

	// pkg: challenge16
	// cpu: Apple M3
	// BenchmarkExpensiveCalculation/Small-8         	 2908297	       402.9 ns/op	       0 B/op	       0 allocs/op
	// BenchmarkExpensiveCalculation/Medium-8        	   23930	     50013 ns/op	       0 B/op	       0 allocs/op
	// BenchmarkExpensiveCalculation/Large-8         	     194	   6210095 ns/op	       0 B/op	       0 allocs/op

	if n <= 0 {
		return 0
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += fibonacci(i)
	}

	return sum
}

// Helper function that computes the fibonacci number at position n
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Linear complexity with iteration
func OptimizedFibonacci(n int) int {

	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// OptimizedCalculation is your optimized version of ExpensiveCalculation
// It should produce identical results but perform better
func OptimizedCalculation(n int) int {
	// Hint: Consider memoization or avoiding redundant calculations
	if n <= 1 {
		return n
	}

	sum := 0

	// Iterate
	prev1, prev2 := 1, 0
	for i := 1; i <= n; i++ {
		prev2, prev1 = prev1, prev2+prev1
		sum += prev2
	}

	return sum
}

// HighAllocationSearch searches for all occurrences of a substring and creates a map with their positions
// TODO: Optimize this function to reduce allocations
func HighAllocationSearch(text, substr string) map[int]string {
	result := make(map[int]string)

	// Convert to lowercase for case-insensitive search
	lowerText := strings.ToLower(text)
	lowerSubstr := strings.ToLower(substr)

	for i := 0; i < len(lowerText); i++ {
		// Check if we can fit the substring starting at position i
		if i+len(lowerSubstr) <= len(lowerText) {
			// Extract the potential match
			potentialMatch := lowerText[i : i+len(lowerSubstr)]

			// Check if it matches
			if potentialMatch == lowerSubstr {
				// Store the original case version
				result[i] = text[i : i+len(substr)]
			}
		}
	}

	return result
}

// OptimizedSearch is your optimized version of HighAllocationSearch
// It should produce identical results but perform better with fewer allocations
func OptimizedSearch(text, substr string) map[int]string {
	result := make(map[int]string)

	// Convert to lowercase for case-insensitive search
	lowerText := strings.ToLower(text)
	lowerSubstr := strings.ToLower(substr)

	for i := 0; i < len(lowerText); i++ {
		// Check if we can fit the substring starting at position i
		if i+len(lowerSubstr) <= len(lowerText) {

			// Check if it matches
			if lowerText[i:i+len(lowerSubstr)] == lowerSubstr {
				// Store the original case version
				result[i] = text[i : i+len(substr)]
			}
		}
	}

	return result // Replace this with your optimized implementation
}

// A function to simulate CPU-intensive work for benchmarking
// You don't need to optimize this; it's just used for testing
func SimulateCPUWork(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
		// Just waste CPU cycles
		for i := 0; i < 1000000; i++ {
			_ = i
		}
	}
}
