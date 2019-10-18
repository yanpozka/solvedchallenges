package main

import "fmt"

func main() {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}
	d := deleteDuplicates(l)
	fmt.Printf("%d -> %v\n", d.Val, d.Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	for curr != nil {
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		}
		curr = curr.Next
	}

	return head
}
