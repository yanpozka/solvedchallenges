package btree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) (result [][]int) {
	if root == nil {
		return
	}

	type pair struct {
		level int
		tree  *TreeNode
	}

	result = [][]int{}
	q := []*pair{&pair{0, root}}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node.level == len(result) {
			result = append(result, []int{})
		}
		result[node.level] = append(result[node.level], node.tree.Val)

		if node.tree.Left != nil {
			q = append(q, &pair{node.level + 1, node.tree.Left})
		}
		if node.tree.Right != nil {
			q = append(q, &pair{node.level + 1, node.tree.Right})
		}
	}

	for i := len(result)/2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}

	return
}
