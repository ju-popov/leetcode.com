# 53. Maximum Subarray

**Difficulty**: Easy

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Divide and Conquer](https://leetcode.com/tag/divide-and-conquer/)
- [Dynamic Programming](https://leetcode.com/tag/dynamic-programming/)

## Problem:

Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return *its sum*.

A **subarray** is a **contiguous** part of an array.

**Example 1:**

```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

**Example 2:**

```
Input: nums = [1]
Output: 1
```

**Example 3:**

```
Input: nums = [5,4,-1,7,8]
Output: 23
```

**Constraints:**

- `1 <= nums.length <= 105`
- `-104 <= nums[i] <= 104`

**Follow up:** If you have figured out the `O(n)` solution, try coding another solution using the **divide and conquer** approach, which is more subtle.

## Solution:

```go
func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := nums[0]

	for index := 1; index < len(nums); index++ {
		if currentSum < 0 {
			currentSum = 0
		}

		currentSum += nums[index]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
```

## Similar Questions:

- [Best Time to Buy and Sell Stock](https://github.com/ju-popov/leetcode.com/tree/main/problems/best-time-to-buy-and-sell-stock/)
- [Maximum Product Subarray](https://github.com/ju-popov/leetcode.com/tree/main/problems/maximum-product-subarray/)
- [Degree of an Array](https://github.com/ju-popov/leetcode.com/tree/main/problems/degree-of-an-array/)
- [Longest Turbulent Subarray](https://github.com/ju-popov/leetcode.com/tree/main/problems/longest-turbulent-subarray/)
- [Maximum Absolute Sum of Any Subarray](https://github.com/ju-popov/leetcode.com/tree/main/problems/maximum-absolute-sum-of-any-subarray/)
- [Maximum Subarray Sum After One Operation](https://github.com/ju-popov/leetcode.com/tree/main/problems/maximum-subarray-sum-after-one-operation/)
