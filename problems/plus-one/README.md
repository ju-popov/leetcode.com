# 66. Plus One

**Difficulty**: Easy

## Related Topics:

- [Array](https://leetcode.com/tag/array/)
- [Math](https://leetcode.com/tag/math/)

## Problem:

You are given a **large integer** represented as an integer array `digits`, where each `digits[i]` is the `ith` digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading `0`'s.

Increment the large integer by one and return *the resulting array of digits*.

**Example 1:**

```
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
```

**Example 2:**

```
Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
```

**Example 3:**

```
Input: digits = [0]
Output: [1]
Explanation: The array represents the integer 0.
Incrementing by one gives 0 + 1 = 1.
Thus, the result should be [1].
```

**Example 4:**

```
Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].
```

**Constraints:**

- `1 <= digits.length <= 100`
- `0 <= digits[i] <= 9`
- `digits` does not contain any leading `0`'s.

## Solution:

```go
func plusOne(digits []int) []int {
	carry := 1
	indexDigits := len(digits) - 1

	for (indexDigits >= 0) && (carry > 0) {
		sum := carry + digits[indexDigits]
		carry = sum / 10
		digits[indexDigits] = sum % 10
		indexDigits--
	}

	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}
```

## Similar Questions:

- [Multiply Strings](https://github.com/ju-popov/leetcode.com/tree/main/problems/multiply-strings/)
- [Add Binary](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-binary/)
- [Plus One Linked List](https://github.com/ju-popov/leetcode.com/tree/main/problems/plus-one-linked-list/)
- [Add to Array-Form of Integer](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-to-array-form-of-integer/)
