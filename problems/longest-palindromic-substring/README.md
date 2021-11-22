# 5. Longest Palindromic Substring

**Difficulty**: Medium

## Related Topics:

- [String](https://leetcode.com/tag/string/)
- [Dynamic Programming](https://leetcode.com/tag/dynamic-programming/)

## Problem:

Given a string `s`, return *the longest palindromic substring* in `s`.

**Example 1:**

```
Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```

**Example 2:**

```
Input: s = "cbbd"
Output: "bb"
```

**Example 3:**

```
Input: s = "a"
Output: "a"
```

**Example 4:**

```
Input: s = "ac"
Output: "a"
```

**Constraints:**

- `1 <= s.length <= 1000`
- `s` consist of only digits and English letters.

## Solution:

```go
func expandAroundCenter(s string, length int, left int, right int) (int, int) {
	resultLeft := 0
	resultRight := 0

	for (left >= 0) && (right <= length-1) {
		if s[left] == s[right] {
			resultLeft = left
			resultRight = right
			left--
			right++
		} else {
			break
		}
	}

	return resultLeft, resultRight
}

func longestPalindrome(s string) string {
	length := len(s)

	if length <= 1 {
		return s
	}

	resultLeft := 0
	resultRight := 0

	for index := 0; index < length; index++ {
		left, right := expandAroundCenter(s, length, index, index)
		if (right - left) > (resultRight - resultLeft) {
			resultLeft, resultRight = left, right
		}
	}

	for index := 0; index < length-1; index++ {
		left, right := expandAroundCenter(s, length, index, index+1)
		if (right - left) > (resultRight - resultLeft) {
			resultLeft, resultRight = left, right
		}
	}

	return s[resultLeft : resultRight+1]
}
```

## Similar Questions:

- [Shortest Palindrome](https://github.com/ju-popov/leetcode.com/tree/main/problems/shortest-palindrome/)
- [Palindrome Permutation](https://github.com/ju-popov/leetcode.com/tree/main/problems/palindrome-permutation/)
- [Palindrome Pairs](https://github.com/ju-popov/leetcode.com/tree/main/problems/palindrome-pairs/)
- [Longest Palindromic Subsequence](https://github.com/ju-popov/leetcode.com/tree/main/problems/longest-palindromic-subsequence/)
- [Palindromic Substrings](https://github.com/ju-popov/leetcode.com/tree/main/problems/palindromic-substrings/)
