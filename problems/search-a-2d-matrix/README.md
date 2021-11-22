# 74. Search a 2D Matrix

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Binary Search](https://leetcode.com/tag/binary-search/)
- [Matrix](https://leetcode.com/tag/matrix/)


## Problem:

Write an efficient algorithm that searches for a value in an `m x n` matrix. This matrix has the following properties:

- Integers in each row are sorted from left to right.
- The first integer of each row is greater than the last integer of the previous row.

**Example 1:**

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
```

**Example 2:**

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
```

**Constraints:**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 100`
- `-104 <= matrix[i][j], target <= 104`

## Solution:

```go
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
```

## Similar Questions:

- [Search a 2D Matrix II](https://github.com/ju-popov/leetcode.com/tree/main/problems/search-a-2d-matrix-ii/)
