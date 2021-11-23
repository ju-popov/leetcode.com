# 540. Single Element in a Sorted Array

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Binary Search](https://leetcode.com/tag/binary-search/)

## Problem:

You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.

Return *the single element that appears only once*.

Your solution must run in `O(log n)` time and `O(1)` space.

**Example 1:**

```
Input: nums = [1,1,2,3,3,4,4,8,8]
Output: 2
```

**Example 2:**

```
Input: nums = [3,3,7,7,10,11,11]
Output: 10
```

**Constraints:**

- `1 <= nums.length <= 105`
- `0 <= nums[i] <= 105`

## Solution:

```go
func singleNonDuplicate(nums []int) int {
	left := -1
	right := len(nums) / 2

	for {
		if right-left == 1 {
			return nums[right*2]
		}

		mid := left + (right-left)/2

		if nums[mid*2] == nums[mid*2+1] {
			left = mid
		} else {
			right = mid
		}
	}
}
```
