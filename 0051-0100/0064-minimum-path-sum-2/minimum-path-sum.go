package main

import "fmt"

/*

64. Minimum Path Sum

https://leetcode.com/problems/minimum-path-sum/

*/

func minInt(v1 int, v2 int) int {
	if v1 <= v2 {
		return v1
	}

	return v2
}

func minPathSum(grid [][]int) int {
	heigth := len(grid)
	width := len(grid[0])

	for y := 0; y < heigth; y++ {
		for x := 0; x < width; x++ {
			if x == 0 {
				if y == 0 {
					// first column and row
					continue
				} else {
					// first column
					top := grid[y-1][x]
					grid[y][x] += top
				}
			} else {
				if y == 0 {
					// first row
					left := grid[y][x-1]
					grid[y][x] += left
				} else {
					// all else
					top := grid[y-1][x]
					left := grid[y][x-1]
					grid[y][x] += minInt(top, left)
				}
			}
		}
	}

	return grid[heigth-1][width-1]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})) // 7
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))            // 12
}
