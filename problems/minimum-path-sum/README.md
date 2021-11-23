# 64. Minimum Path Sum

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Dynamic Programming](https://leetcode.com/tag/dynamic-programming/)
- [Matrix](https://leetcode.com/tag/matrix/)

## Problem:

Given a `m x n` `grid` filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

**Note:** You can only move either down or right at any point in time.

**Example 1:**

```
Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
```

**Example 2:**

```
Input: grid = [[1,2,3],[4,5,6]]
Output: 12
```

**Constraints:**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 200`
- `0 <= grid[i][j] <= 100`

## Solution:

```go
func minInt(v1 int, v2 int) int {
	if v1 <= v2 {
		return v1
	}

	return v2
}

func minPathSum(grid [][]int) int {
	heigth := len(grid)
	width := len(grid[0])

	for y := 0; y < heigth; y++ {
		for x := 0; x < width; x++ {
			if x == 0 {
				if y == 0 {
					// first column and row
					continue
				} else {
					// first column
					top := grid[y-1][x]
					grid[y][x] += top
				}
			} else {
				if y == 0 {
					// first row
					left := grid[y][x-1]
					grid[y][x] += left
				} else {
					// all else
					top := grid[y-1][x]
					left := grid[y][x-1]
					grid[y][x] += minInt(top, left)
				}
			}
		}
	}

	return grid[heigth-1][width-1]
}
```

## Similar Questions:

- [Unique Paths](https://github.com/ju-popov/leetcode.com/tree/main/problems/unique-paths/)
- [Dungeon Game](https://github.com/ju-popov/leetcode.com/tree/main/problems/dungeon-game/)
- [Cherry Pickup](https://github.com/ju-popov/leetcode.com/tree/main/problems/cherry-pickup/)
- [Maximum Number of Points with Cost](https://github.com/ju-popov/leetcode.com/tree/main/problems/maximum-number-of-points-with-cost/)
