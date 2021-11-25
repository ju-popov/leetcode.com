# 83. Remove Duplicates from Sorted List

**Difficulty**: Easy

## Related Topics:

- [Linked List](https://leetcode.com/tag/linked-list/)

## Problem:

Given the `head` of a sorted linked list, *delete all duplicates such that each element appears only once*. Return *the linked list **sorted** as well*.

**Example 1:**

```
Input: head = [1,1,2]
Output: [1,2]
```

**Example 2:**

```
Input: head = [1,1,2,3,3]
Output: [1,2,3]
```

**Constraints:**

- The number of nodes in the list is in the range `[0, 300]`.
- `-100 <= Node.val <= 100`
- The list is guaranteed to be **sorted** in ascending order.

## Solution:

```go
func deleteDuplicates(head *ListNode) *ListNode {
	if (head == nil) || (head.Next == nil) {
		return head
	}

	dummyNode := &ListNode{}
	last := dummyNode

	prev := (*ListNode)(nil)
	for current := head; current != nil; current = current.Next {
		if (prev == nil) || (prev.Val != current.Val) {
			last.Next = current
			last = current
		}

		prev = current
	}

	last.Next = nil

	return dummyNode.Next
}
```

## Similar Questions:

- [Remove Duplicates from Sorted List II](https://github.com/ju-popov/leetcode.com/tree/main/problems/remove-duplicates-from-sorted-list-ii/)
- [Remove Duplicates From an Unsorted Linked List](https://github.com/ju-popov/leetcode.com/tree/main/problems/remove-duplicates-from-an-unsorted-linked-list/)
