# 78. Subsets

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Backtracking](https://leetcode.com/tag/backtracking/)
- [Bit Manipulation](https://leetcode.com/tag/bit-manipulation/)

## Problem:

Given an integer array `nums` of **unique** elements, return *all possible subsets (the power set)*.

The solution set **must not** contain duplicate subsets. Return the solution in **any order**.

**Example 1:**

```
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
```

**Example 2:**

```
Input: nums = [0]
Output: [[],[0]]
```

**Constraints:**

- `1 <= nums.length <= 10`
- `-10 <= nums[i] <= 10`
- All the numbers of `nums` are **unique**.

## Solution:

```go
func subsets(nums []int) [][]int {
	results := make([][]int, 0)
	length := float64(len(nums))

	for i := 0; i < int(math.Pow(2, length)); i++ {
		result := make([]int, 0)

		for j, bitmask := 0, 1; j < len(nums); j++ {
			if (i & bitmask) != 0 {
				result = append(result, nums[j])
			}
			
			bitmask <<= 1
		}

		results = append(results, result)
	}

	return results
}
```

## Similar Questions:

- [Subsets II](https://github.com/ju-popov/leetcode.com/tree/main/problems/subsets-ii/)
- [Generalized Abbreviation](https://github.com/ju-popov/leetcode.com/tree/main/problems/generalized-abbreviation/)
- [Letter Case Permutation](https://github.com/ju-popov/leetcode.com/tree/main/problems/letter-case-permutation/)
- [Find Array Given Subset Sums](https://github.com/ju-popov/leetcode.com/tree/main/problems/find-array-given-subset-sums/)
- [Count Number of Maximum Bitwise-OR Subsets](https://github.com/ju-popov/leetcode.com/tree/main/problems/count-number-of-maximum-bitwise-or-subsets/)
