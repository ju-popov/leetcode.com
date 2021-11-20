package main

import "fmt"

/*

64. Minimum Path Sum

https://leetcode.com/problems/minimum-path-sum/

*/

type memoryKey struct {
	y int
	x int
}

func minInt(v1 int, v2 int) int {
	if v1 <= v2 {
		return v1
	}

	return v2
}

func minPathSumHelper(grid [][]int, y int, x int, memory map[memoryKey]int) int {
	key := memoryKey{y: y, x: x}

	if value, ok := memory[key]; ok {
		return value
	}

	var result int

	if y == len(grid)-1 {
		if x == len(grid[y])-1 {
			// Last row and column
			result = grid[y][x]
		} else {
			// Last row
			right := minPathSumHelper(grid, y, x+1, memory)
			result = grid[y][x] + right
		}
	} else {
		if x == len(grid[y])-1 {
			// Last column
			down := minPathSumHelper(grid, y+1, x, memory)
			result = grid[y][x] + down
		} else {
			// Everything else
			right := minPathSumHelper(grid, y, x+1, memory)
			down := minPathSumHelper(grid, y+1, x, memory)
			result = grid[y][x] + minInt(right, down)
		}
	}

	memory[key] = result

	return result
}

func minPathSum(grid [][]int) int {
	memory := make(map[memoryKey]int)

	return minPathSumHelper(grid, 0, 0, memory)
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})) // 7
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))            // 12
}
