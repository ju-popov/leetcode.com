package main

import "fmt"

/*

62. Unique Paths

https://leetcode.com/problems/unique-paths/

Approach 1: Dynamic Programming

Complexity Analysis

Time complexity: O(N×M).

Space complexity: O(N×M).

*/

type memoryKey struct {
	m int
	n int
}

func uniquePathsHelper(m int, n int, memory map[memoryKey]int) int {
	if (m < 1) || (n < 1) {
		return 0
	}

	if (m == 1) && (n == 1) {
		return 1
	}

	if value, ok := memory[memoryKey{m: m, n: n}]; ok {
		return value
	}

	right := uniquePathsHelper(m, n-1, memory)
	down := uniquePathsHelper(m-1, n, memory)

	memory[memoryKey{m: m, n: n}] = right + down

	return memory[memoryKey{m: m, n: n}]
}

func uniquePaths(m int, n int) int {
	memory := make(map[memoryKey]int)

	return uniquePathsHelper(m, n, memory)
}

func main() {
	fmt.Println(uniquePaths(3, 7)) // 28
	fmt.Println(uniquePaths(3, 2)) // 3
	fmt.Println(uniquePaths(7, 3)) // 28
	fmt.Println(uniquePaths(3, 3)) // 6
}
