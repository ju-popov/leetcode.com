# 15. 3Sum

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Two Pointers](https://leetcode.com/tag/two-pointers/)
- [Sorting](https://leetcode.com/tag/sorting/)

## Problem:

Given an integer array nums, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

**Example 1:**

```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
```

**Example 2:**

```
Input: nums = []
Output: []
```

**Example 3:**

```
Input: nums = [0]
Output: []
```

**Constraints:**

- `0 <= nums.length <= 3000`
- `-105 <= nums[i] <= 105`

## Solution:

```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	target := 0
	result := make([][]int, 0)

	if len(nums) < 3 {
		return result
	}

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

			sum := nums[index] + nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[index], nums[left], nums[right]})
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

- [Two Sum](https://github.com/ju-popov/leetcode.com/tree/main/problems/two-sum/)
- [3Sum Closest](https://github.com/ju-popov/leetcode.com/tree/main/problems/3sum-closest/)
- [4Sum](https://github.com/ju-popov/leetcode.com/tree/main/problems/4sum/)
- [3Sum Smaller](https://github.com/ju-popov/leetcode.com/tree/main/problems/3sum-smaller/)
