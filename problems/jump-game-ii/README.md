# 45. Jump Game II

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Dynamic Programming](https://leetcode.com/tag/dynamic-programming/)
- [Greedy](https://leetcode.com/tag/greedy/)

## Problem:

Given an array of non-negative integers `nums`, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

You can assume that you can always reach the last index.

**Example 1:**

```
Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

**Example 2:**

```
Input: nums = [2,3,0,1,4]
Output: 2
```

**Constraints:**

- `1 <= nums.length <= 104`
- `0 <= nums[i] <= 1000`

## Solution:

```go
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	pos := 0
	result := 0

	for {
		if nums[pos] == 0 {
			return -1
		}

		totalMaxRich := 0
		nextPos := 0

		for step := 1; step <= nums[pos]; step++ {
			if pos+step == len(nums)-1 {
				return result + 1
			}

			maxRich := nums[pos+step] + step

			if maxRich > totalMaxRich {
				totalMaxRich = maxRich
				nextPos = pos + step
			}
		}

		pos = nextPos
		result++
	}
}
```

## Similar Questions:

- [Jump Game](https://github.com/ju-popov/leetcode.com/tree/main/problems/jump-game/)
- [Jump Game III](https://github.com/ju-popov/leetcode.com/tree/main/problems/jump-game-iii/)
- [Jump Game VII](https://github.com/ju-popov/leetcode.com/tree/main/problems/jump-game-vii/)
