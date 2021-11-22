# 13. Roman to Integer

**Difficulty**: Easy

## Related Topics:

- [Hash Table](https://leetcode.com/tag/hash-table/)
- [Math](https://leetcode.com/tag/math/)
- [String](https://leetcode.com/tag/string/)

## Problem:

Roman numerals are represented by seven different symbols: `I`, `V`, `X`, `L`, `C`, `D` and `M`.

```
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```

For example, `2` is written as `II` in Roman numeral, just two one's added together. `12` is written as `XII`, which is simply `X + II`. The number `27` is written as `XXVII`, which is `XX + V + II`.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not `IIII`. Instead, the number four is written as `IV`. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as `IX`. There are six instances where subtraction is used:

- `I` can be placed before `V` (5) and `X` (10) to make 4 and 9.
- `X` can be placed before `L` (50) and `C` (100) to make 40 and 90.
- `C` can be placed before `D` (500) and `M` (1000) to make 400 and 900.

Given a roman numeral, convert it to an integer.

**Example 1:**

```
Input: s = "III"
Output: 3
```

**Example 2:**

```
Input: s = "IV"
Output: 4
```

**Example 3:**

```
Input: s = "IX"
Output: 9
```

**Example 4:**

```
Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
```

**Example 5:**

```
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```

**Constraints:**

- `1 <= s.length <= 15`
- `s` contains only the characters `('I', 'V', 'X', 'L', 'C', 'D', 'M')`.
- It is **guaranteed** that `s` is a valid roman numeral in the range `[1, 3999]`.

## Solution:

```go
func romanToInt(s string) int {
	result := 0
	length := len(s)
	left := 0

	for left < length {
		if s[left] == 'M' {
			result += 1000
			left++

			continue
		}

		if (left < length-1) && (s[left] == 'C') && (s[left+1] == 'M') {
			result += 900
			left += 2

			continue
		}

		if s[left] == 'D' {
			result += 500
			left++

			continue
		}

		if (left < length-1) && (s[left] == 'C') && (s[left+1] == 'D') {
			result += 400
			left += 2

			continue
		}

		if s[left] == 'C' {
			result += 100
			left++

			continue
		}

		if (left < length-1) && (s[left] == 'X') && (s[left+1] == 'C') {
			result += 90
			left += 2

			continue
		}

		if s[left] == 'L' {
			result += 50
			left++

			continue
		}

		if (left < length-1) && (s[left] == 'X') && (s[left+1] == 'L') {
			result += 40
			left += 2

			continue
		}

		if s[left] == 'X' {
			result += 10
			left++

			continue
		}

		if (left < length-1) && (s[left] == 'I') && (s[left+1] == 'X') {
			result += 9
			left += 2

			continue
		}

		if s[left] == 'V' {
			result += 5
			left++

			continue
		}

		if (left < length-1) && (s[left] == 'I') && (s[left+1] == 'V') {
			result += 4
			left += 2

			continue
		}

		if s[left] == 'I' {
			result++
			left++

			continue
		}
	}

	return result
}
```

## Similar Questions:

- [Integer to Roman](https://github.com/ju-popov/leetcode.com/tree/main/problems/integer-to-roman/)
