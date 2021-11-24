# 48. Rotate Image

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Math](https://leetcode.com/tag/math/)
- [Matrix](https://leetcode.com/tag/matrix/)

## Problem:

You are given an `n x n` 2D `matrix` representing an image, rotate the image by **90** degrees (clockwise).

You have to rotate the image [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm), which means you have to modify the input 2D matrix directly. **DO NOT** allocate another 2D matrix and do the rotation.

**Example 1:**

```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
```

**Example 2:**

```
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
```

**Example 3:**

```
Input: matrix = [[1]]
Output: [[1]]
```

**Example 4:**

```
Input: matrix = [[1,2],[3,4]]
Output: [[3,1],[4,2]]
```

**Constraints:**

- `matrix.length == n`
- `matrix[i].length == n`
- `1 <= n <= 20`
- `-1000 <= matrix[i][j] <= 1000`

## Solution:

```go
func rotate(matrix [][]int) {
	size := len(matrix)

	// Transpose
	for y := 0; y < size; y++ {
		for x := y + 1; x < size; x++ {
			matrix[y][x], matrix[x][y] = matrix[x][y], matrix[y][x]
		}
	}

	// Inverse
	for y := 0; y < size; y++ {
		for x := 0; x < size/2; x++ {
			matrix[y][x], matrix[y][size-1-x] = matrix[y][size-1-x], matrix[y][x]
		}
	}
}
```

## Similar Questions:

- [Determine Whether Matrix Can Be Obtained By Rotation](https://github.com/ju-popov/leetcode.com/tree/main/problems/determine-whether-matrix-can-be-obtained-by-rotation/)
