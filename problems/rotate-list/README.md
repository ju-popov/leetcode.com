# 61. Rotate List

**Difficulty**: Medium

## Related Topics:

- [Linked List](https://leetcode.com/tag/linked-list/)
- [Two Pointers](https://leetcode.com/tag/two-pointers/)

## Problem:

Given the `head` of a linked list, rotate the list to the right by `k` places.

**Example 1:**

```
Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]
```

**Example 2:**

```
Input: head = [0,1,2], k = 4
Output: [2,0,1]
```

**Constraints:**

- The number of nodes in the list is in the range `[0, 500]`.
- `-100 <= Node.val <= 100`
- `0 <= k <= 2 * 109`

## Solution:

```go
func countNodes(head *ListNode) (int, *ListNode) {
	count := 0

	var prev *ListNode
	for current := head; current != nil; current = current.Next {
		prev = current
		count++
	}

	return count, prev
}

func rotateRight(head *ListNode, k int) *ListNode {
	nodesCount, tail := countNodes(head)
	if nodesCount <= 1 {
		return head
	}

	// Discard full loops count
	k -= k / nodesCount * nodesCount
	if k == 0 {
		return head
	}

	newHead := head
	newTail := (*ListNode)(nil)

	// Go forward (nodesCount - k) nodes
	for index := 0; index < nodesCount-k; index++ {
		newTail = newHead
		newHead = newHead.Next
	}

	newTail.Next = nil
	tail.Next = head

	return newHead
}
```

## Similar Questions:

- [Rotate Array](https://github.com/ju-popov/leetcode.com/tree/main/problems/rotate-array/)
- [Split Linked List in Parts](https://github.com/ju-popov/leetcode.com/tree/main/problems/split-linked-list-in-parts/)
