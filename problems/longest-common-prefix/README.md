# 14. Longest Common Prefix

**Difficulty**: Easy

## Related Topics:

- [String](https://leetcode.com/tag/string/)

## Problem:

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string `""`.

**Example 1:**

```
Input: strs = ["flower","flow","flight"]
Output: "fl"
```

**Example 2:**

```
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```

**Constraints:**

- `1 <= strs.length <= 200`
- `0 <= strs[i].length <= 200`
- `strs[i]` consists of only lower-case English letters.

## Solution:

```go
func minInt(values ...int) int {
	result := values[0]

	for index := 1; index < len(values); index++ {
		if values[index] < result {
			result = values[index]
		}
	}

	return result
}

func longestCommonPrefixPosition(str1 string, str2 string) int {
	result := -1

	for index := 0; index < minInt(len(str1), len(str2)); index++ {
		if str1[index] == str2[index] {
			result = index

			continue
		}

		break
	}

	return result
}

func longestCommonPrefix(strs []string) string {
	result := strs[0]

	for index := 1; index < len(strs); index++ {
		pos := longestCommonPrefixPosition(result, strs[index])
		if pos < 0 {
			return ""
		}

		result = result[:pos+1]
	}

	return result
}
```
