# 704. Binary Search

**Difficulty**: Easy

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Binary Search](https://leetcode.com/tag/binary-search/)

## Problem:

Given an array of integers `nums` which is sorted in ascending order, and an integer `target`, write a function to search `target` in `nums`. If `target` exists, then return its index. Otherwise, return `-1`.

You must write an algorithm with `O(log n)` runtime complexity.

**Example 1:**

```
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4
```

**Example 2:**

```
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
```

**Constraints:**

- `1 <= nums.length <= 104`
- `-104 < nums[i], target < 104`
- All the integers in `nums` are **unique**.
- `nums` is sorted in ascending order.

## Solution:

```go
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		value := nums[mid]

		if value == target {
			return mid
		}

		if value < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
```

## Similar Questions:
  
- [Search in a Sorted Array of Unknown Size](https://github.com/ju-popov/leetcode.com/tree/main/problems/search-in-a-sorted-array-of-unknown-size/)
