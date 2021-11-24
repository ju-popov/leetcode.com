# 25. Reverse Nodes in k-Group

**Difficulty**: Hard

## Related Topics:

- [Linked List](https://leetcode.com/tag/linked-list/)
- [Recursion](https://leetcode.com/tag/recursion/)

## Problem:

Given a linked list, reverse the nodes of a linked list *k* at a time and return its modified list.

*k* is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of *k* then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

**Example 1:**

```
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
```

**Example 2:**

```
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]
```

**Example 3:**

```
Input: head = [1,2,3,4,5], k = 1
Output: [1,2,3,4,5]
```

**Example 4:**

```
Input: head = [1], k = 1
Output: [1]
```

**Constraints:**

- The number of nodes in the list is in the range `sz`.
- `1 <= sz <= 5000`
- `0 <= Node.val <= 1000`
- `1 <= k <= sz`

**Follow-up:** Can you solve the problem in O(1) extra memory space?

## Solution:

```go
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}

	var prev *ListNode

	current := head
	nodes := make([]*ListNode, k)

	for {
		// Load next nodes for reversal
		for i := 0; i < k; i++ {
			if current == nil {
				return head
			}

			nodes[i] = current
			current = current.Next
		}

		// Attach new nodes to previous
		if prev == nil {
			// Change head if needed
			head = nodes[k-1]
		} else {
			prev.Next = nodes[k-1]
		}

		// Reverse pointers
		for i := 0; i < k; i++ {
			if i == 0 {
				nodes[i].Next = current
			} else {
				nodes[i].Next = nodes[i-1]
			}
		}

		// Previous is now the last element
		prev = nodes[0]
	}
}
```

## Similar Questions:

- [Swap Nodes in Pairs](https://github.com/ju-popov/leetcode.com/tree/main/problems/swap-nodes-in-pairs/)
- [Swapping Nodes in a Linked List](https://github.com/ju-popov/leetcode.com/tree/main/problems/swapping-nodes-in-a-linked-list/)
- [Reverse Nodes in Even Length Groups](https://github.com/ju-popov/leetcode.com/tree/main/problems/reverse-nodes-in-even-length-groups/)
