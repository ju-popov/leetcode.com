package main

import "fmt"

/*

63. Unique Paths II

https://leetcode.com/problems/unique-paths-ii/

*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	height := len(obstacleGrid)
	width := len(obstacleGrid[0])

	if obstacleGrid[height-1][width-1] == 1 {
		return 0
	}

	for y := height - 1; y >= 0; y-- {
		for x := width - 1; x >= 0; x-- {
			if obstacleGrid[y][x] == 1 {
				obstacleGrid[y][x] = 0

				continue
			}

			if (y == height-1) && (x == width-1) {
				obstacleGrid[y][x] = 1

				continue
			}

			right := 0
			if x < width-1 {
				right = obstacleGrid[y][x+1]
			}

			bottom := 0
			if y < height-1 {
				bottom = obstacleGrid[y+1][x]
			}

			obstacleGrid[y][x] = right + bottom
		}
	}

	return obstacleGrid[0][0]
}

func main() {
	obstacleGrid1 := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	fmt.Println(uniquePathsWithObstacles(obstacleGrid1)) // 2

	obstacleGrid2 := [][]int{
		{0, 1},
		{0, 0},
	}

	fmt.Println(uniquePathsWithObstacles(obstacleGrid2)) // 1

	obstacleGrid3 := [][]int{
		{0, 0},
		{0, 1},
	}

	fmt.Println(uniquePathsWithObstacles(obstacleGrid3)) // 0
}
