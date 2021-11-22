package main

import "fmt"

/*

74. Search a 2D Matrix

https://leetcode.com/problems/search-a-2d-matrix/

#array #binary-search #matrix

*/

func searchMatrix(matrix [][]int, target int) bool {
	heigth := len(matrix)
	width := len(matrix[0])

	left := 0
	right := width*heigth - 1

	for left <= right {
		mid := left + (right-left)/2
		y := mid / width
		x := mid % width
		value := matrix[y][x]

		if value == target {
			return true
		}

		if value < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix(
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}, 3,
	)) // true

	fmt.Println(searchMatrix(
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}, 13,
	)) // false
}
