package sbt

import (
	"strconv"
	"strings"
)

const (
	emptyNode = 9999
	sep       = "/"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	nodes []int
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	this.nodes = make([]int, 0, 10001)

	this.fillLikeHeap(root, 0)

	strs := make([]string, len(this.nodes))
	for ix, n := range this.nodes {
		strs[ix] = strconv.Itoa(n)
	}

	return strings.Join(strs, sep)
}

func (this *Codec) fillLikeHeap(root *TreeNode, index int) {
	diff := index - len(this.nodes)
	for ; diff >= 0; diff-- {
		this.nodes = append(this.nodes, emptyNode)
	}
	if root == nil {
		this.nodes[index] = emptyNode
		return
	}

	this.nodes[index] = root.Val

	this.fillLikeHeap(root.Left, 2*index+1)
	this.fillLikeHeap(root.Right, 2*index+2)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" { // TODO: check with empty node
		return nil
	}

	strs := strings.Split(data, sep)
	nodes := make([]int, len(strs))
	for ix, s := range strs {
		nodes[ix], _ = strconv.Atoi(s)
	}

	return decode(nodes, 0)
}

func decode(nodes []int, index int) *TreeNode {
	if len(nodes) == 0 || index >= len(nodes) || nodes[index] == emptyNode {
		return nil
	}
	return &TreeNode{
		Val:   nodes[index],
		Left:  decode(nodes, 2*index+1),
		Right: decode(nodes, 2*index+2),
	}
}
