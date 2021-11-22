# 18. 4Sum

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Two Pointers](https://leetcode.com/tag/two-pointers/)
- [Sorting](https://leetcode.com/tag/sorting/)

## Problem:

Given an array `nums` of `n` integers, return *an array of all the **unique** quadruplets* `[nums[a], nums[b], nums[c], nums[d]]` such that:

- `0 <= a, b, c, d < n`
- `a`, `b`, `c`, and `d` are **distinct**.
- `nums[a] + nums[b] + nums[c] + nums[d] == target`

You may return the answer in **any order**.

**Example 1:**

```
Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
```

**Example 2:**

```
Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]
```

**Constraints:**

- `1 <= nums.length <= 200`
- `-109 <= nums[i] <= 109`
- `-109 <= target <= 109`

## Solution:

```go
func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)

	sort.Ints(nums)

	for index1 := 0; index1 < len(nums)-3; index1++ {
		if (index1 != 0) && (nums[index1] == nums[index1-1]) {
			continue
		}

		for index2 := index1 + 1; index2 < len(nums)-2; index2++ {
			if (index2 != index1+1) && (nums[index2] == nums[index2-1]) {
				continue
			}

			left := index2 + 1
			right := len(nums) - 1

			for left < right {
				if (left != index2+1) && (nums[left] == nums[left-1]) {
					left++

					continue
				}

				sum := nums[index1] + nums[index2] + nums[left] + nums[right]

				if sum == target {
					result = append(result, []int{nums[index1], nums[index2], nums[left], nums[right]})
				}

				if sum > target {
					right--

					continue
				}

				left++
			}
		}
	}

	return result
}
```

## Similar Questions:

- [Two Sum](https://github.com/ju-popov/leetcode.com/tree/main/problems/two-sum/)
- [3Sum](https://github.com/ju-popov/leetcode.com/tree/main/problems/3sum/)
- [4Sum II](https://github.com/ju-popov/leetcode.com/tree/main/problems/4sum-ii/)
- [Count Special Quadruplets](https://github.com/ju-popov/leetcode.com/tree/main/problems/count-special-quadruplets/)
