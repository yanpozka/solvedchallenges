package main

import "fmt"

func main() {

	n1 := &Node{Val: 2}
	n2 := &Node{Val: 3}
	root := &Node{
		Val:      1,
		Children: []*Node{n1, n2},
	}
	fmt.Println(levelOrder(root))
	fmt.Println(levelOrderDFS(root))
}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrderDFS(root *Node) [][]int {
	if root == nil {
		return nil
	}
	v := new(visitor)
	v.dfs(root, 0)
	return v.levels
}

type visitor struct {
	levels [][]int
}

func (v *visitor) dfs(root *Node, level int) {
	if level >= len(v.levels) {
		v.levels = append(v.levels, []int{})
	}
	v.levels[level] = append(v.levels[level], root.Val)

	for ix := range root.Children {
		v.dfs(root.Children[ix], level+1)
	}
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	type item struct {
		n     *Node
		level int
	}

	var levels [][]int

	q := []*item{{root, 0}}
	for len(q) > 0 {
		ci := q[0]
		q = q[1:]

		if ci.level >= len(levels) {
			levels = append(levels, []int{})
		}
		levels[ci.level] = append(levels[ci.level], ci.n.Val)

		for ix := range ci.n.Children {
			q = append(q, &item{ci.n.Children[ix], ci.level + 1})
		}
	}

	return levels
}
