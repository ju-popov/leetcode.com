package main

import "fmt"

/*

37. Sudoku Solver

https://leetcode.com/problems/sudoku-solver/

*/

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

func printBoard(board [][]byte) {
	for y := 0; y < len(board); y++ {
		fmt.Println(string(board[y]))
	}
}

func main() {
	board := [][]byte{
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

	solveSudoku(board)

	// 534678912
	// 672195348
	// 198342567
	// 859761423
	// 426853791
	// 713924856
	// 961537284
	// 287419635
	// 345286179

	printBoard(board)
}
