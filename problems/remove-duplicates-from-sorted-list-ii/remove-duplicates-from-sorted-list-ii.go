package main

import (
	"bytes"
	"fmt"
)

/*

82. Remove Duplicates from Sorted List II

https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

#linked-list #two-pointers

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

func main() {
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(deleteDuplicates(head1)) // 1 -> 2 -> 5

	head2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}

	fmt.Println(deleteDuplicates(head2)) // 2 -> 3

	head3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
			},
		},
	}

	fmt.Println(deleteDuplicates(head3)) // 1
}
