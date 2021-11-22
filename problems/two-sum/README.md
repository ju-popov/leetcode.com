# 1. Two Sum

**Difficulty**: Easy

## Related Topics:

* [Array](https://leetcode.com/tag/array/)
* [Hash Table](https://leetcode.com/tag/hash-table/)

## Problem:

Given an array of integers `nums` and an integer `target`, return *indices of the two numbers such that they add up to target*.

You may assume that each input would have **exactly one solution**, and you may not use the *same* element twice.

You can return the answer in any order.

**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

**Constraints:**

- `2 <= nums.length <= 104`
- `-109 <= nums[i] <= 109`
- `-109 <= target <= 109`
- **Only one valid answer exists.**

**Follow-up:**Can you come up with an algorithm that is less than `O(n2)`time complexity?

## Solution:

```go
func twoSum(nums []int, target int) []int {
	visited := make(map[int]int)

	for index, value := range nums {
		if _, ok := visited[target-value]; ok {
			return []int{visited[target-value], index}
		}

		visited[value] = index
	}

	return []int{-1, -1}
}
```

## Similar Questions:

- [3Sum](https://github.com/ju-popov/leetcode.com/tree/main/problems/3sum/)
- [4Sum](https://github.com/ju-popov/leetcode.com/tree/main/problems/4sum/)
- [Two Sum II - Input Array Is Sorted](https://github.com/ju-popov/leetcode.com/tree/main/problems/two-sum-ii-input-array-is-sorted/)
- [Two Sum III - Data structure design](https://github.com/ju-popov/leetcode.com/tree/main/problems/two-sum-iii-data-structure-design/)
- [Subarray Sum Equals K](https://github.com/ju-popov/leetcode.com/tree/main/problems/subarray-sum-equals-k/)
- [Two Sum IV - Input is a BST](https://github.com/ju-popov/leetcode.com/tree/main/problems/two-sum-iv-input-is-a-bst/)
- [Two Sum Less Than K](https://github.com/ju-popov/leetcode.com/tree/main/problems/two-sum-less-than-k/)
- [Max Number of K-Sum Pairs](https://github.com/ju-popov/leetcode.com/tree/main/problems/max-number-of-k-sum-pairs/)
- [Count Good Meals](https://github.com/ju-popov/leetcode.com/tree/main/problems/count-good-meals/)
- [Count Number of Pairs With Absolute Difference K](https://github.com/ju-popov/leetcode.com/tree/main/problems/count-number-of-pairs-with-absolute-difference-k/)
- [Number of Pairs of Strings With Concatenation Equal to Target](https://github.com/ju-popov/leetcode.com/tree/main/problems/number-of-pairs-of-strings-with-concatenation-equal-to-target/)
