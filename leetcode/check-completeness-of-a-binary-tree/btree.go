package btree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	type pair struct {
		level int
		tree  *TreeNode
	}

	levels := [][]*int{}
	q := []*pair{&pair{0, root}}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node.level == len(levels) {
			levels = append(levels, []*int{})
		}
		if node.tree == nil {
			levels[node.level] = append(levels[node.level], nil)
			continue
		}
		levels[node.level] = append(levels[node.level], &node.tree.Val)

		q = append(q, &pair{node.level + 1, node.tree.Left})
		q = append(q, &pair{node.level + 1, node.tree.Right})
	}

	// for _, res := range levels {
	// 	for _, n := range res {
	// 		if n != nil {
	// 			fmt.Printf("%3d, ", *n)
	// 		} else {
	// 			fmt.Print("nil, ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	for ix := 0; ix < len(levels)-2; ix++ {
		res := levels[ix]
		for ix := range res {
			if res[ix] == nil {
				return false
			}
		}
	}

	res := levels[len(levels)-2]
	var foundNil bool
	for ix := range res {
		if res[ix] == nil {
			foundNil = true
		} else {
			if foundNil {
				return false
			}
		}
	}

	return true
}
