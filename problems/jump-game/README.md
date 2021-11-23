# 55. Jump Game

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Dynamic Programming](https://leetcode.com/tag/dynamic-programming/)
- [Greedy](https://leetcode.com/tag/greedy/)

## Problem:

You are given an integer array `nums`. You are initially positioned at the array's **first index**, and each element in the array represents your maximum jump length at that position.

Return `true`*if you can reach the last index, or*`false`*otherwise*.

**Example 1:**

```
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

**Example 2:**

```
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
```

**Constraints:**

- `1 <= nums.length <= 104`
- `0 <= nums[i] <= 105`

## Solution:

```go
func canJump(nums []int) bool {
	lastPos := len(nums) - 1

	for index := len(nums) - 1; index >= 0; index-- {
		if index+nums[index] >= lastPos {
			lastPos = index
		}
	}

	return lastPos == 0
}
```

## Similar Questions:

- [Jump Game II](https://leetcode.com/problems/jump-game-ii/)
- [Jump Game III](https://leetcode.com/problems/jump-game-iii/)
- [Jump Game VII](https://leetcode.com/problems/jump-game-vii/)
