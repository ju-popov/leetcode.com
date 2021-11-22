# 20. Valid Parentheses

**Difficulty**: Easy

## Related Topics:

- [String](https://leetcode.com/tag/string/)
- [Stack](https://leetcode.com/tag/stack/)

## Problem:

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

**Example 1:**

```
Input: s = "()"
Output: true
```

**Example 2:**

```
Input: s = "()[]{}"
Output: true
```

**Example 3:**

```
Input: s = "(]"
Output: false
```

**Example 4:**

```
Input: s = "([)]"
Output: false
```

**Example 5:**

```
Input: s = "{[]}"
Output: true
```

**Constraints:**

- `1 <= s.length <= 104`
- `s` consists of parentheses only `'()[]{}'`.

## Solution:

```go
var parentheses = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, char := range s {
		if value, ok := parentheses[char]; ok {
			// opening
			stack = append(stack, value)
		} else {
			// closing
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] != char {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
```

## Similar Questions:

- [Generate Parentheses](https://github.com/ju-popov/leetcode.com/tree/main/problems/generate-parentheses/)
- [Longest Valid Parentheses](https://github.com/ju-popov/leetcode.com/tree/main/problems/longest-valid-parentheses/)
- [Remove Invalid Parentheses](https://github.com/ju-popov/leetcode.com/tree/main/problems/remove-invalid-parentheses/)
- [Check If Word Is Valid After Substitutions](https://github.com/ju-popov/leetcode.com/tree/main/problems/check-if-word-is-valid-after-substitutions/)
