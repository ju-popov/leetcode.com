package main

import "fmt"

/*

79. Word Search

https://leetcode.com/problems/word-search/

#array #backtracking #matrix

*/

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

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED")) // true

	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "SEE")) // true

	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCB")) // false

	fmt.Println(exist([][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}, "AAB")) // true

	fmt.Println(exist([][]byte{
		{'a', 'a', 'b', 'a', 'a', 'b'},
		{'b', 'a', 'b', 'a', 'b', 'b'},
		{'b', 'a', 'b', 'b', 'b', 'b'},
		{'a', 'a', 'b', 'a', 'b', 'a'},
		{'b', 'b', 'a', 'a', 'a', 'b'},
		{'b', 'b', 'b', 'a', 'b', 'a'},
	}, "aaaababab")) // true
}
