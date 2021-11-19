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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}

	var prev *ListNode

	current := head
	nodes := make([]*ListNode, k)

	for {
		// Load next nodes for reversal
		for i := 0; i < k; i++ {
			if current == nil {
				return head
			}

			nodes[i] = current
			current = current.Next
		}

		// Attach new nodes to previous
		if prev == nil {
			// Change head if needed
			head = nodes[k-1]
		} else {
			prev.Next = nodes[k-1]
		}

		// Reverse pointers
		for i := 0; i < k; i++ {
			if i == 0 {
				nodes[i].Next = current
			} else {
				nodes[i].Next = nodes[i-1]
			}
		}

		// Previous is now the last element
		prev = nodes[0]
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
