# 82. Remove Duplicates from Sorted List II

**Difficulty**: Medium

## Related Topics:

- [Linked List](https://leetcode.com/tag/linked-list/)
- [Two Pointers](https://leetcode.com/tag/two-pointers/)

## Problem:

Given the `head` of a sorted linked list, *delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list*. Return *the linked list **sorted** as well*.

**Example 1:**

```
Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]
```

**Example 2:**

```
Input: head = [1,1,1,2,3]
Output: [2,3]
```

**Constraints:**

- The number of nodes in the list is in the range `[0, 300]`.
- `-100 <= Node.val <= 100`
- The list is guaranteed to be **sorted** in ascending order.

## Solution:

```go
func nextUniqueNode(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head.Next
	}

	if head.Val != head.Next.Val {
		return head, head.Next
	}

	prev := (*ListNode)(nil)
	for (head != nil) && ((prev == nil) || (head.Val == prev.Val)) {
		prev = head
		head = head.Next
	}

	return nil, head
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyNode := &ListNode{}
	lastNode := dummyNode

	for current := head; current != nil; {
		var node *ListNode
		node, current = nextUniqueNode(current)

		if node != nil {
			lastNode.Next = node
			lastNode = node
			lastNode.Next = nil
		}
	}

	return dummyNode.Next
}
```

## Similar Questions:

- [Remove Duplicates from Sorted List](https://github.com/ju-popov/leetcode.com/tree/main/problems/remove-duplicates-from-sorted-list/)
- [Remove Duplicates From an Unsorted Linked List](https://github.com/ju-popov/leetcode.com/tree/main/problems/remove-duplicates-from-an-unsorted-linked-list/)
