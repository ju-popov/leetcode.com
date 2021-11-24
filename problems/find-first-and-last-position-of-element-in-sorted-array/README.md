# 34. Find First and Last Position of Element in Sorted Array

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Binary Search](https://leetcode.com/tag/binary-search/)

## Problem:

Given an array of integers `nums` sorted in non-decreasing order, find the starting and ending position of a given `target` value.

If `target` is not found in the array, return `[-1, -1]`.

You must write an algorithm with `O(log n)` runtime complexity.

**Example 1:**

```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```

**Example 2:**

```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```

**Example 3:**

```
Input: nums = [], target = 0
Output: [-1,-1]
```

**Constraints:**

- `0 <= nums.length <= 105`
- `-109 <= nums[i] <= 109`
- `nums` is a non-decreasing array.
- `-109 <= target <= 109`

## Solution:

```go
func binarySeach(nums []int, target int, first bool) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			if first {
				if (mid == left) || (nums[mid-1] != target) {
					return mid
				}

				right = mid - 1
			} else {
				if (mid == right) || (nums[mid+1] != target) {
					return mid
				}

				left = mid + 1
			}

			continue
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	first := binarySeach(nums, target, true)
	if first == -1 {
		return []int{-1, -1}
	}

	last := binarySeach(nums, target, false)

	return []int{first, last}
}
```

## Similar Questions:

- [First Bad Version](https://github.com/ju-popov/leetcode.com/tree/main/problems/first-bad-version/)
- [Plates Between Candles](https://github.com/ju-popov/leetcode.com/tree/main/problems/plates-between-candles/)
