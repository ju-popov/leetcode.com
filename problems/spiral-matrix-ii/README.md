# 59. Spiral Matrix II

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Matrix](https://leetcode.com/tag/matrix/)
- [Simulation](https://leetcode.com/tag/simulation/)

## Problem:

Given a positive integer `n`, generate an `n x n` `matrix` filled with elements from `1` to `n2` in spiral order.

**Example 1:**

```
Input: n = 3
Output: [[1,2,3],[8,9,4],[7,6,5]]
```

**Example 2:**

```
Input: n = 1
Output: [[1]]
```

**Constraints:**

- `1 <= n <= 20`

## Solution:

```go
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
```

## Similar Questions:
  
- [Spiral Matrix](https://leetcode.com/problems/spiral-matrix/)
- [Spiral Matrix III](https://leetcode.com/problems/spiral-matrix-iii/)
