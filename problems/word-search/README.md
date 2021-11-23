# 79. Word Search

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Backtracking](https://leetcode.com/tag/backtracking/)
- [Matrix](https://leetcode.com/tag/matrix/)

## Problem:

Given an `m x n` grid of characters `board` and a string `word`, return `true` *if* `word` *exists in the grid*.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

**Example 1:**

```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
```

**Example 2:**

```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true
```

**Example 3:**

```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false
```

**Constraints:**

- `m == board.length`
- `n = board[i].length`
- `1 <= m, n <= 6`
- `1 <= word.length <= 15`
- `board` and `word` consists of only lowercase and uppercase English letters.

**Follow up:** Could you use search pruning to make your solution faster with a larger `board`?

## Solution:

```go
var neighbors = [][]int{
	{0, -1}, // left
	{0, +1}, // right
	{+1, 0}, // up
	{-1, 0}, // down
}

func existHelper(board [][]byte, word string, y int, x int) bool {
	if len(word) == 1 {
		return true
	}

	height := len(board)
	width := len(board[0])

	// Mark as visited
	board[y][x] = ' '

	for _, neighbor := range neighbors {
		neighborY := y + neighbor[0]
		neighborX := x + neighbor[1]

		if (neighborY >= 0) &&
			(neighborY < height) &&
			(neighborX >= 0) &&
			(neighborX < width) &&
			board[neighborY][neighborX] == word[1] &&
			existHelper(board, word[1:], neighborY, neighborX) {
			return true
		}
	}

	// Remove mark
	board[y][x] = word[0]

	return false
}

func exist(board [][]byte, word string) bool {
	height := len(board)
	width := len(board[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if board[y][x] == word[0] {
				if existHelper(board, word, y, x) {
					return true
				}
			}
		}
	}

	return false
}
```

## Similar Questions:

- [Word Search II](https://github.com/ju-popov/leetcode.com/tree/main/problems/word-search-ii/)
