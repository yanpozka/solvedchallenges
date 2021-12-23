package main

import "fmt"

func main() {
	h := &ListNode{Val: 2, Next: &ListNode{Val: 1}}
	h = removeElements(h, 2)
	fmt.Println(h.Val, h.Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		if cur.Val == val {
			if prev == nil {
				cur = cur.Next
				head = cur
			} else {
				prev.Next = cur.Next
				cur = cur.Next
			}
		} else {
			prev = cur
			cur = cur.Next
		}
	}

	if prev == nil {
		return nil
	}
	return head
}
