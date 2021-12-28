package main

import (
	"fmt"
)

func main() {
	t := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 5},
		Right: &TreeNode{Val: -5, Right: &TreeNode{Val: 5}},
	}
	fmt.Println(findFrequentTreeSum(t))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("%d", t.Val)
}

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	sums := map[int]int{}
	dfs(root, sums)
	fmt.Println(sums)
	var sol []int
	var max int
	for sum, freq := range sums {
		if freq > max {
			max = freq
			sol = sol[:0] // no reallocation
			sol = append(sol, sum)
		} else if freq == max {
			sol = append(sol, sum)
		}
	}
	return sol
}

func dfs(root *TreeNode, data map[int]int) int {
	if root == nil {
		return 0
	}
	sum := root.Val + dfs(root.Left, data) + dfs(root.Right, data)
	data[sum]++
	return sum
}
