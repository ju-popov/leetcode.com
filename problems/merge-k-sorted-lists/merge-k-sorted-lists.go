package main

import (
	"bytes"
	"fmt"
)

/*

23. Merge k Sorted Lists

https://leetcode.com/problems/merge-k-sorted-lists/

#linked-list #divide-and-conquer #heap-priority-queue #merge-sort

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	buf := bytes.NewBufferString("")

	for current := head; current != nil; current = current.Next {
		if current != head {
			if _, err := fmt.Fprint(buf, " -> "); err != nil {
				panic(err)
			}
		}

		if _, err := fmt.Fprintf(buf, "%v", current.Val); err != nil {
			panic(err)
		}
	}

	return buf.String()
}

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

func main() {
	lists := []*ListNode{
		{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
		{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		{
			Val: 2,
			Next: &ListNode{
				Val: 6,
			},
		},
	}

	fmt.Println(mergeKLists(lists))
}
