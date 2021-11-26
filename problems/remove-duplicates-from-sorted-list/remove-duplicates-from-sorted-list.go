package main

import (
	"bytes"
	"fmt"
)

/*

83. Remove Duplicates from Sorted List

https://leetcode.com/problems/remove-duplicates-from-sorted-list/

#linked-list

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

func main() {
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
	}

	fmt.Println(deleteDuplicates(head1)) // 1 -> 2

	head2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}

	fmt.Println(deleteDuplicates(head2)) // 1 -> 2 -> 3
}
