# 62. Unique Paths

**Difficulty**: Medium

## Related Topics:

- [Math](https://leetcode.com/tag/math/)
- [Dynamic Programming](https://leetcode.com/tag/dynamic-programming/)
- [Combinatorics](https://leetcode.com/tag/combinatorics/)

## Problem:

A robot is located at the top-left corner of a `m x n` grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

**Example 1:**

```
Input: m = 3, n = 7
Output: 28
```

**Example 2:**

```
Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down
```

**Example 3:**

```
Input: m = 7, n = 3
Output: 28
```

**Example 4:**

```
Input: m = 3, n = 3
Output: 6
```

**Constraints:**

- `1 <= m, n <= 100`
- It's guaranteed that the answer will be less than or equal to `2 * 109`.

## Solution:

```go
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
```

## Similar Questions:

- [Unique Paths II](https://github.com/ju-popov/leetcode.com/tree/main/problems/unique-paths-ii/)
- [Minimum Path Sum](https://github.com/ju-popov/leetcode.com/tree/main/problems/minimum-path-sum/)
- [Dungeon Game](https://github.com/ju-popov/leetcode.com/tree/main/problems/dungeon-game/)
