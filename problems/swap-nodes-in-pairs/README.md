# 24. Swap Nodes in Pairs

**Difficulty**: Medium

## Related Topics:

- [Linked List](https://leetcode.com/tag/linked-list/)
- [Recursion](https://leetcode.com/tag/recursion/)

## Problem:

Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

**Example 1:**

```
Input: head = [1,2,3,4]
Output: [2,1,4,3]
```

**Example 2:**

```
Input: head = []
Output: []
```

**Example 3:**

```
Input: head = [1]
Output: [1]
```

**Constraints:**

- The number of nodes in the list is in the range `[0, 100]`.
- `0 <= Node.val <= 100`

## Solution:

```go
func swapPairs(head *ListNode) *ListNode {
	const k = 2 // Whap is pairs

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

- [Reverse Nodes in k-Group](https://github.com/ju-popov/leetcode.com/tree/main/problems/reverse-nodes-in-k-group/)
- [Swapping Nodes in a Linked List](https://github.com/ju-popov/leetcode.com/tree/main/problems/swapping-nodes-in-a-linked-list/)
