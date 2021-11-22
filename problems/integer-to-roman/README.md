# 12. Integer to Roman

**Difficulty**: Medium

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

Given an integer, convert it to a roman numeral.

**Example 1:**

```
Input: num = 3
Output: "III"
```

**Example 2:**

```
Input: num = 4
Output: "IV"
```

**Example 3:**

```
Input: num = 9
Output: "IX"
```

**Example 4:**

```
Input: num = 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
```

**Example 5:**

```
Input: num = 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```

**Constraints:**

- `1 <= num <= 3999`

## Solution:

```go
func intToRoman(num int) string {
	result := ""

	for num >= 1000 {
		result += "M"
		num -= 1000
	}

	for num >= 900 {
		result += "CM"
		num -= 900
	}

	for num >= 500 {
		result += "D"
		num -= 500
	}

	for num >= 400 {
		result += "CD"
		num -= 400
	}

	for num >= 100 {
		result += "C"
		num -= 100
	}

	for num >= 90 {
		result += "XC"
		num -= 90
	}

	for num >= 50 {
		result += "L"
		num -= 50
	}

	for num >= 40 {
		result += "XL"
		num -= 40
	}

	for num >= 10 {
		result += "X"
		num -= 10
	}

	for num >= 9 {
		result += "IX"
		num -= 9
	}

	for num >= 5 {
		result += "V"
		num -= 5
	}

	for num >= 4 {
		result += "IV"
		num -= 4
	}

	for num >= 1 {
		result += "I"
		num--
	}

	return result
}
```

## Similar Questions:

- [Roman to Integer](https://github.com/ju-popov/leetcode.com/tree/main/problems/roman-to-integer/)
- [Integer to English Words](https://github.com/ju-popov/leetcode.com/tree/main/problems/integer-to-english-words/)
