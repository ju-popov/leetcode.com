package main

/*

51. N-Queens

https://leetcode.com/problems/n-queens/

*/

func isThreatend(queens [][2]int, posX int, posY int) bool {
	for _, queen := range queens {
		queenX := queen[0]
		queenY := queen[1]

		// Horizontal or Vertical
		if (queenX == posX) || (queenY == posY) {
			return true
		}

		// Diagonal \
		if (queenX - queenY) == (posX - posY) {
			return true
		}

		// Diagonal /
		if (queenX + queenY) == (posX + posY) {
			return true
		}
	}

	return false
}

func positionNQueens(n int, queens [][2]int, count int, result [][]string) [][]string {
	if count == 0 {
		return result
	}

}

func solveNQueens(n int) [][]string {
	queens := make([][2]int, 0)
	results := make([][]string, 0)

	return positionNQueens(n, queens, n, results)
}
