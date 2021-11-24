# 40. Combination Sum II

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Backtracking](https://leetcode.com/tag/backtracking/)

## Problem:

Given a collection of candidate numbers (`candidates`) and a target number (`target`), find all unique combinations in `candidates` where the candidate numbers sum to `target`.

Each number in `candidates` may only be used **once** in the combination.

**Note:** The solution set must not contain duplicate combinations.

**Example 1:**

```
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output: 
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
```

**Example 2:**

```
Input: candidates = [2,5,2,1,2], target = 5
Output: 
[
[1,2,2],
[5]
]
```

**Constraints:**

- `1 <= candidates.length <= 100`
- `1 <= candidates[i] <= 50`
- `1 <= target <= 30`

## Solution:

```go
func combinationSum2Helper(candidates []int, target int, candidateIndex int) [][]int {
	result := make([][]int, 0)

	if candidateIndex >= len(candidates) {
		return result
	}

	candidate := candidates[candidateIndex]

	if candidate == target {
		result = append(result, []int{candidate})
	}

	// With current candidate
	if target > candidate {
		for _, subSolution := range combinationSum2Helper(candidates, target-candidate, candidateIndex+1) {
			solution := append([]int{candidate}, subSolution...)
			result = append(result, solution)
		}
	}

	// Without current candidate - Skip diplicate candidates
	for index := candidateIndex + 1; index < len(candidates); index++ {
		if candidates[index] != candidates[index-1] {
			result = append(result, combinationSum2Helper(candidates, target, index)...)

			break
		}
	}

	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	return combinationSum2Helper(candidates, target, 0)
}
```

## Similar Questions:

- [Combination Sum](https://github.com/ju-popov/leetcode.com/tree/main/problems/combination-sum/)
