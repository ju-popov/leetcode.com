package main

import (
	"bytes"
	"fmt"
)

/*

2. Add Two Numbers

https://leetcode.com/problems/add-two-numbers/

Approach 1: Elementary Math

Complexity Analysis

Time complexity : O(max(m,n)). Assume that mm and nn represents the length of
l1 and l2 respectively, the algorithm above iterates at most max(m,n) times.

Space complexity : O(max(m,n)). The length of the new list is at most
max(m,n)+1.

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	current := dummyNode

	carry := 0

	for (l1 != nil) || (l2 != nil) || (carry != 0) {
		total := carry

		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{
			Val: total % 10,
		}

		carry = total / 10

		current = current.Next
	}

	return dummyNode.Next
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	fmt.Println(addTwoNumbers(l1, l2)) // 7 -> 0 -> 8
}
