package main

import "fmt"

func main() {
	n := &ListNode{Val: 1}
	fmt.Println(*getIntersectionNode(n, n))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	inverse := map[*ListNode]*ListNode{headA: headA}

	for curA := headA; curA != nil; curA = curA.Next {
		if curA.Next != nil {
			inverse[curA.Next] = curA
		}
	}
	// fmt.Println(inverse)

	for curB := headB; curB != nil; curB = curB.Next {
		if _, found := inverse[curB]; found {
			return curB
		}
	}

	return nil
}
