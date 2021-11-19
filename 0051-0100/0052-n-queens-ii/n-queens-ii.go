package main

import "fmt"

/*

	52. N-Queens II

	https://leetcode.com/problems/n-queens-ii/

*/

func possibleXPositions(n int, y int, queens [][2]int) []int {
	result := make([]int, 0)

	for x := 0; x < n; x++ {
		threatened := false

		for _, queen := range queens {
			queenY := queen[0]
			queenX := queen[1]

			// Horizontal or Vertical
			if (queenX == x) || (queenY == y) {
				threatened = true

				break
			}

			// Diagonal \
			if (queenX - queenY) == (x - y) {
				threatened = true

				break
			}

			// Diagonal /
			if (queenX + queenY) == (x + y) {
				threatened = true

				break
			}
		}

		if !threatened {
			result = append(result, x)
		}
	}

	return result
}

func positionNQueens(n int, queens [][2]int, count int) int {
	result := 0

	if count == 0 {
		return 1
	}

	y := len(queens)

	for _, x := range possibleXPositions(n, y, queens) {
		newQueens := [][2]int{{y, x}}
		newQueens = append(newQueens, queens...)

		result += positionNQueens(n, newQueens, count-1)
	}

	return result
}

func totalNQueens(n int) int {
	queens := make([][2]int, 0)

	return positionNQueens(n, queens, n)
}

func main() {
	fmt.Println(totalNQueens(4)) // 2
	fmt.Println(totalNQueens(1)) // 1
}
