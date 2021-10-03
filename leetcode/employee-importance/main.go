package main

import (
	"fmt"
)

func main() {
	fmt.Println(getImportance([]*Employee{
		{Id: 1, Importance: 5, Subordinates: []int{2, 3}},
		{Id: 2, Importance: 3},
		{Id: 3, Importance: 3, Subordinates: []int{4}},
		{Id: 4, Importance: 1},
	}, 1))
}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	graph := map[int]*Employee{}

	for ix, e := range employees {
		graph[e.Id] = employees[ix]
	}

	return depthSearch(graph, id)
}

func depthSearch(graph map[int]*Employee, rootID int) int {
	em := graph[rootID]
	sum := em.Importance

	for _, sID := range em.Subordinates {
		if sID == rootID {
			continue
		}
		sum += depthSearch(graph, sID)
	}

	return sum
}
