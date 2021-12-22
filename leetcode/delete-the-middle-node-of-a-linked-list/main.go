package main

import "fmt"

func main() {
	{
		head := &ListNode{
			Val:  1,
			Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		}
		head = deleteMiddle(head)
		fmt.Println(head.Val, head.Next.Val, head.Next.Next.Val, head.Next.Next.Next)
	}
	{
		head := &ListNode{Val: 2, Next: &ListNode{Val: 1}}
		head = deleteMiddle(head)
		fmt.Println(head.Val, head.Next)
	}
	{
		head := &ListNode{Val: 2, Next: nil}
		head = deleteMiddle(head)
		fmt.Println(head)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	var size int

	for cur := head; cur != nil; cur = cur.Next {
		size++
	}
	if size == 1 {
		return nil
	}

	tmp := head
	for i, h := 0, size/2-1; i < h; i++ {
		tmp = tmp.Next
	}
	if tmp.Next != nil {
		tmp.Next = tmp.Next.Next
	}

	return head
}
