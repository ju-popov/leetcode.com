# 19. Remove Nth Node From End of List

**Difficulty**: Medium

## Related Topics:

- [Linked List](https://leetcode.com/tag/linked-list/)
- [Two Pointers](https://leetcode.com/tag/two-pointers/)

## Problem:

Given the `head` of a linked list, remove the `nth` node from the end of the list and return its head.

**Example 1:**

```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```

**Example 2:**

```
Input: head = [1], n = 1
Output: []
```

**Example 3:**

```
Input: head = [1,2], n = 1
Output: [1]
```

**Constraints:**

- The number of nodes in the list is `sz`.
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

**Follow up:** Could you do this in one pass?

## Solution:

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first := dummy
	second := dummy

	for counter := 0; counter < n+1; counter++ {
		if first == nil {
			return dummy.Next
		}

		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}
```

## Similar Questions:

- [Swapping Nodes in a Linked List](https://github.com/ju-popov/leetcode.com/tree/main/problems/swapping-nodes-in-a-linked-list/)
- [Delete N Nodes After M Nodes of a Linked List](https://github.com/ju-popov/leetcode.com/tree/main/problems/delete-n-nodes-after-m-nodes-of-a-linked-list/)
