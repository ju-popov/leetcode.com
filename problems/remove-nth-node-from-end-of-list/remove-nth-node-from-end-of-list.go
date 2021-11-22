package main

/*

19. Remove Nth Node From End of List

https://leetcode.com/problems/remove-nth-node-from-end-of-list/

#linked-list #two-pointers

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

func main() {
	inputList1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	fmt.Println(removeNthFromEnd(inputList1, 1)) // 1 -> 2 -> 3 -> 4
}
