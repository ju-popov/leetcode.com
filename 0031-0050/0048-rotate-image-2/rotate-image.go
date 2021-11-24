package main

import "fmt"

/*

48. Rotate Image

https://leetcode.com/problems/rotate-image/

*/

func rotate(matrix [][]int) {
	size := len(matrix)

	// Transpose
	for y := 0; y < size; y++ {
		for x := y + 1; x < size; x++ {
			matrix[y][x], matrix[x][y] = matrix[x][y], matrix[y][x]
		}
	}

	// Inverse
	for y := 0; y < size; y++ {
		for x := 0; x < size/2; x++ {
			matrix[y][x], matrix[y][size-1-x] = matrix[y][size-1-x], matrix[y][x]
		}
	}
}

func main() {
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(matrix1)

	fmt.Println(matrix1) // [[7 4 1] [8 5 2] [9 6 3]]

	matrix2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	rotate(matrix2)

	fmt.Println(matrix2) // [[15 13 2 5] [14 3 4 1] [12 6 8 9] [16 7 10 11]]

	matrix3 := [][]int{
		{1},
	}

	rotate(matrix3)

	fmt.Println(matrix3) // [[1]]

	matrix4 := [][]int{
		{1, 2},
		{3, 4},
	}

	rotate(matrix4)

	fmt.Println(matrix4) // [[3 1] [4 2]]
}
