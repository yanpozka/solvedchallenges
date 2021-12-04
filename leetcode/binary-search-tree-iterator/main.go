package main

import "fmt"

func main() {
	bstit := Constructor(&TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	})
	fmt.Println(bstit.HasNext())
	fmt.Println(bstit.Next())
	fmt.Println(bstit.HasNext())
	fmt.Println(bstit.Next())
	fmt.Println(bstit.HasNext())
	fmt.Println(bstit.Next())
	fmt.Println(bstit.HasNext())
	fmt.Println(bstit.Next())
	fmt.Println(bstit.HasNext())
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	nodes []int
	ix    int
}

func Constructor(root *TreeNode) BSTIterator {
	bstit := BSTIterator{}
	bstit.inOrder(root)
	return bstit
}

func (this *BSTIterator) inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	this.inOrder(root.Left)
	this.nodes = append(this.nodes, root.Val)
	this.inOrder(root.Right)
}

func (this *BSTIterator) Next() int {
	if this.ix < len(this.nodes) {
		n := this.nodes[this.ix]
		this.ix++
		return n
	}
	return -1
}

func (this *BSTIterator) HasNext() bool {
	return this.ix < len(this.nodes)
}
