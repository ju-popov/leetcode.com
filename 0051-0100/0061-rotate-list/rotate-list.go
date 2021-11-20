package main

import (
	"bytes"
	"fmt"
)

/*

61. Rotate List

https://leetcode.com/problems/rotate-list/

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

func countNodes(head *ListNode) (int, *ListNode) {
	count := 0

	var prev *ListNode
	for current := head; current != nil; current = current.Next {
		prev = current
		count++
	}

	return count, prev
}

func rotateRight(head *ListNode, k int) *ListNode {
	nodesCount, tail := countNodes(head)
	if nodesCount <= 1 {
		return head
	}

	// Discard full loops count
	k -= k / nodesCount * nodesCount
	if k == 0 {
		return head
	}

	newHead := head
	newTail := (*ListNode)(nil)

	// Go forward (nodesCount - k) nodes
	for index := 0; index < nodesCount-k; index++ {
		newTail = newHead
		newHead = newHead.Next
	}

	newTail.Next = nil
	tail.Next = head

	return newHead
}

func main() {
	list1 := &ListNode{
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

	list2 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
	}

	fmt.Println(rotateRight(list1, 2)) // 4 -> 5 -> 1 -> 2 -> 3
	fmt.Println(rotateRight(list2, 4)) // 2 -> 0 -> 1
}
