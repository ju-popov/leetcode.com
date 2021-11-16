package main

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	frontNode := head

	for counter := 0; counter < n; counter++ {
		if frontNode == nil {
			return head
		}

		frontNode = frontNode.Next
	}

	var prevNode *ListNode
	currentNode := head

	for frontNode != nil {
		frontNode = frontNode.Next
		prevNode = currentNode
		currentNode = currentNode.Next
	}

	if prevNode == nil {
		return currentNode.Next
	}

	prevNode.Next = currentNode.Next

	return head
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

	fmt.Println(removeNthFromEnd(inputList1, 1))
}
