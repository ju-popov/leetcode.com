# 75. Sort Colors

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Two Pointers](https://leetcode.com/tag/two-pointers/)
- [Sorting](https://leetcode.com/tag/sorting/)

## Problem:

Given an array `nums` with `n` objects colored red, white, or blue, sort them **in-place**so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers `0`, `1`, and `2` to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

**Example 1:**

```
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

**Example 2:**

```
Input: nums = [2,0,1]
Output: [0,1,2]
```

**Example 3:**

```
Input: nums = [0]
Output: [0]
```

**Example 4:**

```
Input: nums = [1]
Output: [1]
```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 300`
- `nums[i]` is `0`, `1`, or `2`.

**Follow up:** Could you come up with a one-pass algorithm using only constant extra space?

## Solution:

```go
func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	index := 0

	for index <= right {
		switch nums[index] {
		case 0:
			nums[left], nums[index] = nums[index], nums[left]
			left++
			index++
		case 2:
			nums[right], nums[index] = nums[index], nums[right]
			right--
		default:
			index++
		}
	}
}
```

## Similar Questions:

- [Sort List](https://github.com/ju-popov/leetcode.com/tree/main/problems/sort-list/)
- [Wiggle Sort](https://github.com/ju-popov/leetcode.com/tree/main/problems/wiggle-sort/)
- [Wiggle Sort II](https://github.com/ju-popov/leetcode.com/tree/main/problems/wiggle-sort-ii/)
