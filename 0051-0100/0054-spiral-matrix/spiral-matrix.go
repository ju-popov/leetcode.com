package main

import "fmt"

/*

54. Spiral Matrix

https://leetcode.com/problems/spiral-matrix/

*/

func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)

	yMin := 0
	yMax := len(matrix) - 1
	xMin := 0
	xMax := len(matrix[0]) - 1

	for (yMin <= yMax) && (xMin <= xMax) {
		// Left to Right
		for x := xMin; x <= xMax; x++ {
			result = append(result, matrix[yMin][x])
		}

		// Up to Down
		for y := yMin + 1; y <= yMax-1; y++ {
			result = append(result, matrix[y][xMax])
		}

		// Right to Left
		if yMax > yMin {
			for x := xMax; x >= xMin; x-- {
				result = append(result, matrix[yMax][x])
			}
		}

		// Down to Left
		if xMax > xMin {
			for y := yMax - 1; y >= yMin+1; y-- {
				result = append(result, matrix[y][xMin])
			}
		}

		xMin++
		xMax--
		yMin++
		yMax--
	}

	return result
}

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	})) // [1 2 3 4 8 12 11 10 9 5 6 7]
}
