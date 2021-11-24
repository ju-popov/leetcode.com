# 49. Group Anagrams

**Difficulty**: Medium

## Related Topics:

- [Hash Table](https://leetcode.com/tag/hash-table/)
- [String](https://leetcode.com/tag/string/)
- [Sorting](https://leetcode.com/tag/sorting/)

## Problem:

Given an array of strings `strs`, group **the anagrams** together. You can return the answer in **any order**.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

**Example 1:**

```
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

**Example 2:**

```
Input: strs = [""]
Output: [[""]]
```

**Example 3:**

```
Input: strs = ["a"]
Output: [["a"]]
```

**Constraints:**

- `1 <= strs.length <= 104`
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.

## Solution:

```go
func sortString(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)

	return strings.Join(chars, "")
}

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)

	memory := make(map[string][]string)

	for _, str := range strs {
		sorted := sortString(str)
		memory[sorted] = append(memory[sorted], str)
	}

	for _, values := range memory {
		result = append(result, values)
	}

	return result
}
```

## Similar Questions:

- [Valid Anagram](https://github.com/ju-popov/leetcode.com/tree/main/problems/valid-anagram/)
- [Group Shifted Strings](https://github.com/ju-popov/leetcode.com/tree/main/problems/group-shifted-strings/)
