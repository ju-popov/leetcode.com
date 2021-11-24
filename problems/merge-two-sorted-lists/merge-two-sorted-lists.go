package main

/*

21. Merge Two Sorted Lists

https://leetcode.com/problems/merge-two-sorted-lists/

#linked-list #recursion

*/

import (
	"bytes"
	"fmt"
)

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

func main() {
	inputList1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	inputList2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	fmt.Println(mergeTwoLists(inputList1, inputList2)) // 1 -> 1 -> 2 -> 3 -> 4 -> 4
}
