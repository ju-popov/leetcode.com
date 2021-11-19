package main

import "fmt"

/*

51. N-Queens

https://leetcode.com/problems/n-queens/

*/

func solution(n int, queens [][2]int) []string {
	result := make([]string, 0)

	for y := 0; y < n; y++ {
		row := ""

		for x := 0; x < n; x++ {
			found := false

			for _, queen := range queens {
				if queen[0] == y && queen[1] == x {
					found = true

					break
				}
			}

			if found {
				row += "Q"
			} else {
				row += "."
			}
		}

		result = append(result, row)
	}

	return result
}

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

func positionNQueens(n int, queens [][2]int, count int, result [][]string) [][]string {
	if count == 0 {
		result = append(result, solution(n, queens))

		return result
	}

	y := len(queens)

	for _, x := range possibleXPositions(n, y, queens) {
		newQueens := [][2]int{{y, x}}
		newQueens = append(newQueens, queens...)

		result = positionNQueens(n, newQueens, count-1, result)
	}

	return result
}

func solveNQueens(n int) [][]string {
	queens := make([][2]int, 0)
	results := make([][]string, 0)

	return positionNQueens(n, queens, n, results)
}

func main() {
	fmt.Println(solveNQueens(4))
	fmt.Println(solveNQueens(1))
}
