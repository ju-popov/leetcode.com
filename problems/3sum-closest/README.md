# 16. 3Sum Closest

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Two Pointers](https://leetcode.com/tag/two-pointers/)
- [Sorting](https://leetcode.com/tag/sorting/)

## Problem:

Given an integer array `nums` of length `n` and an integer `target`, find three integers in `nums` such that the sum is closest to `target`.

Return *the sum of the three integers*.

You may assume that each input would have exactly one solution.

**Example 1:**

```
Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
```

**Example 2:**

```
Input: nums = [0,0,0], target = 1
Output: 0
```

**Constraints:**

- `3 <= nums.length <= 1000`
- `-1000 <= nums[i] <= 1000`
- `-104 <= target <= 104`

## Solution:

```go
func calcResult(n1 int, n2 int, n3 int, target int) (int, int) {
	total := n1 + n2 + n3

	offset := target - total
	if offset < 0 {
		offset = -offset
	}

	return total, offset
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	result, resultOffset := calcResult(nums[0], nums[1], nums[2], target)

	for index := 0; index < len(nums)-2; index++ {
		if (index != 0) && (nums[index] == nums[index-1]) {
			continue
		}

		left := index + 1
		right := len(nums) - 1

		for left < right {
			if (left != index+1) && (nums[left] == nums[left-1]) {
				left++

				continue
			}

			sum, sumOffset := calcResult(nums[index], nums[left], nums[right], target)

			if sumOffset == 0 {
				return sum
			}

			if sumOffset < resultOffset {
				result, resultOffset = sum, sumOffset
			}

			if sum > target {
				right--

				continue
			}

			left++
		}
	}

	return result
}
```

## Similar Questions:

- [3Sum](https://github.com/ju-popov/leetcode.com/tree/main/problems/3sum/)
- [3Sum Smaller](https://github.com/ju-popov/leetcode.com/tree/main/problems/3sum-smaller/)
