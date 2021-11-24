# 29. Divide Two Integers

**Difficulty**: Medium

## Related Topics:

- [Math](https://leetcode.com/tag/math/)
- [Bit Manipulation](https://leetcode.com/tag/bit-manipulation/)

## Problem:

Given two integers `dividend` and `divisor`, divide two integers **without** using multiplication, division, and mod operator.

The integer division should truncate toward zero, which means losing its fractional part. For example, `8.345` would be truncated to `8`, and `-2.7335` would be truncated to `-2`.

Return *the **quotient** after dividing*`dividend`*by*`divisor`.

**Note:**Assume we are dealing with an environment that could only store integers within the **32-bit** signed integer range: `[−231, 231 − 1]`. For this problem, if the quotient is **strictly greater than** `231 - 1`, then return `231 - 1`, and if the quotient is **strictly less than** `-231`, then return `-231`.

**Example 1:**

```
Input: dividend = 10, divisor = 3
Output: 3
Explanation: 10/3 = 3.33333.. which is truncated to 3.
```

**Example 2:**

```
Input: dividend = 7, divisor = -3
Output: -2
Explanation: 7/-3 = -2.33333.. which is truncated to -2.
```

**Example 3:**

```
Input: dividend = 0, divisor = 1
Output: 0
```

**Example 4:**

```
Input: dividend = 1, divisor = 1
Output: 1
```

**Constraints:**

- `-231 <= dividend, divisor <= 231 - 1`
- `divisor != 0`

## Solution:

```go
const (
	maxInt32 = 2147483647
	minInt32 = -2147483648
)

func abs(value int) int {
	if value >= 0 {
		return value
	}

	return -value
}

func addResult(result int32, value int32) int32 {
	if value > 0 {
		if maxInt32-result < value {
			return maxInt32
		}
	} else {
		if minInt32-result > value {
			return minInt32
		}
	}

	return result + value
}

func divideHelper(dividend int, positiveDivisor int) (int32, int) {
	var result int32

	if abs(dividend) < positiveDivisor {
		return 0, dividend
	}

	if abs(dividend) >= positiveDivisor+positiveDivisor {
		var value int32
		value, dividend = divideHelper(dividend, positiveDivisor+positiveDivisor)
		result = addResult(result, value)
		result = addResult(result, value)
	}

	if abs(dividend) >= positiveDivisor {
		if dividend > 0 {
			dividend -= positiveDivisor
			result = addResult(result, 1)
		} else if dividend < 0 {
			dividend += positiveDivisor
			result = addResult(result, -1)
		}
	}

	return result, dividend
}

func divide(dividend int, divisor int) int {
	if divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	}

	result, _ := divideHelper(dividend, divisor)

	return int(result)
}
```
