# 448. Find All Numbers Disappeared in an Array


**Difficulty**: Easy

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Hash Table](https://leetcode.com/tag/hash-table/)

## Problem:

Given an array `nums` of `n` integers where `nums[i]` is in the range `[1, n]`, return *an array of all the integers in the range* `[1, n]` *that do not appear in* `nums`.

**Example 1:**

```
Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
```

**Example 2:**

```
Input: nums = [1,1]
Output: [2]
```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 105`
- `1 <= nums[i] <= n`

**Follow up:** Could you do it without extra space and in `O(n)` runtime? You may assume the returned list does not count as extra space.

## Solution:

```go
func findDisappearedNumbers(nums []int) []int {
	memory := make([]bool, len(nums))

	for _, num := range nums {
		memory[num-1] = true
	}

	result := make([]int, 0)

	for index := 0; index < len(nums); index++ {
		if !memory[index] {
			result = append(result, index+1)
		}
	}

	return result
}
```

## Similar Questions:

- [First Missing Positive](https://github.com/ju-popov/leetcode.com/tree/main/problems/first-missing-positive/)
- [Find All Duplicates in an Array](https://github.com/ju-popov/leetcode.com/tree/main/problems/find-all-duplicates-in-an-array/)
- [Find Unique Binary String](https://github.com/ju-popov/leetcode.com/tree/main/problems/find-unique-binary-string/)
