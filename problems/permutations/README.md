# 46. Permutations

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Backtracking](https://leetcode.com/tag/backtracking/)

## Problem:

Given an array `nums` of distinct integers, return *all the possible permutations*. You can return the answer in **any order**.

**Example 1:**

```
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

**Example 2:**

```
Input: nums = [0,1]
Output: [[0,1],[1,0]]
```

**Example 3:**

```
Input: nums = [1]
Output: [[1]]
```

**Constraints:**

- `1 <= nums.length <= 6`
- `-10 <= nums[i] <= 10`
- All the integers of `nums` are **unique**.

## Solution:

```go
func permute(nums []int) [][]int {
	results := make([][]int, 0)

	if len(nums) == 0 {
		results = append(results, []int{})

		return results
	}

	for index := 0; index < len(nums); index++ {
		removedValue := nums[index]

		newNums := append([]int{}, nums[:index]...)
		newNums = append(newNums, nums[index+1:]...)

		subResults := permute(newNums)

		for _, subResult := range subResults {
			result := append([]int{removedValue}, subResult...)
			results = append(results, result)
		}
	}

	return results
}
```

## Similar Questions:

- [Next Permutation](https://github.com/ju-popov/leetcode.com/tree/main/problems/next-permutation/)
- [Permutations II](https://github.com/ju-popov/leetcode.com/tree/main/problems/permutations-ii/)
- [Permutation Sequence](https://github.com/ju-popov/leetcode.com/tree/main/problems/permutation-sequence/)
- [Combinations](https://github.com/ju-popov/leetcode.com/tree/main/problems/combinations/)
