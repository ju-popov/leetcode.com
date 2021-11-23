# 461. Hamming Distance

**Difficulty**: Easy

## Related Topics:

- [Bit Manipulation](https://leetcode.com/tag/bit-manipulation/)

## Problem:

The [Hamming distance](https://en.wikipedia.org/wiki/Hamming_distance) between two integers is the number of positions at which the corresponding bits are different.

Given two integers `x` and `y`, return *the **Hamming distance** between them*.

**Example 1:**

```
Input: x = 1, y = 4
Output: 2
Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
The above arrows point to positions where the corresponding bits are different.
```

**Example 2:**

```
Input: x = 3, y = 1
Output: 1
```

**Constraints:**

- `0 <= x, y <= 231 - 1`

## Solution:

```go
func hammingDistance(x int, y int) int {
	result := 0

	for (x > 0) || (y > 0) {
		if x%2 != y%2 {
			result++
		}

		x /= 2
		y /= 2
	}

	return result
}
```

## Similar Questions:

- [Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/)
- [Total Hamming Distance](https://leetcode.com/problems/total-hamming-distance/)
