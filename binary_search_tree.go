// https://play.golang.org/p/BGa6oDXUxY
//
package main

import (
	"container/list"
	"fmt"
)

// iterator of a in-order visit of a BST
type itr struct {
	current *tree
	stack   *list.List
}

func (i *itr) hasMore() bool {
	return i.stack.Len() > 0
}

func (i *itr) next() int {
	for i.stack.Len() > 0 {
		switch tc := i.stack.Remove(i.stack.Front()).(type) {
		case *tree:
			if tc.right != nil {
				i.stack.PushFront(tc.right)
			}

			i.stack.PushFront(tc.d)

			if tc.left != nil {
				i.stack.PushFront(tc.left)
			}
		case int:
			return tc
		}
	}

	return -1
}

type tree struct {
	d     int
	left  *tree
	right *tree
}

func (t *tree) Itr() *itr {
	stack := list.New()
	stack.PushFront(t)
	i := &itr{stack: stack, current: t}

	return i
}

// in-order result without recursivity
func (t *tree) inOrder() []int {
	stack := list.New()
	stack.PushFront(t)
	var result []int

	for stack.Len() > 0 {
		val := stack.Remove(stack.Front())
		if tc, isTree := val.(*tree); isTree {
			if tc.right != nil {
				stack.PushFront(tc.right)
			}
			stack.PushFront(tc.d)
			if tc.left != nil {
				stack.PushFront(tc.left)
			}
		} else {
			result = append(result, val.(int))
		}
	}
	return result
}

func (t *tree) isLeaf() bool {
	if t == nil {
		return false
	}
	return t.left == nil && t.right == nil
}

// get in-order list from a pre-order list on a BST
func getInOrder(pre []int) []int {
	if len(pre) == 0 {
		return nil
	}

	mid := pre[0]
	cursor := 1
	for ; cursor < len(pre) && pre[cursor] <= mid; cursor++ {
	}

	result := getInOrder(pre[1:cursor])
	result = append(result, mid)
	result = append(result, getInOrder(pre[cursor:])...)

	return result
}

// construct a binary search tree (BST) from a pre-order list
func rebuild(po []int) *tree {
	if len(po) == 0 {
		return nil
	}

	cursor := 1
	for ; cursor < len(po) && po[cursor] <= po[0]; cursor++ {
	}

	root := &tree{
		d:     po[0],
		left:  rebuild(po[1:cursor]),
		right: rebuild(po[cursor:]),
	}

	return root
}

func equals(a, b *tree) bool {
	if a == nil && b == nil {
		return true
	}

	if a.Height() != b.Height() {
		return false
	}
	if a.d != b.d {
		return false
	}

	return equals(a.left, b.left) && equals(a.right, b.right)
}

// PreOrder returns pre order list
func (t *tree) PreOrder() []int {
	if t == nil {
		return nil
	}
	result := []int{t.d}
	result = append(result, t.left.PreOrder()...)
	result = append(result, t.right.PreOrder()...)

	return result
}

// Insert a la' Binary Search Tree
func (t *tree) Insert(v int) {
	if t == nil {
		return
	}

	if v < t.d {
		if t.left == nil {
			t.left = &tree{d: v}
			return
		}
		t.left.Insert(v)
	} else {
		if t.right == nil {
			t.right = &tree{d: v}
			return
		}
		t.right.Insert(v)
	}
}

//
func (t *tree) Height() int {
	if t == nil {
		return 0
	}
	lh := t.left.Height()
	if rh := t.right.Height(); rh > lh {
		return rh + 1
	}
	return lh + 1
}

func (t *tree) printM(row, sc, ec int, m [][]int) {
	if t == nil {
		return
	}

	m[row][(sc+ec)/2] = t.d
	t.left.printM(row+2, sc, (sc+ec)/2, m)
	t.right.printM(row+2, (sc+ec)/2+1, ec, m)
}

// Prints a binary search tree in this way:
//				(root)
//                                9
//
//                4                               13
//
//        6               8
//
//                    5
//
func (t *tree) Print() {
	if t == nil {
		return
	}
	var m [][]int
	h := t.Height()
	var sq int

	if h <= 2 {
		var l, r int
		if t.left != nil {
			l = t.left.d
		}
		if t.right != nil {
			r = t.right.d
		}
		m = [][]int{{0, t.d, 0}, {0, 0, 0}, {l, 0, r}}
		sq = 3
	} else {
		m = make([][]int, h*2)
		sq = 1 << uint32(h-1)
		sq += (sq - 2) * 4 // ??

		for row := range m {
			m[row] = make([]int, sq)
		}

		t.printM(0, 0, sq, m)
	}

	for row := range m {
		for c := 0; c < sq; c++ {
			if m[row][c] != 0 {
				fmt.Print(m[row][c], " ")
			} else {
				fmt.Print(" ", " ")
			}
		}
		fmt.Println()
	}
}

func main() {
	r := &tree{
		d: 10,
		left: &tree{
			d: 7,
			left: &tree{
				d:     5,
				right: &tree{d: 6},
			},
		},
		right: &tree{
			d:    22,
			left: &tree{d: 13},
		},
	}

	fmt.Println(r.inOrder())
	it := r.Itr()
	for it.hasMore() {
		fmt.Print(it.next(), " ")
	}
	fmt.Println()

	a := &tree{
		d:    32,
		left: &tree{d: 13},
	}
	fmt.Println(getInOrder(a.PreOrder()))

	fmt.Println("order?", getInOrder(r.PreOrder()))

	r1 := &tree{d: 10}
	r1.Insert(7)
	r1.Insert(22)
	r1.Insert(13)
	r1.Insert(5)
	r1.Insert(6)
	r1.Print()
	fmt.Println("order?", getInOrder(r1.PreOrder()))

	fmt.Println("r == r1 ??", equals(r, r1))
	fmt.Println("r1 == r1 ??", equals(r1, r1))

	r2 := rebuild(r1.PreOrder())
	r2.Print()
	fmt.Println("r1 == r2 ??", equals(r1, r2))
	fmt.Println("r2 == r2 ??", equals(r2, r2))
}
