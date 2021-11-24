# 50. Pow(x, n)

**Difficulty**: Medium

## Related Topics:

- [Math](https://leetcode.com/tag/math/)
- [Recursion](https://leetcode.com/tag/recursion/)

## Problem:

Implement [pow(x, n)](http://www.cplusplus.com/reference/valarray/pow/), which calculates `x` raised to the power `n` (i.e., `xn`).

**Example 1:**

```
Input: x = 2.00000, n = 10
Output: 1024.00000
```

**Example 2:**

```
Input: x = 2.10000, n = 3
Output: 9.26100
```

**Example 3:**

```
Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
```

**Constraints:**

- `-100.0 < x < 100.0`
- `-231 <= n <= 231-1`
- `-104 <= xn <= 104`

## Solution:

```go
func myPowHelper(x float64, n int, memory map[int]float64) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if value, ok := memory[n]; ok {
		return value
	}

	result := myPowHelper(x, n/2, memory) * myPowHelper(x, n-n/2, memory)

	memory[n] = result

	return result
}

func myPow(x float64, n int) float64 {
	memory := make(map[int]float64)

	if n >= 0 {
		return myPowHelper(x, n, memory)
	}

	return 1 / myPowHelper(x, -n, memory)
}
```

## Similar Questions:

- [Sqrt(x)](https://github.com/ju-popov/leetcode.com/tree/main/problems/sqrtx/)
- [Super Pow](https://github.com/ju-popov/leetcode.com/tree/main/problems/super-pow/)
