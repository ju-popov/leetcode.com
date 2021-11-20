package main

import "fmt"

/*

59. Spiral Matrix II

https://leetcode.com/problems/spiral-matrix-ii/

*/

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	min := 0
	max := n - 1
	counter := 1

	for min <= max {
		// left to right
		y := min

		for x := min; x <= max; x++ {
			result[y][x] = counter
			counter++
		}

		// top to bottom
		x := max

		for y := min + 1; y <= max-1; y++ {
			result[y][x] = counter
			counter++
		}

		// right to left
		if min != max {
			y := max

			for x := max; x >= min; x-- {
				result[y][x] = counter
				counter++
			}
		}

		// botton to left
		if min != max {
			x := min

			for y := max - 1; y >= min+1; y-- {
				result[y][x] = counter
				counter++
			}
		}

		min++
		max--
	}

	return result
}

func main() {
	fmt.Println(generateMatrix(1)) // [[1]]
	fmt.Println(generateMatrix(2)) // [[1 2] [4 3]]
	fmt.Println(generateMatrix(3)) // [[1 2 3] [8 9 4] [7 6 5]]
	fmt.Println(generateMatrix(4)) // [[1 2 3 4] [12 13 14 5] [11 16 15 6] [10 9 8 7]]
}
