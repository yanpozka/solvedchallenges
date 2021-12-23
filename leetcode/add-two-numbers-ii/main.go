package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	l1 := &ListNode{Val: 7, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	s := addTwoNumbers(l1, l2)
	fmt.Println(s.Val, s.Next.Val, s.Next.Next.Val, s.Next.Next.Next.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	buff := new(strings.Builder)
	cur := l1
	for cur != nil {
		buff.WriteString(strconv.Itoa(cur.Val))
		cur = cur.Next
	}
	a := new(big.Int)
	a.SetString(buff.String(), 10)

	buff.Reset()
	cur = l2
	for cur != nil {
		buff.WriteString(strconv.Itoa(cur.Val))
		cur = cur.Next
	}
	b := new(big.Int)
	b.SetString(buff.String(), 10)

	a.Add(a, b)
	res := a.String()

	head := &ListNode{Val: int(res[0] - '0')}
	cur = head
	for ix := 1; ix < len(res); ix++ {
		cur.Next = &ListNode{Val: int(res[ix] - '0')}
		cur = cur.Next
	}

	return head
}
