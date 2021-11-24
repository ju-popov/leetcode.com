# 21. Merge Two Sorted Lists

**Difficulty**: Easy

## Related Topics:

- [Linked List](https://leetcode.com/tag/linked-list/)
- [Recursion](https://leetcode.com/tag/recursion/)

## Problem:

Merge two sorted linked lists and return it as a **sorted** list. The list should be made by splicing together the nodes of the first two lists.

**Example 1:**

```
Input: l1 = [1,2,4], l2 = [1,3,4]
Output: [1,1,2,3,4,4]
```

**Example 2:**

```
Input: l1 = [], l2 = []
Output: []
```

**Example 3:**

```
Input: l1 = [], l2 = [0]
Output: [0]
```

**Constraints:**

- The number of nodes in both lists is in the range `[0, 50]`.
- `-100 <= Node.val <= 100`
- Both `l1` and `l2` are sorted in **non-decreasing** order.

## Solution:

```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	last := dummyNode

	for (l1 != nil) && (l2 != nil) {
		if l1.Val <= l2.Val {
			last.Next = l1
			last = l1
			l1 = l1.Next
		} else {
			last.Next = l2
			last = l2
			l2 = l2.Next
		}
	}

	if l1 != nil {
		last.Next = l1
	} else if l2 != nil {
		last.Next = l2
	}

	return dummyNode.Next
}
```

## Similar Questions:

- [Merge k Sorted Lists](https://github.com/ju-popov/leetcode.com/tree/main/problems/merge-k-sorted-lists/)
- [Merge Sorted Array](https://github.com/ju-popov/leetcode.com/tree/main/problems/merge-sorted-array/)
- [Sort List](https://github.com/ju-popov/leetcode.com/tree/main/problems/sort-list/)
- [Shortest Word Distance II](https://github.com/ju-popov/leetcode.com/tree/main/problems/shortest-word-distance-ii/)
- [Add Two Polynomials Represented as Linked Lists](https://github.com/ju-popov/leetcode.com/tree/main/problems/add-two-polynomials-represented-as-linked-lists/)
- [Longest Common Subsequence Between Sorted Arrays](https://github.com/ju-popov/leetcode.com/tree/main/problems/longest-common-subsequence-between-sorted-arrays/)
