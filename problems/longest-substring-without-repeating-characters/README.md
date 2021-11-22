# 3. Longest Substring Without Repeating Characters

**Difficulty**: Medium

## Related Topics:

- [Hash Table](https://leetcode.com/tag/hash-table/)
- [String](https://leetcode.com/tag/string/)
- [Sliding Window](https://leetcode.com/tag/sliding-window/)

## Problem:

Given a string `s`, find the length of the **longest substring** without repeating characters.

**Example 1:**

```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

**Example 2:**

```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

**Example 3:**

```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

**Example 4:**

```
Input: s = ""
Output: 0
```

**Constraints:**

- `0 <= s.length <= 5 * 104`
- `s` consists of English letters, digits, symbols and spaces.

## Solution:

```go
func maxInt(values ...int) int {
	result := values[0]

	for _, value := range values {
		if value > result {
			result = value
		}
	}

	return result
}

func lengthOfLongestSubstring(s string) int {
	result := 0
	visited := make(map[byte]bool)

	left := -1
	right := 0

	for right < len(s) {
		char := s[right]

		if visited[char] {
			for {
				left++

				visited[s[left]] = false

				if s[left] == char {
					break
				}
			}
		}

		visited[char] = true

		result = maxInt(result, right-left)

		right++
	}

	return result
}
```

# Similar Questions:

- [Longest Substring with At Most Two Distinct Characters](https://github.com/ju-popov/leetcode.com/tree/main/problems/longest-substring-with-at-most-two-distinct-characters/)
- [Longest Substring with At Most K Distinct Characters](https://github.com/ju-popov/leetcode.com/tree/main/problems/longest-substring-with-at-most-k-distinct-characters/)
- [Subarrays with K Different Integers](https://github.com/ju-popov/leetcode.com/tree/main/problems/subarrays-with-k-different-integers/)
- [Maximum Erasure Value](https://github.com/ju-popov/leetcode.com/tree/main/problems/maximum-erasure-value/)
- [Number of Equal Count Substrings](https://github.com/ju-popov/leetcode.com/tree/main/problems/number-of-equal-count-substrings/)
