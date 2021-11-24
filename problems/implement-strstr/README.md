# 28. Implement strStr()

**Difficulty**: Easy

## Related Topics:

- [Two Pointers](https://leetcode.com/tag/two-pointers/)
- [String](https://leetcode.com/tag/string/)
- [String Matching](https://leetcode.com/tag/string-matching/)

## Problem:

Implement [strStr()](http://www.cplusplus.com/reference/cstring/strstr/).

Return the index of the first occurrence of needle in haystack, or `-1` if `needle` is not part of `haystack`.

**Clarification:**

What should we return when `needle` is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when `needle` is an empty string. This is consistent to C's [strstr()](http://www.cplusplus.com/reference/cstring/strstr/) and Java's [indexOf()](https://docs.oracle.com/javase/7/docs/api/java/lang/String.html#indexOf(java.lang.String)).

**Example 1:**

```
Input: haystack = "hello", needle = "ll"
Output: 2
```

**Example 2:**

```
Input: haystack = "aaaaa", needle = "bba"
Output: -1
```

**Example 3:**

```
Input: haystack = "", needle = ""
Output: 0
```

**Constraints:**

- `0 <= haystack.length, needle.length <= 5 * 104`
- `haystack` and `needle` consist of only lower-case English characters.

## Solution:

```go
func isMatch(haystack string, needle string) bool {
	for index := 0; index < len(needle); index++ {
		if haystack[index] != needle[index] {
			return false
		}
	}

	return true
}

func strStr(haystack string, needle string) int {
	for index := 0; index < len(haystack)-len(needle)+1; index++ {
		if isMatch(haystack[index:], needle) {
			return index
		}
	}

	return -1
}
```

## Similar Questions:

- [Shortest Palindrome](https://github.com/ju-popov/leetcode.com/tree/main/problems/shortest-palindrome/)
- [Repeated Substring Pattern](https://github.com/ju-popov/leetcode.com/tree/main/problems/repeated-substring-pattern/)
