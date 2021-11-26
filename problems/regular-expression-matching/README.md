# 10. Regular Expression Matching

**Difficulty**: Hard

## Related Topics:

- [String](https://leetcode.com/tag/string/)
- [Dynamic Programming](https://leetcode.com/tag/dynamic-programming/)
- [Recursion](https://leetcode.com/tag/recursion/)

## Problem:

Given an input string `s` and a pattern `p`, implement regular expression matching with support for `'.'` and `'*'` where:

- `'.'` Matches any single character.​​​​
- `'*'` Matches zero or more of the preceding element.

The matching should cover the **entire** input string (not partial).

**Example 1:**

```
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

**Example 2:**

```
Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
```

**Example 3:**

```
Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
```

**Example 4:**

```
Input: s = "aab", p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
```

**Example 5:**

```
Input: s = "mississippi", p = "mis*is*p*."
Output: false
```

**Constraints:**

- `1 <= s.length <= 20`
- `1 <= p.length <= 30`
- `s` contains only lowercase English letters.
- `p` contains only lowercase English letters, `'.'`, and `'*'`.
- It is guaranteed for each appearance of the character `'*'`, there will be a previous valid character to match.

## Solution:

```go
type memoryKey struct {
	s string
	p string
}

func isMatchHelper(s string, p string, memory map[memoryKey]bool) bool {
	if p == "" {
		return s == ""
	}

	key := memoryKey{s: s, p: p}

	if value, ok := memory[key]; ok {
		return value
	}

	symbol := p[0]
	asterisk := len(p) > 1 && p[1] == '*'

	// (dot)*
	if (symbol == '.') && asterisk {
		result := false

		for counter := 0; (!result) && (counter <= len(s)); counter++ {
			result = isMatchHelper(s[counter:], p[2:], memory)
		}

		memory[key] = result

		return result
	}

	// (dot)
	if symbol == '.' {
		if s == "" {
			memory[key] = false

			return false
		}

		result := isMatchHelper(s[1:], p[1:], memory)
		memory[key] = result

		return result
	}

	// (symbol)*
	if asterisk {
		result := isMatchHelper(s, p[2:], memory)

		for counter := 0; (!result) && (counter < len(s)) && (s[counter] == symbol); counter++ {
			result = isMatchHelper(s[counter+1:], p[2:], memory)
		}

		memory[key] = result

		return result
	}

	// (symbol)
	if (s == "") || (s[0] != p[0]) {
		memory[key] = false

		return false
	}

	result := isMatchHelper(s[1:], p[1:], memory)
	memory[key] = result

	return result
}

func isMatch(s string, p string) bool {
	memory := make(map[memoryKey]bool)

	return isMatchHelper(s, p, memory)
}
```

## Similar Questions:

- [Wildcard Matching](https://github.com/ju-popov/leetcode.com/tree/main/problems/wildcard-matching/)
