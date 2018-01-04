// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
//
// https://play.golang.org/p/tBsWsVL8fC8
//
package main

import "fmt"

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 5},
				},
			},
		},
	}
	l.print()

	tmp := removeNthFromEnd(l, 1)
	tmp.print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	for e := l; e != nil; e = e.Next {
		fmt.Print(e.Val, " -> ")
	}
	fmt.Println()
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n <= 0 {
		return head
	}
	cursor, prev := head, head
	var count int

	for ; cursor != nil; cursor = cursor.Next {
		if count > n {
			prev = prev.Next
		}
		count++
	}
	if count == n {
		return head.Next
	}

	if prev.Next != nil {
		prev.Next = prev.Next.Next
	} else {
		return nil
	}

	return head
}
