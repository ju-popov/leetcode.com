# 73. Set Matrix Zeroes

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Hash Table](https://leetcode.com/tag/hash-table/)
- [Matrix](https://leetcode.com/tag/matrix/)

## Problem:

Given an `m x n` integer matrix `matrix`, if an element is `0`, set its entire row and column to `0`'s, and return *the matrix*.

You must do it [in place](https://en.wikipedia.org/wiki/In-place_algorithm).

**Example 1:**

```
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]
```

**Example 2:**

```
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
```

**Constraints:**

- `m == matrix.length`
- `n == matrix[0].length`
- `1 <= m, n <= 200`
- `-231 <= matrix[i][j] <= 231 - 1`

**Follow up:**

- A straightforward solution using `O(mn)` space is probably a bad idea.
- A simple improvement uses `O(m + n)` space, but still not the best solution.
- Could you devise a constant space solution?

## Solution:

```go
func setZeroes(matrix [][]int) {
	zeroInColumn0 := false
	zeroInRow0 := false

	// matrix[0][0..n] - Mark if first column contains zeroes
	y := 0

	for x := 0; x < len(matrix[0]); x++ {
		if matrix[y][x] == 0 {
			zeroInColumn0 = true

			break
		}
	}

	// matrix[0..m][0] - Mark if first row contains zeroes
	x := 0

	for y := 0; y < len(matrix); y++ {
		if matrix[y][x] == 0 {
			zeroInRow0 = true

			break
		}
	}

	// matrix[1..m][1..n] - Mark rows and columns that containt zeroes
	for y := 1; y < len(matrix); y++ {
		for x := 1; x < len(matrix[y]); x++ {
			if matrix[y][x] != 0 {
				continue
			}

			matrix[0][x] = 0
			matrix[y][0] = 0
		}
	}

	// matrix[1..m][1..n] - Fill with zeros
	for y := 1; y < len(matrix); y++ {
		for x := 1; x < len(matrix[y]); x++ {
			if (matrix[0][x] == 0) || (matrix[y][0] == 0) {
				matrix[y][x] = 0
			}
		}
	}

	// matrix[0][0..n] - Fill first column with zeroes
	if zeroInColumn0 {
		y := 0

		for x := 0; x < len(matrix[y]); x++ {
			matrix[y][x] = 0
		}
	}

	// matrix[0..m][0] - Fill first row with zeroes
	if zeroInRow0 {
		x := 0

		for y := 0; y < len(matrix); y++ {
			matrix[y][x] = 0
		}
	}
}
```

## Similar Questions:
- [Game of Life](https://github.com/ju-popov/leetcode.com/tree/main/problems/game-of-life/)
