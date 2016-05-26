package main

import (
	"container/list"
	"fmt"
)

var N, min int
var nodes, acuml []int
var tree []*list.List
var visited []bool

func dfs(x int) int {
	visited[x] = true

	var sum int
	for elm := tree[x].Front(); elm != nil; elm = elm.Next() {
		if y := elm.Value.(int); !visited[y] { // edge (x, y)
			sum += dfs(y)
		}
	}

	acuml[x] = sum + nodes[x]
	return acuml[x]
}

func checkwithdfs(x int) {
	visited[x] = true

	for elm := tree[x].Front(); elm != nil; elm = elm.Next() {
		if y := elm.Value.(int); !visited[y] { // edge (x, y)
			res := acuml[y] - (acuml[1] - acuml[y])
			if res < 0 {
				res *= -1
			}
			if res < min {
				min = res
			}
			checkwithdfs(y)
		}
	}
}

func main() {
	var p, s int
	fmt.Scan(&N)

	min = 1000000000
	nodes = make([]int, N+1)
	acuml = make([]int, N+1)
	tree = make([]*list.List, N+1)

	for ix := 1; ix < N+1; ix++ {
		fmt.Scan(&nodes[ix])
		tree[ix] = list.New()
	}

	for ix := 1; ix < N; ix++ {
		fmt.Scan(&p, &s)
		tree[p].PushBack(s)
		tree[s].PushBack(p)
	}

	visited = make([]bool, N+1)
	dfs(1)
	visited = make([]bool, N+1)
	checkwithdfs(1)
	fmt.Println(min)
}
