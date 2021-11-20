package main

import "fmt"

/*

63. Unique Paths II

https://leetcode.com/problems/unique-paths-ii/

*/

type memoryKey struct {
	y int
	x int
}

func uniquePathsWithObstaclesHelper(obstacleGrid [][]int, y int, x int, memory map[memoryKey]int) int {
	// Otside
	if (y >= len(obstacleGrid)) || (x >= len(obstacleGrid[y])) {
		return 0
	}

	// Obstacle
	if obstacleGrid[y][x] == 1 {
		return 0
	}

	// Destination
	if (y == len(obstacleGrid)-1) && (x == len(obstacleGrid[y])-1) {
		return 1
	}

	key := memoryKey{y: y, x: x}

	if value, ok := memory[key]; ok {
		return value
	}

	right := uniquePathsWithObstaclesHelper(obstacleGrid, y, x+1, memory)
	bottom := uniquePathsWithObstaclesHelper(obstacleGrid, y+1, x, memory)

	result := right + bottom

	memory[key] = result

	return result
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	memory := make(map[memoryKey]int)

	return uniquePathsWithObstaclesHelper(obstacleGrid, 0, 0, memory)
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
