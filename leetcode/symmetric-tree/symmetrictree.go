package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	type pair struct {
		level int
		t     *TreeNode
	}

	q := list.New()
	q.PushBack(pair{level: 0, t: root})
	var levels [][]*int

	add := func(l int, v *int) {
		if l >= len(levels) {
			levels = append(levels, []*int{})
		}
		levels[l] = append(levels[l], v)
	}

	for q.Len() > 0 {
		p := q.Remove(q.Front()).(pair)
		if p.t == nil {
			add(p.level, nil)
			continue
		}
		add(p.level, &p.t.Val)

		q.PushBack(pair{level: p.level + 1, t: p.t.Left})
		q.PushBack(pair{level: p.level + 1, t: p.t.Right})
	}

	return isSym(levels)
}

func isSym(levels [][]*int) bool {

	// 	for l, nums := range levels {
	// 		fmt.Println("level", l)
	// 		for ix := range nums {
	// 			if nums[ix] == nil {
	// 				fmt.Print("nil, ")
	// 			} else {
	// 				fmt.Printf("%v, ", *nums[ix])
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}

	for _, nums := range levels {
		for a, b := 0, len(nums)-1; a < b; a++ {
			if nums[a] == nil && nums[b] == nil {
				b--
				continue
			}
			if nums[a] == nil || nums[b] == nil {
				return false
			}
			if *nums[a] != *nums[b] {
				return false
			}
			b--
		}
	}
	return true
}
