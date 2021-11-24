# 23. Merge k Sorted Lists


**Difficulty**: Hard

## Related Topics:

- [Linked List](https://leetcode.com/tag/linked-list/)
- [Divide and Conquer](https://leetcode.com/tag/divide-and-conquer/)
- [Heap (Priority Queue)](https://leetcode.com/tag/heap-priority-queue/)
- [Merge Sort](https://leetcode.com/tag/merge-sort/)

## Problem:

You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order.

*Merge all the linked-lists into one sorted linked-list and return it.*

**Example 1:**

```
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
```

**Example 2:**

```
Input: lists = []
Output: []
```

**Example 3:**

```
Input: lists = [[]]
Output: []
```

**Constraints:**

- `k == lists.length`
- `0 <= k <= 10^4`
- `0 <= lists[i].length <= 500`
- `-10^4 <= lists[i][j] <= 10^4`
- `lists[i]` is sorted in **ascending order**.
- The sum of `lists[i].length` won't exceed `10^4`.

## Solution:

```go
func nextNode(lists []*ListNode) (*ListNode, int) {
	var minListNode *ListNode

	minIndex := 0

	for index, list := range lists {
		if list != nil {
			if (minListNode == nil) || (list.Val < minListNode.Val) {
				minListNode = list
				minIndex = index
			}
		}
	}

	return minListNode, minIndex
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummyNode := &ListNode{}
	last := dummyNode

	for {
		next, index := nextNode(lists)
		if next == nil {
			break
		}

		last.Next = next
		last = last.Next
		lists[index] = lists[index].Next
	}

	return dummyNode.Next
}
```

## Similar Questions:
  
- [Merge Two Sorted Lists](https://github.com/ju-popov/leetcode.com/tree/main/problems/merge-two-sorted-lists/)
- [Ugly Number II](https://github.com/ju-popov/leetcode.com/tree/main/problems/ugly-number-ii/)

