package maxancestor

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// solution #1 (fastest solution)
	return maxDiff(root, nil)

	// slowest solution, but using way less memory
	// solution #2
	// ancs := map[*TreeNode]*TreeNode{root: nil}
	// return maxDiffMap(root, ancs)

	// solution #3
	// ancs := map[int]int{root.Val: -1}
	// return maxDiffMapInt(root, ancs)
}

func maxDiffMapInt(root *TreeNode, ancs map[int]int) int {
	if root == nil {
		return 0
	}
	var max int
	tmp := root.Val
	for {
		parentVal := ancs[tmp]
		if parentVal == -1 {
			break
		}
		diff := int(math.Abs(float64(parentVal - root.Val)))
		if diff > max {
			max = diff
		}
		tmp = parentVal
	}

	if root.Left != nil {
		ancs[root.Left.Val] = root.Val
		if leftMax := maxDiffMapInt(root.Left, ancs); leftMax > max {
			max = leftMax
		}
	}
	if root.Right != nil {
		ancs[root.Right.Val] = root.Val
		if rightMax := maxDiffMapInt(root.Right, ancs); rightMax > max {
			max = rightMax
		}
	}

	return max
}

func maxDiffMap(root *TreeNode, ancs map[*TreeNode]*TreeNode) int {
	if root == nil {
		return 0
	}
	var max int
	tmp := root
	for {
		parent := ancs[tmp]
		if parent == nil {
			break
		}
		diff := int(math.Abs(float64(parent.Val - root.Val)))
		if diff > max {
			max = diff
		}
		tmp = parent
	}

	if root.Left != nil {
		ancs[root.Left] = root
		if leftMax := maxDiffMap(root.Left, ancs); leftMax > max {
			max = leftMax
		}
	}
	if root.Right != nil {
		ancs[root.Right] = root
		if rightMax := maxDiffMap(root.Right, ancs); rightMax > max {
			max = rightMax
		}
	}

	return max
}

func maxDiff(root *TreeNode, ancestors []int) int {
	if root == nil {
		return 0
	}

	var max int
	for _, aVal := range ancestors {
		diff := int(math.Abs(float64(aVal - root.Val)))
		if diff > max {
			max = diff
		}
	}

	leftAncs := make([]int, len(ancestors)+1)
	copy(leftAncs, ancestors)
	leftAncs[len(leftAncs)-1] = root.Val
	if leftMax := maxDiff(root.Left, leftAncs); leftMax > max {
		max = leftMax
	}

	rightAncs := make([]int, len(ancestors)+1)
	copy(rightAncs, ancestors)
	rightAncs[len(rightAncs)-1] = root.Val
	if rightMax := maxDiff(root.Right, rightAncs); rightMax > max {
		max = rightMax
	}

	return max
}
