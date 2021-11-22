# 17. Letter Combinations of a Phone Number

**Difficulty**: Medium

## Related Topics:

- [Hash Table](https://leetcode.com/tag/hash-table/)
- [String](https://leetcode.com/tag/string/)
- [Backtracking](https://leetcode.com/tag/backtracking/)

## Problem:

Given a string containing digits from `2-9` inclusive, return all possible letter combinations that the number could represent. Return the answer in **any order**.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

**Example 1:**

```
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

**Example 2:**

```
Input: digits = ""
Output: []
```

**Example 3:**

```
Input: digits = "2"
Output: ["a","b","c"]
```

**Constraints:**

- `0 <= digits.length <= 4`
- `digits[i]` is a digit in the range `['2', '9']`.

## Solution:

```go
var digitsMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	results := make([]string, 0)
	if len(digits) == 0 {
		return results
	}

	// Add one empty result
	results = append(results, []string{""}...)

	for index := len(digits) - 1; index >= 0; index-- {
		digit := digits[index]
		newResults := make([]string, 0)

		for _, char := range digitsMap[digit] {
			for _, result := range results {
				newResults = append(newResults, char+result)
			}
		}

		results = newResults
	}

	return results
}
```

## Similar Questions:

- [Generate Parentheses](https://github.com/ju-popov/leetcode.com/tree/main/problems/generate-parentheses/)
- [Combination Sum](https://github.com/ju-popov/leetcode.com/tree/main/problems/combination-sum/)
- [Binary Watch](https://github.com/ju-popov/leetcode.com/tree/main/problems/binary-watch/)
