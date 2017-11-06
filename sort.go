package main

import (
	"sort"
)

// By is the type of a "less" function that defines the ordering of its NodeStruct arguments.
type By func(p1, p2 *NodeStruct) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(nodes []*NodeStruct) {
	ps := &nodeSorter{nodes, by}
	sort.Sort(ps)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type nodeSorter struct {
	nodes []*NodeStruct
	by    func(p1, p2 *NodeStruct) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *nodeSorter) Len() int {
	return len(s.nodes)
}

// Swap is part of sort.Interface.
func (s *nodeSorter) Swap(i, j int) {
	s.nodes[i], s.nodes[j] = s.nodes[j], s.nodes[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *nodeSorter) Less(i, j int) bool {
	return s.by(s.nodes[i], s.nodes[j])
}
