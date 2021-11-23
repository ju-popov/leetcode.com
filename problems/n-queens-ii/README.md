# 52. N-Queens II

**Difficulty**: Hard

## Related Topics:

- [Backtracking](https://leetcode.com/tag/backtracking/)

## Problem:

The **n-queens** puzzle is the problem of placing `n` queens on an `n x n` chessboard such that no two queens attack each other.

Given an integer `n`, return *the number of distinct solutions to the **n-queens puzzle***.

**Example 1:**

```
Input: n = 4
Output: 2
Explanation: There are two distinct solutions to the 4-queens puzzle as shown.
```

**Example 2:**

```
Input: n = 1
Output: 1
```

**Constraints:**

- `1 <= n <= 9`

## Solution:

```go
func possibleXPositions(n int, y int, queens [][2]int) []int {
	result := make([]int, 0)

	for x := 0; x < n; x++ {
		threatened := false

		for _, queen := range queens {
			queenY := queen[0]
			queenX := queen[1]

			// Horizontal or Vertical
			if (queenX == x) || (queenY == y) {
				threatened = true

				break
			}

			// Diagonal \
			if (queenX - queenY) == (x - y) {
				threatened = true

				break
			}

			// Diagonal /
			if (queenX + queenY) == (x + y) {
				threatened = true

				break
			}
		}

		if !threatened {
			result = append(result, x)
		}
	}

	return result
}

func positionNQueens(n int, queens [][2]int, count int) int {
	result := 0

	if count == 0 {
		return 1
	}

	y := len(queens)

	for _, x := range possibleXPositions(n, y, queens) {
		newQueens := [][2]int{{y, x}}
		newQueens = append(newQueens, queens...)

		result += positionNQueens(n, newQueens, count-1)
	}

	return result
}

func totalNQueens(n int) int {
	queens := make([][2]int, 0)

	return positionNQueens(n, queens, n)
}
```

## Similar Questions:

- [N-Queens](https://leetcode.com/problems/n-queens/)
