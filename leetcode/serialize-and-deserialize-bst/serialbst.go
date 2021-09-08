package serialbst

import (
	"strconv"
	"strings"
)

const sep = "/"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	nodes []int
}

func Constructor() Codec {
	return Codec{
		nodes: make([]int, 0),
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	this.nodes = this.nodes[:0] // reset
	this.preOrder(root)

	strs := make([]string, len(this.nodes))
	for ix, n := range this.nodes {
		strs[ix] = strconv.Itoa(n)
	}

	return strings.Join(strs, sep)
}

func (this *Codec) preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	this.nodes = append(this.nodes, root.Val)
	this.preOrder(root.Left)
	this.preOrder(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	parts := strings.Split(data, sep)
	nodes := make([]int, len(parts))
	for ix, strVal := range parts {
		nodes[ix], _ = strconv.Atoi(strVal)
	}

	return this.decPreOrder(nodes)
}

func (this *Codec) decPreOrder(nodes []int) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}

	rootVal := nodes[0]
	ix := 1
	for ; ix < len(nodes); ix++ {
		if nodes[ix] >= rootVal {
			break
		}
	}

	return &TreeNode{
		Val:   rootVal,
		Left:  this.decPreOrder(nodes[1:ix]),
		Right: this.decPreOrder(nodes[ix:]),
	}
}
