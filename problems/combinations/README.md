# 77. Combinations

**Difficulty**: Medium

## Related Topics:
- [Array](https://leetcode.com/tag/array/)
- [Backtracking](https://leetcode.com/tag/backtracking/)

## Problem:

Given two integers `n` and `k`, return *all possible combinations of* `k` *numbers out of the range* `[1, n]`.

You may return the answer in **any order**.

**Example 1:**

```
Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
```

**Example 2:**

```
Input: n = 1, k = 1
Output: [[1]]
```

**Constraints:**

- `1 <= n <= 20`
- `1 <= k <= n`

## Solution:

```go
func combine(n int, k int) [][]int {
	results := make([][]int, 0)

	if k == 1 {
		for i := 1; i <= n; i++ {
			results = append(results, []int{i})
		}

		return results
	}

	for _, prevResult := range combine(n-1, k-1) {
		for i := prevResult[len(prevResult)-1] + 1; i <= n; i++ {
			result := make([]int, 0)
			result = append(result, prevResult...)
			result = append(result, i)
			results = append(results, result)
		}
	}

	return results
}
```

## Similar Questions:

- [Combination Sum](https://github.com/ju-popov/leetcode.com/tree/main/problems/combination-sum/)
- [Permutations](https://github.com/ju-popov/leetcode.com/tree/main/problems/permutations/)
