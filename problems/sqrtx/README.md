# 69. Sqrt(x)

**Difficulty**: Easy

## Related Topics:

- [Math](https://leetcode.com/tag/math/)
- [Binary Search](https://leetcode.com/tag/binary-search/)

## Problem:

Given a non-negative integer `x`, compute and return *the square root of* `x`.

Since the return type is an integer, the decimal digits are **truncated**, and only **the integer part** of the result is returned.

**Note:**You are not allowed to use any built-in exponent function or operator, such as `pow(x, 0.5)` or `x ** 0.5`.

**Example 1:**

```
Input: x = 4
Output: 2
```

**Example 2:**

```
Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since the decimal part is truncated, 2 is returned.
```

**Constraints:**

- `0 <= x <= 231 - 1`

## Solution:

```go
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left := 1
	right := x / 2

	for left <= right {
		mid := left + (right-left)/2
		value := mid * mid

		if value == x {
			return mid
		}

		if value > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}
```

## Similar Questions:

- [Pow(x, n)](https://github.com/ju-popov/leetcode.com/tree/main/problems/powx-n/)
- [Valid Perfect Square](https://github.com/ju-popov/leetcode.com/tree/main/problems/valid-perfect-square/)
