# 51. N-Queens

**Difficulty**: Hard

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Backtracking](https://leetcode.com/tag/backtracking/)

## Problem:

The **n-queens** puzzle is the problem of placing `n` queens on an `n x n` chessboard such that no two queens attack each other.

Given an integer `n`, return *all distinct solutions to the **n-queens puzzle***. You may return the answer in **any order**.

Each solution contains a distinct board configuration of the n-queens' placement, where `'Q'` and `'.'` both indicate a queen and an empty space, respectively.

**Example 1:**

```
Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
```

**Example 2:**

```
Input: n = 1
Output: [["Q"]]
```

**Constraints:**

- `1 <= n <= 9`

## Solution:

```go
func solution(n int, queens [][2]int) []string {
	result := make([]string, 0)

	for y := 0; y < n; y++ {
		row := ""

		for x := 0; x < n; x++ {
			found := false

			for _, queen := range queens {
				if queen[0] == y && queen[1] == x {
					found = true

					break
				}
			}

			if found {
				row += "Q"
			} else {
				row += "."
			}
		}

		result = append(result, row)
	}

	return result
}

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

func positionNQueens(n int, queens [][2]int, count int, result [][]string) [][]string {
	if count == 0 {
		result = append(result, solution(n, queens))

		return result
	}

	y := len(queens)

	for _, x := range possibleXPositions(n, y, queens) {
		newQueens := [][2]int{{y, x}}
		newQueens = append(newQueens, queens...)

		result = positionNQueens(n, newQueens, count-1, result)
	}

	return result
}

func solveNQueens(n int) [][]string {
	queens := make([][2]int, 0)
	results := make([][]string, 0)

	return positionNQueens(n, queens, n, results)
}
```

## Similar Questions:

- [N-Queens II](https://leetcode.com/problems/n-queens-ii/)
- [Grid Illumination](https://leetcode.com/problems/grid-illumination/)
