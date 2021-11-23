# 54. Spiral Matrix

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Matrix](https://leetcode.com/tag/matrix/)
- [Simulation](https://leetcode.com/tag/simulation/)

## Problem:

Given an `m x n` `matrix`, return *all elements of the* `matrix` *in spiral order*.

**Example 1:**

```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
```

**Example 2:**

```
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
```

**Constraints:**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 10`
- `-100 <= matrix[i][j] <= 100`

## Solution:

```go
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
```

## Similar Questions:

- [Spiral Matrix II](https://github.com/ju-popov/leetcode.com/tree/main/problems/spiral-matrix-ii/)
- [Spiral Matrix III](https://github.com/ju-popov/leetcode.com/tree/main/problems/spiral-matrix-iii/)
