package main

import "fmt"

/*

73. Set Matrix Zeroes

https://leetcode.com/problems/set-matrix-zeroes/

#array #hash-table #matrix

*/

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

func main() {
	matrix1 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	setZeroes(matrix1) // [[1 0 1] [0 0 0] [1 0 1]]

	fmt.Println(matrix1)

	matrix2 := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	setZeroes(matrix2)

	fmt.Println(matrix2) // [[0 0 0 0] [0 4 5 0] [0 3 1 0]]
}
