# 31. Next Permutation

**Difficulty**: Medium

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Two Pointers](https://leetcode.com/tag/two-pointers/)

## Problem:

Implement **next permutation**, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).

The replacement must be **in place** and use only constant extra memory.

**Example 1:**

```
Input: nums = [1,2,3]
Output: [1,3,2]
```

**Example 2:**

```
Input: nums = [3,2,1]
Output: [1,2,3]
```

**Example 3:**

```
Input: nums = [1,1,5]
Output: [1,5,1]
```

**Example 4:**

```
Input: nums = [1]
Output: [1]
```

**Constraints:**

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 100`

## Solution:

```go
func nextPermutation(nums []int) []int {
	// Swap Left Value
	var swapLeftIndex int

	for swapLeftIndex = len(nums) - 2; swapLeftIndex >= 0; swapLeftIndex-- {
		if nums[swapLeftIndex] < nums[swapLeftIndex+1] {
			break
		}
	}

	// Swap Right Value
	if swapLeftIndex != -1 {
		swapRightIndex := -1

		for index := len(nums) - 1; index > swapLeftIndex; index-- {
			if nums[index] > nums[swapLeftIndex] {
				if (swapRightIndex == -1) || (nums[index] < nums[swapRightIndex]) {
					swapRightIndex = index
				}
			}
		}

		nums[swapLeftIndex], nums[swapRightIndex] = nums[swapRightIndex], nums[swapLeftIndex]
	}

	// Reverse the rest (sort ascending)
	start := swapLeftIndex + 1
	end := len(nums) - 1

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

	return nums
}
```

## Similar Questions:

- [Permutations](https://github.com/ju-popov/leetcode.com/tree/main/problems/permutations/)
- [Permutations II](https://github.com/ju-popov/leetcode.com/tree/main/problems/permutations-ii/)
- [Permutation Sequence](https://github.com/ju-popov/leetcode.com/tree/main/problems/permutation-sequence/)
- [Palindrome Permutation II](https://github.com/ju-popov/leetcode.com/tree/main/problems/palindrome-permutation-ii/)
- [Minimum Adjacent Swaps to Reach the Kth Smallest Number](https://github.com/ju-popov/leetcode.com/tree/main/problems/minimum-adjacent-swaps-to-reach-the-kth-smallest-number/)
