# 22. Generate Parentheses

**Difficulty**: Medium

## Related Topics:

- [String](https://leetcode.com/tag/string/)
- [Dynamic Programming](https://leetcode.com/tag/dynamic-programming/)
- [Backtracking](https://leetcode.com/tag/backtracking/)

## Problem:

Given `n` pairs of parentheses, write a function to *generate all combinations of well-formed parentheses*.

**Example 1:**

```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
```

**Example 2:**

```
Input: n = 1
Output: ["()"]
```

**Constraints:**

- `1 <= n <= 8`

## Solution:

```go
func generateParenthesis(n int) []string {
	results := []string{""}

	for index := 0; index < 2*n; index++ {
		newResults := make([]string, 0)

		for _, result := range results {
			opening := strings.Count(result, "(")
			closing := len(result) - opening

			if opening < n {
				newResults = append(newResults, result+"(")
			}

			if opening > closing {
				newResults = append(newResults, result+")")
			}
		}

		results = newResults
	}

	return results
}
```

## Similar Questions:

- [Letter Combinations of a Phone Number](https://github.com/ju-popov/leetcode.com/tree/main/problems/letter-combinations-of-a-phone-number/)
- [Valid Parentheses](https://github.com/ju-popov/leetcode.com/tree/main/problems/valid-parentheses/)
