# 7. Reverse Integer

**Difficulty**: Medium

## Related Topics:

- [Math](https://leetcode.com/tag/math/)

## Problem:

Given a signed 32-bit integer `x`, return `x`*with its digits reversed*. If reversing `x` causes the value to go outside the signed 32-bit integer range `[-231, 231 - 1]`, then return `0`.

**Assume the environment does not allow you to store 64-bit integers (signed or unsigned).**

**Example 1:**

```
Input: x = 123
Output: 321
```

**Example 2:**

```
Input: x = -123
Output: -321
```

**Example 3:**

```
Input: x = 120
Output: 21
```

**Example 4:**

```
Input: x = 0
Output: 0
```

**Constraints:**

- `-231 <= x <= 231 - 1`

## Solution:

```go
const (
	maxInt32 = 2147483647
	minInt32 = -2147483648
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	result := int32(0)

	for x != 0 {
		digit := int32(x % 10)

		if result > maxInt32/10 {
			return 0
		}

		if (result == maxInt32/10) && (digit > maxInt32%10) {
			return 0
		}

		if result < minInt32/10 {
			return 0
		}

		if (result == minInt32/10) && (digit < minInt32%10) {
			return 0
		}

		result = result*10 + digit
		x /= 10
	}

	return int(result)
}
```

## Similar Questions:

- [String to Integer (atoi)](https://github.com/ju-popov/leetcode.com/tree/main/problems/string-to-integer-atoi/)
- [Reverse Bits](https://github.com/ju-popov/leetcode.com/tree/main/problems/reverse-bits/)
