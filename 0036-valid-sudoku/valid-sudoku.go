package main

import "fmt"

/*

36. Valid Sudoku

https://leetcode.com/problems/valid-sudoku/

*/

func isValidArea(board [][]byte, yStart int, yEnd int, xStart int, xEnd int) bool {
	digits := make(map[byte]bool)

	fmt.Println(xStart, xEnd, yStart, yEnd)

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

func main() {
	board1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board2 := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board3 := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}

	fmt.Println(isValidSudoku(board1)) // true
	fmt.Println(isValidSudoku(board2)) // false
	fmt.Println(isValidSudoku(board3)) // false
}
