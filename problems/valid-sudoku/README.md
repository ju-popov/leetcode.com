# 36. Valid Sudoku

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Hash Table](https://leetcode.com/tag/hash-table/)
- [Matrix](https://leetcode.com/tag/matrix/)

## Problem:

Determine if a `9 x 9` Sudoku board is valid. Only the filled cells need to be validated **according to the following rules**:

1. Each row must contain the digits `1-9` without repetition.
2. Each column must contain the digits `1-9` without repetition.
3. Each of the nine `3 x 3` sub-boxes of the grid must contain the digits `1-9` without repetition.

**Note:**

- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.

**Example 1:**

```
Input: board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
```

**Example 2:**

```
Input: board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
```

**Constraints:**

- `board.length == 9`
- `board[i].length == 9`
- `board[i][j]` is a digit `1-9` or `'.'`.

## Solution:

```go
func isValidArea(board [][]byte, yStart int, yEnd int, xStart int, xEnd int) bool {
	digits := make(map[byte]bool)

	for y := yStart; y <= yEnd; y++ {
		row := board[y]

		for x := xStart; x <= xEnd; x++ {
			char := row[x]

			if (char >= '0') && (char <= '9') {
				if digits[char] {
					return false
				}

				digits[char] = true
			}
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	// Rows
	for y := 0; y < len(board); y++ {
		if !isValidArea(board, y, y, 0, len(board[0])-1) {
			return false
		}
	}

	// Columns
	for x := 0; x < len(board[0]); x++ {
		if !isValidArea(board, 0, len(board)-1, x, x) {
			return false
		}
	}

	// Grids
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if !isValidArea(board, y*3, y*3+2, x*3, x*3+2) {
				return false
			}
		}
	}

	return true
}
```

## Similar Questions:

- [Sudoku Solver](https://github.com/ju-popov/leetcode.com/tree/main/problems/sudoku-solver/)
