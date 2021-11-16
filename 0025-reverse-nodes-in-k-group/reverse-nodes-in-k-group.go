package main

import (
	"bytes"
	"fmt"
)

// https://leetcode.com/problems/reverse-nodes-in-k-group/

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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}

	current := head
	nodes := make([]*ListNode, k+1)

	for {
		for i := 1; i <= k; i++ {
			if current == nil {
				return head
			}

			nodes[i] = current
			current = current.Next
		}

		if nodes[0] == nil {
			head = nodes[k]
		} else {
			nodes[0].Next = nodes[k]
		}

		nodes[0] = nodes[1]
		nodes[1].Next = current

		for i := 2; i <= k; i++ {
			nodes[i].Next = nodes[i-1]
		}
	}
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
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	head2 := &ListNode{
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

	head3 := &ListNode{
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

	head4 := &ListNode{
		Val: 1,
	}

	fmt.Println(reverseKGroup(head1, 2))
	fmt.Println(reverseKGroup(head2, 3))
	fmt.Println(reverseKGroup(head3, 1))
	fmt.Println(reverseKGroup(head4, 1))
}
