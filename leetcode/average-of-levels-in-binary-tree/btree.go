package btree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	type pair struct {
		level int
		tree  *TreeNode
	}

	levels := [][]int{}
	q := []*pair{&pair{0, root}}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node.level == len(levels) {
			levels = append(levels, []int{})
		}
		levels[node.level] = append(levels[node.level], node.tree.Val)

		if node.tree.Left != nil {
			q = append(q, &pair{node.level + 1, node.tree.Left})
		}
		if node.tree.Right != nil {
			q = append(q, &pair{node.level + 1, node.tree.Right})
		}
	}

	avgs := make([]float64, 0, len(levels))
	for _, res := range levels {
		var sum float64
		for _, n := range res {
			sum += float64(n)
		}
		avgs = append(avgs, sum/float64(len(res)))
	}

	return avgs
}
