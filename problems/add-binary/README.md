# 67. Add Binary

**Difficulty**: Easy

## Related Topics:

- [Math](https://leetcode.com/tag/math/)
- [String](https://leetcode.com/tag/string/)
- [Bit Manipulation](https://leetcode.com/tag/bit-manipulation/)
- [Simulation](https://leetcode.com/tag/simulation/)

## Problem:

Given two binary strings `a` and `b`, return *their sum as a binary string*.

**Example 1:**

```
Input: a = "11", b = "1"
Output: "100"
```

**Example 2:**

```
Input: a = "1010", b = "1011"
Output: "10101"
```

**Constraints:**

- `1 <= a.length, b.length <= 104`
- `a` and `b` consist only of `'0'` or `'1'` characters.
- Each string does not contain leading zeros except for the zero itself.

## Solution:

```go
func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func addBinary(a string, b string) string {
	result := make([]byte, max(len(a), len(b)))
	indexA := len(a) - 1
	indexB := len(b) - 1
	indexResult := len(result) - 1
	carry := byte(0)

	for (indexA >= 0) || (indexB >= 0) {
		sum := carry

		if indexA >= 0 {
			sum += a[indexA] - '0'
		}

		if indexB >= 0 {
			sum += b[indexB] - '0'
		}

		carry = sum / 2
		result[indexResult] = '0' + sum%2

		indexA--
		indexB--
		indexResult--
	}

	if carry > 0 {
		return "1" + string(result)
	}

	return string(result)
}
```

## Similar Questions:

- [Add Two Numbers](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-two-numbers/)
- [Multiply Strings](https://github.com/ju-popov/leetcode.com/tree/main/problems/multiply-strings/)
- [Plus One](https://github.com/ju-popov/leetcode.com/tree/main/problems/plus-one/)
- [Add to Array-Form of Integer](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-to-array-form-of-integer/)
