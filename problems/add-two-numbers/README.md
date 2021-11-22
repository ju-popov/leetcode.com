# 2. Add Two Numbers

**Difficulty**: Medium

## Related Topics:

-  [Linked List](https://leetcode.com/tag/linked-list/)
-  [Math](https://leetcode.com/tag/math/)
-  [Recursion](https://leetcode.com/tag/recursion/)

## Problem:

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order**, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example 1:**

```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
```

**Example 2:**

```
Input: l1 = [0], l2 = [0]
Output: [0]
```

**Example 3:**

```
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
```

**Constraints:**

- The number of nodes in each linked list is in the range `[1, 100]`.
- `0 <= Node.val <= 9`
- It is guaranteed that the list represents a number that does not have leading zeros.

## Solution:

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	current := dummyNode

	carry := 0

	for (l1 != nil) || (l2 != nil) || (carry != 0) {
		total := carry

		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{
			Val: total % 10,
		}

		carry = total / 10

		current = current.Next
	}

	return dummyNode.Next
}
```

## Similar Questions:

- [Multiply Strings](https://github.com/ju-popov/leetcode.com/tree/main/problems/multiply-strings/)
- [Add Binary](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-binary/)
- [Sum of Two Integers](https://github.com/ju-popov/leetcode.com/tree/main/problems/sum-of-two-integers/)
- [Add Strings](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-strings/)
- [Add Two Numbers II](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-two-numbers-ii/)
- [Add to Array-Form of Integer](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-to-array-form-of-integer/)
- [Add Two Polynomials Represented as Linked Lists](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-two-polynomials-represented-as-linked-lists/)
