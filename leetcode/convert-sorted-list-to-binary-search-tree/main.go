package main

import "fmt"

func main() {
	l := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{Val: 5},
		},
	}

	r := sortedListToBST(l)

	fmt.Printf("%v <- %v -> %v\n", r.Val, r.Left.Val, r.Right.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var size int
	for e := head; e != nil; e = e.Next {
		size++
	}
	if size == 1 {
		return &TreeNode{Val: head.Val}
	}

	cursor := head
	var prev *ListNode

	for ix := 0; cursor != nil && ix < size/2; cursor, ix = cursor.Next, ix+1 {
		prev = cursor
	}
	prev.Next = nil

	return &TreeNode{
		Val:   cursor.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(cursor.Next),
	}
}
