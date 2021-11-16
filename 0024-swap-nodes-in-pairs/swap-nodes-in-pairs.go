package main

import (
	"bytes"
	"fmt"
)

// https://leetcode.com/problems/swap-nodes-in-pairs/

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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// prev -> current1 -> current2 -> next
// prev -> current2 -> current1 -> next

func swapPairs(head *ListNode) *ListNode {
	current1 := head

	var prev *ListNode
	for (current1 != nil) && (current1.Next != nil) {
		current2 := current1.Next
		next := current2.Next

		if prev == nil {
			head = current2
		} else {
			prev.Next = current2
		}

		prev = current1

		current2.Next = current1
		current1.Next = next

		current1 = next
	}

	return head
}

func main() {
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}

	head3 := &ListNode{
		Val: 1,
	}

	fmt.Println(swapPairs(head1))
	fmt.Println(swapPairs(nil))
	fmt.Println(swapPairs(head3))
}
