# 63. Unique Paths II

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Dynamic Programming](https://leetcode.com/tag/dynamic-programming/)
- [Matrix](https://leetcode.com/tag/matrix/)

## Problem:

A robot is located at the top-left corner of a `m x n` grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

Now consider if some obstacles are added to the grids. How many unique paths would there be?

An obstacle and space is marked as `1` and `0` respectively in the grid.

**Example 1:**

```
Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
Output: 2
Explanation: There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
```

**Example 2:**

```
Input: obstacleGrid = [[0,1],[0,0]]
Output: 1
```

**Constraints:**

- `m == obstacleGrid.length`
- `n == obstacleGrid[i].length`
- `1 <= m, n <= 100`
- `obstacleGrid[i][j]` is `0` or `1`.

## Solution:

```go
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	height := len(obstacleGrid)
	width := len(obstacleGrid[0])

	if obstacleGrid[height-1][width-1] == 1 {
		return 0
	}

	for y := height - 1; y >= 0; y-- {
		for x := width - 1; x >= 0; x-- {
			if obstacleGrid[y][x] == 1 {
				obstacleGrid[y][x] = 0

				continue
			}

			if (y == height-1) && (x == width-1) {
				obstacleGrid[y][x] = 1

				continue
			}

			right := 0
			if x < width-1 {
				right = obstacleGrid[y][x+1]
			}

			bottom := 0
			if y < height-1 {
				bottom = obstacleGrid[y+1][x]
			}

			obstacleGrid[y][x] = right + bottom
		}
	}

	return obstacleGrid[0][0]
}
```

## Similar Questions:

- [Unique Paths](https://github.com/ju-popov/leetcode.com/tree/main/problems/unique-paths/)
- [Unique Paths III](https://github.com/ju-popov/leetcode.com/tree/main/problems/unique-paths-iii/)
