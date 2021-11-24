# 47. Permutations II

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Backtracking](https://leetcode.com/tag/backtracking/)

## Problem:

Given a collection of numbers, `nums`, that might contain duplicates, return *all possible unique permutations **in any order**.*

**Example 1:**

```
Input: nums = [1,1,2]
Output:
[[1,1,2],
 [1,2,1],
 [2,1,1]]
```

**Example 2:**

```
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

**Constraints:**

- `1 <= nums.length <= 8`
- `-10 <= nums[i] <= 10`

## Solution:

```go
func permuteUniqueHelper(nums []int) [][]int {
	results := make([][]int, 0)

	if len(nums) == 0 {
		results = append(results, []int{})

		return results
	}

	for index := 0; index < len(nums); index++ {
		// Skip duplicate values (in sorted array)
		if index > 0 && (nums[index] == nums[index-1]) {
			continue
		}

		// Remove value at index position
		newNums := append([]int{}, nums[:index]...)
		newNums = append(newNums, nums[index+1:]...)

		for _, subResult := range permuteUniqueHelper(newNums) {
			// value at index position become first
			result := append([]int{nums[index]}, subResult...)
			results = append(results, result)
		}
	}

	return results
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	return permuteUniqueHelper(nums)
}
```

## Similar Questions:

- [Next Permutation](https://github.com/ju-popov/leetcode.com/tree/main/problems/next-permutation/)
- [Permutations](https://github.com/ju-popov/leetcode.com/tree/main/problems/permutations/)
- [Palindrome Permutation II](https://github.com/ju-popov/leetcode.com/tree/main/problems/palindrome-permutation-ii/)
- [Number of Squareful Arrays](https://github.com/ju-popov/leetcode.com/tree/main/problems/number-of-squareful-arrays/)
