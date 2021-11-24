# 37. Sudoku Solver

**Difficulty**: Hard

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Backtracking](https://leetcode.com/tag/backtracking/)
- [Matrix](https://leetcode.com/tag/matrix/)

## Problem:

Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy **all of the following rules**:

1. Each of the digits `1-9` must occur exactly once in each row.
2. Each of the digits `1-9` must occur exactly once in each column.
3. Each of the digits `1-9` must occur exactly once in each of the 9 `3x3` sub-boxes of the grid.

The `'.'` character indicates empty cells.

**Example 1:**

```
Input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
Explanation: The input board is shown above and the only valid solution is shown below:
```

**Constraints:**

- `board.length == 9`
- `board[i].length == 9`
- `board[i][j]` is a digit or `'.'`.
- It is **guaranteed** that the input board has only one solution.

## Solution:

```go
func nextUnknownSquare(board [][]byte) (int, int) {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] == '.' {
				return y, x
			}
		}
	}

	return -1, -1
}

func possibleDigits(board [][]byte, y int, x int) []int {
	result := make([]int, 0)

	digits := make([]bool, 9)

	// Horizontal
	for xPos := 0; xPos < len(board[y]); xPos++ {
		if board[y][xPos] != '.' {
			digits[board[y][xPos]-'1'] = true
		}
	}

	// Vertical
	for yPos := 0; yPos < len(board); yPos++ {
		if board[yPos][x] != '.' {
			digits[board[yPos][x]-'1'] = true
		}
	}

	// Square
	for yPos := y / 3 * 3; yPos < y/3*3+3; yPos++ {
		for xPos := x / 3 * 3; xPos < x/3*3+3; xPos++ {
			if board[yPos][xPos] != '.' {
				digits[board[yPos][xPos]-'1'] = true
			}
		}
	}

	for index, value := range digits {
		if !value {
			result = append(result, index+1)
		}
	}

	return result
}

func solveSudokuHelper(board [][]byte) bool {
	y, x := nextUnknownSquare(board)

	if (y == -1) || (x == -1) {
		return true
	}

	for _, digit := range possibleDigits(board, y, x) {
		board[y][x] = byte(digit) + '0'

		if solveSudokuHelper(board) {
			return true
		}

		board[y][x] = '.'
	}

	return false
}

func solveSudoku(board [][]byte) {
	solveSudokuHelper(board)
}
```

## Similar Questions:

- [Valid Sudoku](https://github.com/ju-popov/leetcode.com/tree/main/problems/valid-sudoku/)
- [Unique Paths III](https://github.com/ju-popov/leetcode.com/tree/main/problems/unique-paths-iii/)
